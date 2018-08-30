// Code generated by counterfeiter. DO NOT EDIT.
package opifakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/opi"
)

type FakeTaskDesirer struct {
	DesireStub        func(task *opi.Task) error
	desireMutex       sync.RWMutex
	desireArgsForCall []struct {
		task *opi.Task
	}
	desireReturns struct {
		result1 error
	}
	desireReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(name string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		name string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskDesirer) Desire(task *opi.Task) error {
	fake.desireMutex.Lock()
	ret, specificReturn := fake.desireReturnsOnCall[len(fake.desireArgsForCall)]
	fake.desireArgsForCall = append(fake.desireArgsForCall, struct {
		task *opi.Task
	}{task})
	fake.recordInvocation("Desire", []interface{}{task})
	fake.desireMutex.Unlock()
	if fake.DesireStub != nil {
		return fake.DesireStub(task)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.desireReturns.result1
}

func (fake *FakeTaskDesirer) DesireCallCount() int {
	fake.desireMutex.RLock()
	defer fake.desireMutex.RUnlock()
	return len(fake.desireArgsForCall)
}

func (fake *FakeTaskDesirer) DesireArgsForCall(i int) *opi.Task {
	fake.desireMutex.RLock()
	defer fake.desireMutex.RUnlock()
	return fake.desireArgsForCall[i].task
}

func (fake *FakeTaskDesirer) DesireReturns(result1 error) {
	fake.DesireStub = nil
	fake.desireReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskDesirer) DesireReturnsOnCall(i int, result1 error) {
	fake.DesireStub = nil
	if fake.desireReturnsOnCall == nil {
		fake.desireReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.desireReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskDesirer) Delete(name string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Delete", []interface{}{name})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeTaskDesirer) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeTaskDesirer) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].name
}

func (fake *FakeTaskDesirer) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskDesirer) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskDesirer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.desireMutex.RLock()
	defer fake.desireMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTaskDesirer) recordInvocation(key string, args []interface{}) {
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

var _ opi.TaskDesirer = new(FakeTaskDesirer)
