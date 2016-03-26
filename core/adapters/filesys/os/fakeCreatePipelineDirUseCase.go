// This file was generated by counterfeiter
package os

import "sync"

type FakeCreatePipelineDirUseCase struct {
  ExecuteStub        func(pipelineName string) (err error)
  executeMutex       sync.RWMutex
  executeArgsForCall []struct {
    pipelineName string
  }
  executeReturns     struct {
                       result1 error
                     }
}

func (fake *FakeCreatePipelineDirUseCase) Execute(pipelineName string) (err error) {
  fake.executeMutex.Lock()
  fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
    pipelineName string
  }{pipelineName})
  fake.executeMutex.Unlock()
  if fake.ExecuteStub != nil {
    return fake.ExecuteStub(pipelineName)
  } else {
    return fake.executeReturns.result1
  }
}

func (fake *FakeCreatePipelineDirUseCase) ExecuteCallCount() int {
  fake.executeMutex.RLock()
  defer fake.executeMutex.RUnlock()
  return len(fake.executeArgsForCall)
}

func (fake *FakeCreatePipelineDirUseCase) ExecuteArgsForCall(i int) string {
  fake.executeMutex.RLock()
  defer fake.executeMutex.RUnlock()
  return fake.executeArgsForCall[i].pipelineName
}

func (fake *FakeCreatePipelineDirUseCase) ExecuteReturns(result1 error) {
  fake.ExecuteStub = nil
  fake.executeReturns = struct {
    result1 error
  }{result1}
}