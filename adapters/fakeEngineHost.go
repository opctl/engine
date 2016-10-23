// This file was generated by counterfeiter
package adapters

import (
	"sync"
)

type FakeEngineHost struct {
	EnsureEngineRunningStub        func() (err error)
	ensureEngineRunningMutex       sync.RWMutex
	ensureEngineRunningArgsForCall []struct{}
	ensureEngineRunningReturns     struct {
		result1 error
	}
	GetEngineProtocolRelativeBaseUrlStub        func() (protocolRelativeBaseUrl string, err error)
	getEngineProtocolRelativeBaseUrlMutex       sync.RWMutex
	getEngineProtocolRelativeBaseUrlArgsForCall []struct{}
	getEngineProtocolRelativeBaseUrlReturns     struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEngineHost) EnsureEngineRunning() (err error) {
	fake.ensureEngineRunningMutex.Lock()
	fake.ensureEngineRunningArgsForCall = append(fake.ensureEngineRunningArgsForCall, struct{}{})
	fake.recordInvocation("EnsureEngineRunning", []interface{}{})
	fake.ensureEngineRunningMutex.Unlock()
	if fake.EnsureEngineRunningStub != nil {
		return fake.EnsureEngineRunningStub()
	} else {
		return fake.ensureEngineRunningReturns.result1
	}
}

func (fake *FakeEngineHost) EnsureEngineRunningCallCount() int {
	fake.ensureEngineRunningMutex.RLock()
	defer fake.ensureEngineRunningMutex.RUnlock()
	return len(fake.ensureEngineRunningArgsForCall)
}

func (fake *FakeEngineHost) EnsureEngineRunningReturns(result1 error) {
	fake.EnsureEngineRunningStub = nil
	fake.ensureEngineRunningReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEngineHost) GetEngineProtocolRelativeBaseUrl() (protocolRelativeBaseUrl string, err error) {
	fake.getEngineProtocolRelativeBaseUrlMutex.Lock()
	fake.getEngineProtocolRelativeBaseUrlArgsForCall = append(fake.getEngineProtocolRelativeBaseUrlArgsForCall, struct{}{})
	fake.recordInvocation("GetEngineProtocolRelativeBaseUrl", []interface{}{})
	fake.getEngineProtocolRelativeBaseUrlMutex.Unlock()
	if fake.GetEngineProtocolRelativeBaseUrlStub != nil {
		return fake.GetEngineProtocolRelativeBaseUrlStub()
	} else {
		return fake.getEngineProtocolRelativeBaseUrlReturns.result1, fake.getEngineProtocolRelativeBaseUrlReturns.result2
	}
}

func (fake *FakeEngineHost) GetEngineProtocolRelativeBaseUrlCallCount() int {
	fake.getEngineProtocolRelativeBaseUrlMutex.RLock()
	defer fake.getEngineProtocolRelativeBaseUrlMutex.RUnlock()
	return len(fake.getEngineProtocolRelativeBaseUrlArgsForCall)
}

func (fake *FakeEngineHost) GetEngineProtocolRelativeBaseUrlReturns(result1 string, result2 error) {
	fake.GetEngineProtocolRelativeBaseUrlStub = nil
	fake.getEngineProtocolRelativeBaseUrlReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeEngineHost) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.ensureEngineRunningMutex.RLock()
	defer fake.ensureEngineRunningMutex.RUnlock()
	fake.getEngineProtocolRelativeBaseUrlMutex.RLock()
	defer fake.getEngineProtocolRelativeBaseUrlMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeEngineHost) recordInvocation(key string, args []interface{}) {
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

var _ EngineHost = new(FakeEngineHost)
