// This file was generated by counterfeiter
package core

import "sync"

type fakeOpRunner struct {
  RunStub          func(opRunArgs map[string]string, opRef string, opRunId string, rootOpRunId string) (err error)
  runMutex         sync.RWMutex
  runArgsForCall   []struct {
    opRunArgs   map[string]string
    opRef       string
    opRunId     string
    rootOpRunId string
  }
  runReturns       struct {
                     result1 error
                   }
  invocations      map[string][][]interface{}
  invocationsMutex sync.RWMutex
}

func (fake *fakeOpRunner) Run(opRunArgs map[string]string, opRef string, opRunId string, rootOpRunId string) (err error) {
  fake.runMutex.Lock()
  fake.runArgsForCall = append(fake.runArgsForCall, struct {
    opRunArgs   map[string]string
    opRef       string
    opRunId     string
    rootOpRunId string
  }{opRunArgs, opRef, opRunId, rootOpRunId})
  fake.recordInvocation("Run", []interface{}{opRunArgs, opRef, opRunId, rootOpRunId})
  fake.runMutex.Unlock()
  if fake.RunStub != nil {
    return fake.RunStub(opRunArgs, opRef, opRunId, rootOpRunId)
  } else {
    return fake.runReturns.result1
  }
}

func (fake *fakeOpRunner) RunCallCount() int {
  fake.runMutex.RLock()
  defer fake.runMutex.RUnlock()
  return len(fake.runArgsForCall)
}

func (fake *fakeOpRunner) RunArgsForCall(i int) (map[string]string, string, string, string) {
  fake.runMutex.RLock()
  defer fake.runMutex.RUnlock()
  return fake.runArgsForCall[i].opRunArgs, fake.runArgsForCall[i].opRef, fake.runArgsForCall[i].opRunId, fake.runArgsForCall[i].rootOpRunId
}

func (fake *fakeOpRunner) RunReturns(result1 error) {
  fake.RunStub = nil
  fake.runReturns = struct {
    result1 error
  }{result1}
}

func (fake *fakeOpRunner) Invocations() map[string][][]interface{} {
  fake.invocationsMutex.RLock()
  defer fake.invocationsMutex.RUnlock()
  fake.runMutex.RLock()
  defer fake.runMutex.RUnlock()
  return fake.invocations
}

func (fake *fakeOpRunner) recordInvocation(key string, args []interface{}) {
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
