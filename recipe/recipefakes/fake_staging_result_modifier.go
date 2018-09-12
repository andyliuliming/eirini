// Code generated by counterfeiter. DO NOT EDIT.
package recipefakes

import (
	"sync"

	bap "code.cloudfoundry.org/buildpackapplifecycle"
	"code.cloudfoundry.org/eirini/recipe"
)

type FakeStagingResultModifier struct {
	ModifyStub        func(result bap.StagingResult) (bap.StagingResult, error)
	modifyMutex       sync.RWMutex
	modifyArgsForCall []struct {
		result bap.StagingResult
	}
	modifyReturns struct {
		result1 bap.StagingResult
		result2 error
	}
	modifyReturnsOnCall map[int]struct {
		result1 bap.StagingResult
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStagingResultModifier) Modify(result bap.StagingResult) (bap.StagingResult, error) {
	fake.modifyMutex.Lock()
	ret, specificReturn := fake.modifyReturnsOnCall[len(fake.modifyArgsForCall)]
	fake.modifyArgsForCall = append(fake.modifyArgsForCall, struct {
		result bap.StagingResult
	}{result})
	fake.recordInvocation("Modify", []interface{}{result})
	fake.modifyMutex.Unlock()
	if fake.ModifyStub != nil {
		return fake.ModifyStub(result)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.modifyReturns.result1, fake.modifyReturns.result2
}

func (fake *FakeStagingResultModifier) ModifyCallCount() int {
	fake.modifyMutex.RLock()
	defer fake.modifyMutex.RUnlock()
	return len(fake.modifyArgsForCall)
}

func (fake *FakeStagingResultModifier) ModifyArgsForCall(i int) bap.StagingResult {
	fake.modifyMutex.RLock()
	defer fake.modifyMutex.RUnlock()
	return fake.modifyArgsForCall[i].result
}

func (fake *FakeStagingResultModifier) ModifyReturns(result1 bap.StagingResult, result2 error) {
	fake.ModifyStub = nil
	fake.modifyReturns = struct {
		result1 bap.StagingResult
		result2 error
	}{result1, result2}
}

func (fake *FakeStagingResultModifier) ModifyReturnsOnCall(i int, result1 bap.StagingResult, result2 error) {
	fake.ModifyStub = nil
	if fake.modifyReturnsOnCall == nil {
		fake.modifyReturnsOnCall = make(map[int]struct {
			result1 bap.StagingResult
			result2 error
		})
	}
	fake.modifyReturnsOnCall[i] = struct {
		result1 bap.StagingResult
		result2 error
	}{result1, result2}
}

func (fake *FakeStagingResultModifier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.modifyMutex.RLock()
	defer fake.modifyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStagingResultModifier) recordInvocation(key string, args []interface{}) {
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

var _ recipe.StagingResultModifier = new(FakeStagingResultModifier)