// This file was generated by counterfeiter
package core

import (
  "sync"
  "github.com/dev-op-spec/engine/core/models"
)

type fakeListDevOpsUseCase struct {
  ExecuteStub        func() (devOps []models.DevOpView, err error)
  executeMutex       sync.RWMutex
  executeArgsForCall []struct{}
  executeReturns     struct {
                       result1 []models.DevOpView
                       result2 error
                     }
}

func (fake *fakeListDevOpsUseCase) Execute() (devOps []models.DevOpView, err error) {
  fake.executeMutex.Lock()
  fake.executeArgsForCall = append(fake.executeArgsForCall, struct{}{})
  fake.executeMutex.Unlock()
  if fake.ExecuteStub != nil {
    return fake.ExecuteStub()
  } else {
    return fake.executeReturns.result1, fake.executeReturns.result2
  }
}

func (fake *fakeListDevOpsUseCase) ExecuteCallCount() int {
  fake.executeMutex.RLock()
  defer fake.executeMutex.RUnlock()
  return len(fake.executeArgsForCall)
}

func (fake *fakeListDevOpsUseCase) ExecuteReturns(result1 []models.DevOpView, result2 error) {
  fake.ExecuteStub = nil
  fake.executeReturns = struct {
    result1 []models.DevOpView
    result2 error
  }{result1, result2}
}