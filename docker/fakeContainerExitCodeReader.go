// This file was generated by counterfeiter
package docker

import "sync"

type fakeContainerExitCodeReader struct {
  ReadStub         func(opBundlePath string, opName string, opNamespace string) (opExitCode int, err error)
  readMutex        sync.RWMutex
  readArgsForCall  []struct {
    opBundlePath string
    opName       string
    opNamespace  string
  }
  readReturns      struct {
                     result1 int
                     result2 error
                   }
  invocations      map[string][][]interface{}
  invocationsMutex sync.RWMutex
}

func (fake *fakeContainerExitCodeReader) Read(opBundlePath string, opName string, opNamespace string) (opExitCode int, err error) {
  fake.readMutex.Lock()
  fake.readArgsForCall = append(fake.readArgsForCall, struct {
    opBundlePath string
    opName       string
    opNamespace  string
  }{opBundlePath, opName, opNamespace})
  fake.recordInvocation("Read", []interface{}{opBundlePath, opName, opNamespace})
  fake.readMutex.Unlock()
  if fake.ReadStub != nil {
    return fake.ReadStub(opBundlePath, opName, opNamespace)
  } else {
    return fake.readReturns.result1, fake.readReturns.result2
  }
}

func (fake *fakeContainerExitCodeReader) ReadCallCount() int {
  fake.readMutex.RLock()
  defer fake.readMutex.RUnlock()
  return len(fake.readArgsForCall)
}

func (fake *fakeContainerExitCodeReader) ReadArgsForCall(i int) (string, string, string) {
  fake.readMutex.RLock()
  defer fake.readMutex.RUnlock()
  return fake.readArgsForCall[i].opBundlePath, fake.readArgsForCall[i].opName, fake.readArgsForCall[i].opNamespace
}

func (fake *fakeContainerExitCodeReader) ReadReturns(result1 int, result2 error) {
  fake.ReadStub = nil
  fake.readReturns = struct {
    result1 int
    result2 error
  }{result1, result2}
}

func (fake *fakeContainerExitCodeReader) Invocations() map[string][][]interface{} {
  fake.invocationsMutex.RLock()
  defer fake.invocationsMutex.RUnlock()
  fake.readMutex.RLock()
  defer fake.readMutex.RUnlock()
  return fake.invocations
}

func (fake *fakeContainerExitCodeReader) recordInvocation(key string, args []interface{}) {
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
