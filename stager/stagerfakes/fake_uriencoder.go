// Code generated by counterfeiter. DO NOT EDIT.
package stagerfakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/stager"
)

type FakeURIEncoder struct {
	EncodeStub        func(string) string
	encodeMutex       sync.RWMutex
	encodeArgsForCall []struct {
		arg1 string
	}
	encodeReturns struct {
		result1 string
	}
	encodeReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeURIEncoder) Encode(arg1 string) string {
	fake.encodeMutex.Lock()
	ret, specificReturn := fake.encodeReturnsOnCall[len(fake.encodeArgsForCall)]
	fake.encodeArgsForCall = append(fake.encodeArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Encode", []interface{}{arg1})
	fake.encodeMutex.Unlock()
	if fake.EncodeStub != nil {
		return fake.EncodeStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.encodeReturns.result1
}

func (fake *FakeURIEncoder) EncodeCallCount() int {
	fake.encodeMutex.RLock()
	defer fake.encodeMutex.RUnlock()
	return len(fake.encodeArgsForCall)
}

func (fake *FakeURIEncoder) EncodeArgsForCall(i int) string {
	fake.encodeMutex.RLock()
	defer fake.encodeMutex.RUnlock()
	return fake.encodeArgsForCall[i].arg1
}

func (fake *FakeURIEncoder) EncodeReturns(result1 string) {
	fake.EncodeStub = nil
	fake.encodeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeURIEncoder) EncodeReturnsOnCall(i int, result1 string) {
	fake.EncodeStub = nil
	if fake.encodeReturnsOnCall == nil {
		fake.encodeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.encodeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeURIEncoder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.encodeMutex.RLock()
	defer fake.encodeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeURIEncoder) recordInvocation(key string, args []interface{}) {
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

var _ stager.URIEncoder = new(FakeURIEncoder)
