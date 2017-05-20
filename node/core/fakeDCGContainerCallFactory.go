// Code generated by counterfeiter. DO NOT EDIT.
package core

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type fakeDCGContainerCallFactory struct {
	ConstructStub        func(currentScope map[string]*model.Data, scgContainerCall *model.SCGContainerCall, containerId string, rootOpId string, pkgRef string) (*model.DCGContainerCall, error)
	constructMutex       sync.RWMutex
	constructArgsForCall []struct {
		currentScope     map[string]*model.Data
		scgContainerCall *model.SCGContainerCall
		containerId      string
		rootOpId         string
		pkgRef           string
	}
	constructReturns struct {
		result1 *model.DCGContainerCall
		result2 error
	}
	constructReturnsOnCall map[int]struct {
		result1 *model.DCGContainerCall
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeDCGContainerCallFactory) Construct(currentScope map[string]*model.Data, scgContainerCall *model.SCGContainerCall, containerId string, rootOpId string, pkgRef string) (*model.DCGContainerCall, error) {
	fake.constructMutex.Lock()
	ret, specificReturn := fake.constructReturnsOnCall[len(fake.constructArgsForCall)]
	fake.constructArgsForCall = append(fake.constructArgsForCall, struct {
		currentScope     map[string]*model.Data
		scgContainerCall *model.SCGContainerCall
		containerId      string
		rootOpId         string
		pkgRef           string
	}{currentScope, scgContainerCall, containerId, rootOpId, pkgRef})
	fake.recordInvocation("Construct", []interface{}{currentScope, scgContainerCall, containerId, rootOpId, pkgRef})
	fake.constructMutex.Unlock()
	if fake.ConstructStub != nil {
		return fake.ConstructStub(currentScope, scgContainerCall, containerId, rootOpId, pkgRef)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.constructReturns.result1, fake.constructReturns.result2
}

func (fake *fakeDCGContainerCallFactory) ConstructCallCount() int {
	fake.constructMutex.RLock()
	defer fake.constructMutex.RUnlock()
	return len(fake.constructArgsForCall)
}

func (fake *fakeDCGContainerCallFactory) ConstructArgsForCall(i int) (map[string]*model.Data, *model.SCGContainerCall, string, string, string) {
	fake.constructMutex.RLock()
	defer fake.constructMutex.RUnlock()
	return fake.constructArgsForCall[i].currentScope, fake.constructArgsForCall[i].scgContainerCall, fake.constructArgsForCall[i].containerId, fake.constructArgsForCall[i].rootOpId, fake.constructArgsForCall[i].pkgRef
}

func (fake *fakeDCGContainerCallFactory) ConstructReturns(result1 *model.DCGContainerCall, result2 error) {
	fake.ConstructStub = nil
	fake.constructReturns = struct {
		result1 *model.DCGContainerCall
		result2 error
	}{result1, result2}
}

func (fake *fakeDCGContainerCallFactory) ConstructReturnsOnCall(i int, result1 *model.DCGContainerCall, result2 error) {
	fake.ConstructStub = nil
	if fake.constructReturnsOnCall == nil {
		fake.constructReturnsOnCall = make(map[int]struct {
			result1 *model.DCGContainerCall
			result2 error
		})
	}
	fake.constructReturnsOnCall[i] = struct {
		result1 *model.DCGContainerCall
		result2 error
	}{result1, result2}
}

func (fake *fakeDCGContainerCallFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.constructMutex.RLock()
	defer fake.constructMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *fakeDCGContainerCallFactory) recordInvocation(key string, args []interface{}) {
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