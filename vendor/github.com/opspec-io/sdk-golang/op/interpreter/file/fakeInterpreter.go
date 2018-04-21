// Code generated by counterfeiter. DO NOT EDIT.
package file

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type FakeInterpreter struct {
	InterpretStub        func(scope map[string]*model.Value, expression interface{}, opHandle model.DataHandle, scratchDir string) (*model.Value, error)
	interpretMutex       sync.RWMutex
	interpretArgsForCall []struct {
		scope      map[string]*model.Value
		expression interface{}
		opHandle   model.DataHandle
		scratchDir string
	}
	interpretReturns struct {
		result1 *model.Value
		result2 error
	}
	interpretReturnsOnCall map[int]struct {
		result1 *model.Value
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInterpreter) Interpret(scope map[string]*model.Value, expression interface{}, opHandle model.DataHandle, scratchDir string) (*model.Value, error) {
	fake.interpretMutex.Lock()
	ret, specificReturn := fake.interpretReturnsOnCall[len(fake.interpretArgsForCall)]
	fake.interpretArgsForCall = append(fake.interpretArgsForCall, struct {
		scope      map[string]*model.Value
		expression interface{}
		opHandle   model.DataHandle
		scratchDir string
	}{scope, expression, opHandle, scratchDir})
	fake.recordInvocation("Interpret", []interface{}{scope, expression, opHandle, scratchDir})
	fake.interpretMutex.Unlock()
	if fake.InterpretStub != nil {
		return fake.InterpretStub(scope, expression, opHandle, scratchDir)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.interpretReturns.result1, fake.interpretReturns.result2
}

func (fake *FakeInterpreter) InterpretCallCount() int {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return len(fake.interpretArgsForCall)
}

func (fake *FakeInterpreter) InterpretArgsForCall(i int) (map[string]*model.Value, interface{}, model.DataHandle, string) {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return fake.interpretArgsForCall[i].scope, fake.interpretArgsForCall[i].expression, fake.interpretArgsForCall[i].opHandle, fake.interpretArgsForCall[i].scratchDir
}

func (fake *FakeInterpreter) InterpretReturns(result1 *model.Value, result2 error) {
	fake.InterpretStub = nil
	fake.interpretReturns = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *FakeInterpreter) InterpretReturnsOnCall(i int, result1 *model.Value, result2 error) {
	fake.InterpretStub = nil
	if fake.interpretReturnsOnCall == nil {
		fake.interpretReturnsOnCall = make(map[int]struct {
			result1 *model.Value
			result2 error
		})
	}
	fake.interpretReturnsOnCall[i] = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *FakeInterpreter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInterpreter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ Interpreter = new(FakeInterpreter)