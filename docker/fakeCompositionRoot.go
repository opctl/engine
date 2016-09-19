// This file was generated by counterfeiter
package docker

import (
	"sync"
)

type fakeCompositionRoot struct {
	StartContainerUseCaseStub        func() startContainerUseCase
	startContainerUseCaseMutex       sync.RWMutex
	startContainerUseCaseArgsForCall []struct{}
	startContainerUseCaseReturns     struct {
		result1 startContainerUseCase
	}
	EnsureContainerRemovedUseCaseStub        func() ensureContainerRemovedUseCase
	ensureContainerRemovedUseCaseMutex       sync.RWMutex
	ensureContainerRemovedUseCaseArgsForCall []struct{}
	ensureContainerRemovedUseCaseReturns     struct {
		result1 ensureContainerRemovedUseCase
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeCompositionRoot) StartContainerUseCase() startContainerUseCase {
	fake.startContainerUseCaseMutex.Lock()
	fake.startContainerUseCaseArgsForCall = append(fake.startContainerUseCaseArgsForCall, struct{}{})
	fake.recordInvocation("StartContainerUseCase", []interface{}{})
	fake.startContainerUseCaseMutex.Unlock()
	if fake.StartContainerUseCaseStub != nil {
		return fake.StartContainerUseCaseStub()
	} else {
		return fake.startContainerUseCaseReturns.result1
	}
}

func (fake *fakeCompositionRoot) StartContainerUseCaseCallCount() int {
	fake.startContainerUseCaseMutex.RLock()
	defer fake.startContainerUseCaseMutex.RUnlock()
	return len(fake.startContainerUseCaseArgsForCall)
}

func (fake *fakeCompositionRoot) StartContainerUseCaseReturns(result1 startContainerUseCase) {
	fake.StartContainerUseCaseStub = nil
	fake.startContainerUseCaseReturns = struct {
		result1 startContainerUseCase
	}{result1}
}

func (fake *fakeCompositionRoot) EnsureContainerRemovedUseCase() ensureContainerRemovedUseCase {
	fake.ensureContainerRemovedUseCaseMutex.Lock()
	fake.ensureContainerRemovedUseCaseArgsForCall = append(fake.ensureContainerRemovedUseCaseArgsForCall, struct{}{})
	fake.recordInvocation("EnsureContainerRemovedUseCase", []interface{}{})
	fake.ensureContainerRemovedUseCaseMutex.Unlock()
	if fake.EnsureContainerRemovedUseCaseStub != nil {
		return fake.EnsureContainerRemovedUseCaseStub()
	} else {
		return fake.ensureContainerRemovedUseCaseReturns.result1
	}
}

func (fake *fakeCompositionRoot) EnsureContainerRemovedUseCaseCallCount() int {
	fake.ensureContainerRemovedUseCaseMutex.RLock()
	defer fake.ensureContainerRemovedUseCaseMutex.RUnlock()
	return len(fake.ensureContainerRemovedUseCaseArgsForCall)
}

func (fake *fakeCompositionRoot) EnsureContainerRemovedUseCaseReturns(result1 ensureContainerRemovedUseCase) {
	fake.EnsureContainerRemovedUseCaseStub = nil
	fake.ensureContainerRemovedUseCaseReturns = struct {
		result1 ensureContainerRemovedUseCase
	}{result1}
}

func (fake *fakeCompositionRoot) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startContainerUseCaseMutex.RLock()
	defer fake.startContainerUseCaseMutex.RUnlock()
	fake.ensureContainerRemovedUseCaseMutex.RLock()
	defer fake.ensureContainerRemovedUseCaseMutex.RUnlock()
	return fake.invocations
}

func (fake *fakeCompositionRoot) recordInvocation(key string, args []interface{}) {
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