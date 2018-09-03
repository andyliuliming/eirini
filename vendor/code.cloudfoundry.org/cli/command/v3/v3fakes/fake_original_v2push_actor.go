// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/pushaction"
	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeOriginalV2PushActor struct {
	CreateAndMapDefaultApplicationRouteStub        func(orgGUID string, spaceGUID string, app v2action.Application) (pushaction.Warnings, error)
	createAndMapDefaultApplicationRouteMutex       sync.RWMutex
	createAndMapDefaultApplicationRouteArgsForCall []struct {
		orgGUID   string
		spaceGUID string
		app       v2action.Application
	}
	createAndMapDefaultApplicationRouteReturns struct {
		result1 pushaction.Warnings
		result2 error
	}
	createAndMapDefaultApplicationRouteReturnsOnCall map[int]struct {
		result1 pushaction.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOriginalV2PushActor) CreateAndMapDefaultApplicationRoute(orgGUID string, spaceGUID string, app v2action.Application) (pushaction.Warnings, error) {
	fake.createAndMapDefaultApplicationRouteMutex.Lock()
	ret, specificReturn := fake.createAndMapDefaultApplicationRouteReturnsOnCall[len(fake.createAndMapDefaultApplicationRouteArgsForCall)]
	fake.createAndMapDefaultApplicationRouteArgsForCall = append(fake.createAndMapDefaultApplicationRouteArgsForCall, struct {
		orgGUID   string
		spaceGUID string
		app       v2action.Application
	}{orgGUID, spaceGUID, app})
	fake.recordInvocation("CreateAndMapDefaultApplicationRoute", []interface{}{orgGUID, spaceGUID, app})
	fake.createAndMapDefaultApplicationRouteMutex.Unlock()
	if fake.CreateAndMapDefaultApplicationRouteStub != nil {
		return fake.CreateAndMapDefaultApplicationRouteStub(orgGUID, spaceGUID, app)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createAndMapDefaultApplicationRouteReturns.result1, fake.createAndMapDefaultApplicationRouteReturns.result2
}

func (fake *FakeOriginalV2PushActor) CreateAndMapDefaultApplicationRouteCallCount() int {
	fake.createAndMapDefaultApplicationRouteMutex.RLock()
	defer fake.createAndMapDefaultApplicationRouteMutex.RUnlock()
	return len(fake.createAndMapDefaultApplicationRouteArgsForCall)
}

func (fake *FakeOriginalV2PushActor) CreateAndMapDefaultApplicationRouteArgsForCall(i int) (string, string, v2action.Application) {
	fake.createAndMapDefaultApplicationRouteMutex.RLock()
	defer fake.createAndMapDefaultApplicationRouteMutex.RUnlock()
	return fake.createAndMapDefaultApplicationRouteArgsForCall[i].orgGUID, fake.createAndMapDefaultApplicationRouteArgsForCall[i].spaceGUID, fake.createAndMapDefaultApplicationRouteArgsForCall[i].app
}

func (fake *FakeOriginalV2PushActor) CreateAndMapDefaultApplicationRouteReturns(result1 pushaction.Warnings, result2 error) {
	fake.CreateAndMapDefaultApplicationRouteStub = nil
	fake.createAndMapDefaultApplicationRouteReturns = struct {
		result1 pushaction.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeOriginalV2PushActor) CreateAndMapDefaultApplicationRouteReturnsOnCall(i int, result1 pushaction.Warnings, result2 error) {
	fake.CreateAndMapDefaultApplicationRouteStub = nil
	if fake.createAndMapDefaultApplicationRouteReturnsOnCall == nil {
		fake.createAndMapDefaultApplicationRouteReturnsOnCall = make(map[int]struct {
			result1 pushaction.Warnings
			result2 error
		})
	}
	fake.createAndMapDefaultApplicationRouteReturnsOnCall[i] = struct {
		result1 pushaction.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeOriginalV2PushActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createAndMapDefaultApplicationRouteMutex.RLock()
	defer fake.createAndMapDefaultApplicationRouteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOriginalV2PushActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.OriginalV2PushActor = new(FakeOriginalV2PushActor)
