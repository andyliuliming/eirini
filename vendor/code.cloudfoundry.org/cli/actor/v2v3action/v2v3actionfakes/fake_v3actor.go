// Code generated by counterfeiter. DO NOT EDIT.
package v2v3actionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2v3action"
	"code.cloudfoundry.org/cli/actor/v3action"
)

type FakeV3Actor struct {
	GetApplicationByNameAndSpaceStub        func(string, string) (v3action.Application, v3action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	GetApplicationSummaryByNameAndSpaceStub        func(appName string, spaceGUID string, withObfuscatedValues bool) (v3action.ApplicationSummary, v3action.Warnings, error)
	getApplicationSummaryByNameAndSpaceMutex       sync.RWMutex
	getApplicationSummaryByNameAndSpaceArgsForCall []struct {
		appName              string
		spaceGUID            string
		withObfuscatedValues bool
	}
	getApplicationSummaryByNameAndSpaceReturns struct {
		result1 v3action.ApplicationSummary
		result2 v3action.Warnings
		result3 error
	}
	getApplicationSummaryByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.ApplicationSummary
		result2 v3action.Warnings
		result3 error
	}
	GetOrganizationByNameStub        func(orgName string) (v3action.Organization, v3action.Warnings, error)
	getOrganizationByNameMutex       sync.RWMutex
	getOrganizationByNameArgsForCall []struct {
		orgName string
	}
	getOrganizationByNameReturns struct {
		result1 v3action.Organization
		result2 v3action.Warnings
		result3 error
	}
	getOrganizationByNameReturnsOnCall map[int]struct {
		result1 v3action.Organization
		result2 v3action.Warnings
		result3 error
	}
	ShareServiceInstanceToSpacesStub        func(serviceInstanceGUID string, spaceGUIDs []string) (v3action.RelationshipList, v3action.Warnings, error)
	shareServiceInstanceToSpacesMutex       sync.RWMutex
	shareServiceInstanceToSpacesArgsForCall []struct {
		serviceInstanceGUID string
		spaceGUIDs          []string
	}
	shareServiceInstanceToSpacesReturns struct {
		result1 v3action.RelationshipList
		result2 v3action.Warnings
		result3 error
	}
	shareServiceInstanceToSpacesReturnsOnCall map[int]struct {
		result1 v3action.RelationshipList
		result2 v3action.Warnings
		result3 error
	}
	UnshareServiceInstanceByServiceInstanceAndSpaceStub        func(serviceInstanceGUID string, spaceGUID string) (v3action.Warnings, error)
	unshareServiceInstanceByServiceInstanceAndSpaceMutex       sync.RWMutex
	unshareServiceInstanceByServiceInstanceAndSpaceArgsForCall []struct {
		serviceInstanceGUID string
		spaceGUID           string
	}
	unshareServiceInstanceByServiceInstanceAndSpaceReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	unshareServiceInstanceByServiceInstanceAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpace(arg1 string, arg2 string) (v3action.Application, v3action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{arg1, arg2})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationByNameAndSpaceReturns.result1, fake.getApplicationByNameAndSpaceReturns.result2, fake.getApplicationByNameAndSpaceReturns.result3
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationByNameAndSpaceArgsForCall[i].arg1, fake.getApplicationByNameAndSpaceArgsForCall[i].arg2
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceReturns(result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetApplicationSummaryByNameAndSpace(appName string, spaceGUID string, withObfuscatedValues bool) (v3action.ApplicationSummary, v3action.Warnings, error) {
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

func (fake *FakeV3Actor) GetApplicationSummaryByNameAndSpaceCallCount() int {
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationSummaryByNameAndSpaceArgsForCall)
}

func (fake *FakeV3Actor) GetApplicationSummaryByNameAndSpaceArgsForCall(i int) (string, string, bool) {
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationSummaryByNameAndSpaceArgsForCall[i].appName, fake.getApplicationSummaryByNameAndSpaceArgsForCall[i].spaceGUID, fake.getApplicationSummaryByNameAndSpaceArgsForCall[i].withObfuscatedValues
}

func (fake *FakeV3Actor) GetApplicationSummaryByNameAndSpaceReturns(result1 v3action.ApplicationSummary, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationSummaryByNameAndSpaceStub = nil
	fake.getApplicationSummaryByNameAndSpaceReturns = struct {
		result1 v3action.ApplicationSummary
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetApplicationSummaryByNameAndSpaceReturnsOnCall(i int, result1 v3action.ApplicationSummary, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationSummaryByNameAndSpaceStub = nil
	if fake.getApplicationSummaryByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationSummaryByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.ApplicationSummary
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getApplicationSummaryByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.ApplicationSummary
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetOrganizationByName(orgName string) (v3action.Organization, v3action.Warnings, error) {
	fake.getOrganizationByNameMutex.Lock()
	ret, specificReturn := fake.getOrganizationByNameReturnsOnCall[len(fake.getOrganizationByNameArgsForCall)]
	fake.getOrganizationByNameArgsForCall = append(fake.getOrganizationByNameArgsForCall, struct {
		orgName string
	}{orgName})
	fake.recordInvocation("GetOrganizationByName", []interface{}{orgName})
	fake.getOrganizationByNameMutex.Unlock()
	if fake.GetOrganizationByNameStub != nil {
		return fake.GetOrganizationByNameStub(orgName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getOrganizationByNameReturns.result1, fake.getOrganizationByNameReturns.result2, fake.getOrganizationByNameReturns.result3
}

func (fake *FakeV3Actor) GetOrganizationByNameCallCount() int {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return len(fake.getOrganizationByNameArgsForCall)
}

func (fake *FakeV3Actor) GetOrganizationByNameArgsForCall(i int) string {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return fake.getOrganizationByNameArgsForCall[i].orgName
}

func (fake *FakeV3Actor) GetOrganizationByNameReturns(result1 v3action.Organization, result2 v3action.Warnings, result3 error) {
	fake.GetOrganizationByNameStub = nil
	fake.getOrganizationByNameReturns = struct {
		result1 v3action.Organization
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetOrganizationByNameReturnsOnCall(i int, result1 v3action.Organization, result2 v3action.Warnings, result3 error) {
	fake.GetOrganizationByNameStub = nil
	if fake.getOrganizationByNameReturnsOnCall == nil {
		fake.getOrganizationByNameReturnsOnCall = make(map[int]struct {
			result1 v3action.Organization
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getOrganizationByNameReturnsOnCall[i] = struct {
		result1 v3action.Organization
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) ShareServiceInstanceToSpaces(serviceInstanceGUID string, spaceGUIDs []string) (v3action.RelationshipList, v3action.Warnings, error) {
	var spaceGUIDsCopy []string
	if spaceGUIDs != nil {
		spaceGUIDsCopy = make([]string, len(spaceGUIDs))
		copy(spaceGUIDsCopy, spaceGUIDs)
	}
	fake.shareServiceInstanceToSpacesMutex.Lock()
	ret, specificReturn := fake.shareServiceInstanceToSpacesReturnsOnCall[len(fake.shareServiceInstanceToSpacesArgsForCall)]
	fake.shareServiceInstanceToSpacesArgsForCall = append(fake.shareServiceInstanceToSpacesArgsForCall, struct {
		serviceInstanceGUID string
		spaceGUIDs          []string
	}{serviceInstanceGUID, spaceGUIDsCopy})
	fake.recordInvocation("ShareServiceInstanceToSpaces", []interface{}{serviceInstanceGUID, spaceGUIDsCopy})
	fake.shareServiceInstanceToSpacesMutex.Unlock()
	if fake.ShareServiceInstanceToSpacesStub != nil {
		return fake.ShareServiceInstanceToSpacesStub(serviceInstanceGUID, spaceGUIDs)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.shareServiceInstanceToSpacesReturns.result1, fake.shareServiceInstanceToSpacesReturns.result2, fake.shareServiceInstanceToSpacesReturns.result3
}

func (fake *FakeV3Actor) ShareServiceInstanceToSpacesCallCount() int {
	fake.shareServiceInstanceToSpacesMutex.RLock()
	defer fake.shareServiceInstanceToSpacesMutex.RUnlock()
	return len(fake.shareServiceInstanceToSpacesArgsForCall)
}

func (fake *FakeV3Actor) ShareServiceInstanceToSpacesArgsForCall(i int) (string, []string) {
	fake.shareServiceInstanceToSpacesMutex.RLock()
	defer fake.shareServiceInstanceToSpacesMutex.RUnlock()
	return fake.shareServiceInstanceToSpacesArgsForCall[i].serviceInstanceGUID, fake.shareServiceInstanceToSpacesArgsForCall[i].spaceGUIDs
}

func (fake *FakeV3Actor) ShareServiceInstanceToSpacesReturns(result1 v3action.RelationshipList, result2 v3action.Warnings, result3 error) {
	fake.ShareServiceInstanceToSpacesStub = nil
	fake.shareServiceInstanceToSpacesReturns = struct {
		result1 v3action.RelationshipList
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) ShareServiceInstanceToSpacesReturnsOnCall(i int, result1 v3action.RelationshipList, result2 v3action.Warnings, result3 error) {
	fake.ShareServiceInstanceToSpacesStub = nil
	if fake.shareServiceInstanceToSpacesReturnsOnCall == nil {
		fake.shareServiceInstanceToSpacesReturnsOnCall = make(map[int]struct {
			result1 v3action.RelationshipList
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.shareServiceInstanceToSpacesReturnsOnCall[i] = struct {
		result1 v3action.RelationshipList
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) UnshareServiceInstanceByServiceInstanceAndSpace(serviceInstanceGUID string, spaceGUID string) (v3action.Warnings, error) {
	fake.unshareServiceInstanceByServiceInstanceAndSpaceMutex.Lock()
	ret, specificReturn := fake.unshareServiceInstanceByServiceInstanceAndSpaceReturnsOnCall[len(fake.unshareServiceInstanceByServiceInstanceAndSpaceArgsForCall)]
	fake.unshareServiceInstanceByServiceInstanceAndSpaceArgsForCall = append(fake.unshareServiceInstanceByServiceInstanceAndSpaceArgsForCall, struct {
		serviceInstanceGUID string
		spaceGUID           string
	}{serviceInstanceGUID, spaceGUID})
	fake.recordInvocation("UnshareServiceInstanceByServiceInstanceAndSpace", []interface{}{serviceInstanceGUID, spaceGUID})
	fake.unshareServiceInstanceByServiceInstanceAndSpaceMutex.Unlock()
	if fake.UnshareServiceInstanceByServiceInstanceAndSpaceStub != nil {
		return fake.UnshareServiceInstanceByServiceInstanceAndSpaceStub(serviceInstanceGUID, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.unshareServiceInstanceByServiceInstanceAndSpaceReturns.result1, fake.unshareServiceInstanceByServiceInstanceAndSpaceReturns.result2
}

func (fake *FakeV3Actor) UnshareServiceInstanceByServiceInstanceAndSpaceCallCount() int {
	fake.unshareServiceInstanceByServiceInstanceAndSpaceMutex.RLock()
	defer fake.unshareServiceInstanceByServiceInstanceAndSpaceMutex.RUnlock()
	return len(fake.unshareServiceInstanceByServiceInstanceAndSpaceArgsForCall)
}

func (fake *FakeV3Actor) UnshareServiceInstanceByServiceInstanceAndSpaceArgsForCall(i int) (string, string) {
	fake.unshareServiceInstanceByServiceInstanceAndSpaceMutex.RLock()
	defer fake.unshareServiceInstanceByServiceInstanceAndSpaceMutex.RUnlock()
	return fake.unshareServiceInstanceByServiceInstanceAndSpaceArgsForCall[i].serviceInstanceGUID, fake.unshareServiceInstanceByServiceInstanceAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeV3Actor) UnshareServiceInstanceByServiceInstanceAndSpaceReturns(result1 v3action.Warnings, result2 error) {
	fake.UnshareServiceInstanceByServiceInstanceAndSpaceStub = nil
	fake.unshareServiceInstanceByServiceInstanceAndSpaceReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3Actor) UnshareServiceInstanceByServiceInstanceAndSpaceReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.UnshareServiceInstanceByServiceInstanceAndSpaceStub = nil
	if fake.unshareServiceInstanceByServiceInstanceAndSpaceReturnsOnCall == nil {
		fake.unshareServiceInstanceByServiceInstanceAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.unshareServiceInstanceByServiceInstanceAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3Actor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloudControllerAPIVersionReturns.result1
}

func (fake *FakeV3Actor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeV3Actor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3Actor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3Actor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	fake.shareServiceInstanceToSpacesMutex.RLock()
	defer fake.shareServiceInstanceToSpacesMutex.RUnlock()
	fake.unshareServiceInstanceByServiceInstanceAndSpaceMutex.RLock()
	defer fake.unshareServiceInstanceByServiceInstanceAndSpaceMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3Actor) recordInvocation(key string, args []interface{}) {
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

var _ v2v3action.V3Actor = new(FakeV3Actor)