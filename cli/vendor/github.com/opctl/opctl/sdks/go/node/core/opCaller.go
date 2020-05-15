package core

import (
	"context"
	"errors"
	"path/filepath"
	"time"

	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/op/outputs"
	"github.com/opctl/opctl/sdks/go/opspec/opfile"
	"github.com/opctl/opctl/sdks/go/pubsub"
)

//counterfeiter:generate -o internal/fakes/opCaller.go . opCaller
type opCaller interface {
	// Executes an op call
	Call(
		parentCtx context.Context,
		dcgOpCall *model.DCGOpCall,
		inboundScope map[string]*model.Value,
		parentCallID *string,
		scgOpCall *model.SCGOpCall,
	)
}

func newOpCaller(
	callStore callStore,
	pubSub pubsub.PubSub,
	caller caller,
	dataDirPath string,
) opCaller {
	return _opCaller{
		caller:             caller,
		callStore:          callStore,
		dcgScratchDir:      filepath.Join(dataDirPath, "dcg"),
		outputsInterpreter: outputs.NewInterpreter(),
		opFileGetter:       opfile.NewGetter(),
		pubSub:             pubSub,
	}
}

type _opCaller struct {
	callStore          callStore
	caller             caller
	dcgScratchDir      string
	outputsInterpreter outputs.Interpreter
	opFileGetter       opfile.Getter
	pubSub             pubsub.PubSub
}

func (oc _opCaller) Call(
	parentCtx context.Context,
	dcgOpCall *model.DCGOpCall,
	inboundScope map[string]*model.Value,
	parentCallID *string,
	scgOpCall *model.SCGOpCall,
) {
	// setup cancellation
	opCtx, cancelOp := context.WithCancel(parentCtx)
	defer cancelOp()

	var err error
	outboundScope := map[string]*model.Value{}

	defer func() {
		// defer must be defined before conditional return statements so it always runs
		select {
		case <-parentCtx.Done():
			// if parent context cancelled; NOOP
		default:
			var opOutcome string
			if dcg := oc.callStore.TryGet(dcgOpCall.OpID); nil != dcg && dcg.IsKilled {
				opOutcome = model.OpOutcomeKilled
			} else if nil != err {
				oc.pubSub.Publish(
					model.Event{
						Timestamp: time.Now().UTC(),
						OpErred: &model.OpErredEvent{
							Msg:      err.Error(),
							OpID:     dcgOpCall.OpID,
							OpRef:    dcgOpCall.OpPath,
							RootOpID: dcgOpCall.RootOpID,
						},
					},
				)
				opOutcome = model.OpOutcomeFailed
			} else {
				opOutcome = model.OpOutcomeSucceeded
			}

			event := model.Event{
				Timestamp: time.Now().UTC(),
				OpEnded: &model.OpEndedEvent{
					OpID:     dcgOpCall.OpID,
					OpRef:    dcgOpCall.OpPath,
					Outcome:  opOutcome,
					RootOpID: dcgOpCall.RootOpID,
					Outputs:  outboundScope,
				},
			}

			if nil != err {
				event.OpEnded.Error = &model.CallEndedEventError{
					Message: err.Error(),
				}
			}

			oc.pubSub.Publish(event)
		}

	}()

	opStartedTime := time.Now().UTC()

	oc.pubSub.Publish(
		model.Event{
			Timestamp: opStartedTime,
			OpStarted: &model.OpStartedEvent{
				OpID:     dcgOpCall.OpID,
				OpRef:    dcgOpCall.OpPath,
				RootOpID: dcgOpCall.RootOpID,
			},
		},
	)

	// form scope for op call by combining defined inputs & op dir
	opCallScope := map[string]*model.Value{}
	for varName, varData := range dcgOpCall.Inputs {
		opCallScope[varName] = varData
	}
	opCallScope["/"] = &model.Value{
		Dir: &dcgOpCall.OpPath,
	}

	oc.caller.Call(
		opCtx,
		dcgOpCall.ChildCallID,
		opCallScope,
		dcgOpCall.ChildCallSCG,
		dcgOpCall.OpPath,
		&dcgOpCall.OpID,
		dcgOpCall.RootOpID,
	)

	// subscribe to events
	eventChannel, _ := oc.pubSub.Subscribe(
		opCtx,
		model.EventFilter{
			Roots: []string{dcgOpCall.RootOpID},
			Since: &opStartedTime,
		},
	)

	opOutputs := map[string]*model.Value{}

eventLoop:
	for event := range eventChannel {
		switch {
		case nil != event.OpEnded && event.OpEnded.OpID == dcgOpCall.OpID:
			// parent ended prematurely
			return
		case nil != event.CallEnded && event.CallEnded.CallID == dcgOpCall.ChildCallID:
			if nil != event.CallEnded.Error {
				err = errors.New(event.CallEnded.Error.Message)
				// end on any error
				cancelOp()
				return
			}
			for name, value := range event.CallEnded.Outputs {
				opOutputs[name] = value
			}
			break eventLoop
		}
	}

	var opFile *model.OpFile
	opFile, err = oc.opFileGetter.Get(
		opCtx,
		dcgOpCall.OpPath,
	)
	if nil != err {
		cancelOp()
		return
	}
	opOutputs, err = oc.outputsInterpreter.Interpret(
		opOutputs,
		opFile.Outputs,
		dcgOpCall.OpPath,
		filepath.Join(oc.dcgScratchDir, dcgOpCall.OpID),
	)

	// filter op outboundScope to bound call outboundScope
	for boundName, boundValue := range scgOpCall.Outputs {
		// return bound outboundScope
		if "" == boundValue {
			// implicit value
			boundValue = boundName
		}
		for opOutputName, opOutputValue := range opOutputs {
			if boundValue == opOutputName {
				outboundScope[boundName] = opOutputValue
			}
		}
	}
}
