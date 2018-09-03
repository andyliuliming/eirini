// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2v3action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeAppSummaryActor struct {
	GetApplicationSummaryByNameAndSpaceStub        func(appName string, spaceGUID string, withObfuscatedValues bool) (v2v3action.ApplicationSummary, v2v3action.Warnings, error)
	getApplicationSummaryByNameAndSpaceMutex       sync.RWMutex
	getApplicationSummaryByNameAndSpaceArgsForCall []struct {
		appName              string
		spaceGUID            string
		withObfuscatedValues bool
	}
	getApplicationSummaryByNameAndSpaceReturns struct {
		result1 v2v3action.ApplicationSummary
		result2 v2v3action.Warnings
		result3 error
	}
	getApplicationSummaryByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v2v3action.ApplicationSummary
		result2 v2v3action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAppSummaryActor) GetApplicationSummaryByNameAndSpace(appName string, spaceGUID string, withObfuscatedValues bool) (v2v3action.ApplicationSummary, v2v3action.Warnings, error) {
	fake.getApplicationSummaryByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationSummaryByNameAndSpaceReturnsOnCall[len(fake.getApplicationSummaryByNameAndSpaceArgsForCall)]
	fake.getApplicationSummaryByNameAndSpaceArgsForCall = append(fake.getApplicationSummaryByNameAndSpaceArgsForCall, struct {
		appName              string
		spaceGUID            string
		withObfuscatedValues bool
	}{appName, spaceGUID, withObfuscatedValues})
	fake.recordInvocation("GetApplicationSummaryByNameAndSpace", []interface{}{appName, spaceGUID, withObfuscatedValues})
	fake.getApplicationSummaryByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationSummaryByNameAndSpaceStub != nil {
		return fake.GetApplicationSummaryByNameAndSpaceStub(appName, spaceGUID, withObfuscatedValues)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationSummaryByNameAndSpaceReturns.result1, fake.getApplicationSummaryByNameAndSpaceReturns.result2, fake.getApplicationSummaryByNameAndSpaceReturns.result3
}

func (fake *FakeAppSummaryActor) GetApplicationSummaryByNameAndSpaceCallCount() int {
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationSummaryByNameAndSpaceArgsForCall)
}

func (fake *FakeAppSummaryActor) GetApplicationSummaryByNameAndSpaceArgsForCall(i int) (string, string, bool) {
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationSummaryByNameAndSpaceArgsForCall[i].appName, fake.getApplicationSummaryByNameAndSpaceArgsForCall[i].spaceGUID, fake.getApplicationSummaryByNameAndSpaceArgsForCall[i].withObfuscatedValues
}

func (fake *FakeAppSummaryActor) GetApplicationSummaryByNameAndSpaceReturns(result1 v2v3action.ApplicationSummary, result2 v2v3action.Warnings, result3 error) {
	fake.GetApplicationSummaryByNameAndSpaceStub = nil
	fake.getApplicationSummaryByNameAndSpaceReturns = struct {
		result1 v2v3action.ApplicationSummary
		result2 v2v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAppSummaryActor) GetApplicationSummaryByNameAndSpaceReturnsOnCall(i int, result1 v2v3action.ApplicationSummary, result2 v2v3action.Warnings, result3 error) {
	fake.GetApplicationSummaryByNameAndSpaceStub = nil
	if fake.getApplicationSummaryByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationSummaryByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v2v3action.ApplicationSummary
			result2 v2v3action.Warnings
			result3 error
		})
	}
	fake.getApplicationSummaryByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v2v3action.ApplicationSummary
		result2 v2v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAppSummaryActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAppSummaryActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.AppSummaryActor = new(FakeAppSummaryActor)
