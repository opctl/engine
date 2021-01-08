package core

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/golang-interfaces/iio"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/node/core/containerruntime"
)

//counterfeiter:generate -o internal/fakes/containerCaller.go . containerCaller
type containerCaller interface {
	// Executes a container call
	Call(
		ctx context.Context,
		containerCall *model.ContainerCall,
		inboundScope map[string]*model.Value,
		containerCallSpec *model.ContainerCallSpec,
		rootCallID string,
	) (
		map[string]*model.Value,
		error,
	)
}

func newContainerCaller(
	containerRuntime containerruntime.ContainerRuntime,
	eventChannel chan model.Event,
	stateStore stateStore,
) containerCaller {

	return _containerCaller{
		containerRuntime: containerRuntime,
		eventChannel:     eventChannel,
		stateStore:       stateStore,
		io:               iio.New(),
	}

}

type _containerCaller struct {
	containerRuntime containerruntime.ContainerRuntime
	eventChannel     chan model.Event
	stateStore       stateStore
	io               iio.IIO
}

func (cc _containerCaller) Call(
	ctx context.Context,
	containerCall *model.ContainerCall,
	inboundScope map[string]*model.Value,
	containerCallSpec *model.ContainerCallSpec,
	rootCallID string,
) (
	map[string]*model.Value,
	error,
) {
	outputs := map[string]*model.Value{}
	var exitCode int64

	if nil != containerCall.Image.Ref && nil == containerCall.Image.PullCreds {
		if auth := cc.stateStore.TryGetAuth(*containerCall.Image.Ref); nil != auth {
			containerCall.Image.PullCreds = &auth.Creds
		}
	}

	logStdOutPR, logStdOutPW := cc.io.Pipe()
	logStdErrPR, logStdErrPW := cc.io.Pipe()

	// interpret logs
	logChan := make(chan error, 1)
	go func() {
		logChan <- cc.interpretLogs(
			logStdOutPR,
			logStdErrPR,
			containerCall,
			rootCallID,
		)
	}()

	outputs = cc.interpretOutputs(
		containerCallSpec,
		containerCall,
	)

	rawExitCode, err := cc.containerRuntime.RunContainer(
		ctx,
		containerCall,
		rootCallID,
		cc.eventChannel,
		logStdOutPW,
		logStdErrPW,
	)

	// @TODO: handle no exit code
	if nil != rawExitCode {
		exitCode = *rawExitCode
	}

	if exitCode != 0 {
		err = fmt.Errorf("container exited with %d", exitCode)
	}

	// wait on logChan
	if logChanErr := <-logChan; nil == err {
		// non-destructively set err
		err = logChanErr
	}

	return outputs, err
}

func (this _containerCaller) interpretLogs(
	stdOutReader io.Reader,
	stdErrReader io.Reader,
	containerCall *model.ContainerCall,
	rootCallID string,
) error {
	stdOutLogChan := make(chan error, 1)
	go func() {
		// interpret stdOut
		stdOutLogChan <- readChunks(
			stdOutReader,
			func(chunk []byte) {
				this.eventChannel <- model.Event{
					Timestamp: time.Now().UTC(),
					ContainerStdOutWrittenTo: &model.ContainerStdOutWrittenTo{
						Data:        chunk,
						ContainerID: containerCall.ContainerID,
						OpRef:       containerCall.OpPath,
						RootCallID:  rootCallID,
					},
				}
			})
	}()

	stdErrLogChan := make(chan error, 1)
	go func() {
		// interpret stdErr
		stdErrLogChan <- readChunks(
			stdErrReader,
			func(chunk []byte) {
				this.eventChannel <- model.Event{
					Timestamp: time.Now().UTC(),
					ContainerStdErrWrittenTo: &model.ContainerStdErrWrittenTo{
						Data:        chunk,
						ContainerID: containerCall.ContainerID,
						OpRef:       containerCall.OpPath,
						RootCallID:  rootCallID,
					},
				}
			})
	}()

	// wait on logs
	stdOutLogErr := <-stdOutLogChan
	stdErrLogErr := <-stdErrLogChan

	// return errs
	if nil != stdOutLogErr {
		return stdOutLogErr
	}
	if nil != stdErrLogErr {
		return stdErrLogErr
	}

	return nil
}

func (this _containerCaller) interpretOutputs(
	containerCallSpec *model.ContainerCallSpec,
	containerCall *model.ContainerCall,
) map[string]*model.Value {
	outputs := map[string]*model.Value{}

	for socketAddr, name := range containerCallSpec.Sockets {
		// add socket outputs
		if "0.0.0.0" == socketAddr {
			outputs[name] = &model.Value{Socket: &containerCall.ContainerID}
		}
	}
	for callSpecContainerFilePath, mountSrc := range containerCallSpec.Files {
		mountSrcStr, ok := mountSrc.(string)
		if !ok {
			continue
		}

		if "" == mountSrcStr {
			// skip embedded files
			continue
		}

		// add file outputs
		for callContainerFilePath, callHostFilePath := range containerCall.Files {
			if callSpecContainerFilePath == callContainerFilePath {
				// copy callHostFilePath before taking address; range vars have same address for every iteration
				value := callHostFilePath
				outputs[strings.TrimSuffix(strings.TrimPrefix(mountSrcStr, "$("), ")")] = &model.Value{File: &value}
			}
		}
	}
	for callSpecContainerDirPath, mountSrc := range containerCallSpec.Dirs {
		mountSrcStr, ok := mountSrc.(string)
		if !ok {
			continue
		}

		if "" == mountSrcStr {
			// skip embedded dirs
			continue
		}

		// add dir outputs
		for callContainerDirPath, callHostDirPath := range containerCall.Dirs {
			if callSpecContainerDirPath == callContainerDirPath {
				// copy callHostDirPath before taking address; range vars have same address for every iteration
				value := callHostDirPath
				outputs[strings.TrimSuffix(strings.TrimPrefix(mountSrcStr, "$("), ")")] = &model.Value{Dir: &value}
			}
		}
	}

	return outputs
}
