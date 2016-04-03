// This file was generated by counterfeiter
package core

import (
	"sync"

	"github.com/dev-op-spec/engine/core/models"
)

type fakeRunOperationUseCase struct {
	ExecuteStub        func(req models.RunOperationReq, namesOfAlreadyRunOperations []string) (operationRun models.OperationRunView, err error)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		req                        models.RunOperationReq
		namesOfAlreadyRunOperations []string
	}
	executeReturns struct {
		result1 models.OperationRunView
		result2 error
	}
}

func (fake *fakeRunOperationUseCase) Execute(req models.RunOperationReq, namesOfAlreadyRunOperations []string) (operationRun models.OperationRunView, err error) {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		req                        models.RunOperationReq
		namesOfAlreadyRunOperations []string
	}{req, namesOfAlreadyRunOperations})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(req, namesOfAlreadyRunOperations)
	} else {
		return fake.executeReturns.result1, fake.executeReturns.result2
	}
}

func (fake *fakeRunOperationUseCase) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *fakeRunOperationUseCase) ExecuteArgsForCall(i int) (models.RunOperationReq, []string) {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].req, fake.executeArgsForCall[i].namesOfAlreadyRunOperations
}

func (fake *fakeRunOperationUseCase) ExecuteReturns(result1 models.OperationRunView, result2 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 models.OperationRunView
		result2 error
	}{result1, result2}
}
