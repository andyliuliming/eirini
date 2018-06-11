// Code generated by counterfeiter. DO NOT EDIT.
package eirinifakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/eirini"
	"code.cloudfoundry.org/runtimeschema/cc_messages"
)

type FakeBifrost struct {
	TransferStub        func(ctx context.Context, ccMessages []cc_messages.DesireAppRequestFromCC) error
	transferMutex       sync.RWMutex
	transferArgsForCall []struct {
		ctx        context.Context
		ccMessages []cc_messages.DesireAppRequestFromCC
	}
	transferReturns struct {
		result1 error
	}
	transferReturnsOnCall map[int]struct {
		result1 error
	}
	ListStub        func(ctx context.Context) ([]*models.DesiredLRPSchedulingInfo, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		ctx context.Context
	}
	listReturns struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBifrost) Transfer(ctx context.Context, ccMessages []cc_messages.DesireAppRequestFromCC) error {
	var ccMessagesCopy []cc_messages.DesireAppRequestFromCC
	if ccMessages != nil {
		ccMessagesCopy = make([]cc_messages.DesireAppRequestFromCC, len(ccMessages))
		copy(ccMessagesCopy, ccMessages)
	}
	fake.transferMutex.Lock()
	ret, specificReturn := fake.transferReturnsOnCall[len(fake.transferArgsForCall)]
	fake.transferArgsForCall = append(fake.transferArgsForCall, struct {
		ctx        context.Context
		ccMessages []cc_messages.DesireAppRequestFromCC
	}{ctx, ccMessagesCopy})
	fake.recordInvocation("Transfer", []interface{}{ctx, ccMessagesCopy})
	fake.transferMutex.Unlock()
	if fake.TransferStub != nil {
		return fake.TransferStub(ctx, ccMessages)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.transferReturns.result1
}

func (fake *FakeBifrost) TransferCallCount() int {
	fake.transferMutex.RLock()
	defer fake.transferMutex.RUnlock()
	return len(fake.transferArgsForCall)
}

func (fake *FakeBifrost) TransferArgsForCall(i int) (context.Context, []cc_messages.DesireAppRequestFromCC) {
	fake.transferMutex.RLock()
	defer fake.transferMutex.RUnlock()
	return fake.transferArgsForCall[i].ctx, fake.transferArgsForCall[i].ccMessages
}

func (fake *FakeBifrost) TransferReturns(result1 error) {
	fake.TransferStub = nil
	fake.transferReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBifrost) TransferReturnsOnCall(i int, result1 error) {
	fake.TransferStub = nil
	if fake.transferReturnsOnCall == nil {
		fake.transferReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.transferReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBifrost) List(ctx context.Context) ([]*models.DesiredLRPSchedulingInfo, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		ctx context.Context
	}{ctx})
	fake.recordInvocation("List", []interface{}{ctx})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(ctx)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listReturns.result1, fake.listReturns.result2
}

func (fake *FakeBifrost) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeBifrost) ListArgsForCall(i int) context.Context {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].ctx
}

func (fake *FakeBifrost) ListReturns(result1 []*models.DesiredLRPSchedulingInfo, result2 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeBifrost) ListReturnsOnCall(i int, result1 []*models.DesiredLRPSchedulingInfo, result2 error) {
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []*models.DesiredLRPSchedulingInfo
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeBifrost) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.transferMutex.RLock()
	defer fake.transferMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBifrost) recordInvocation(key string, args []interface{}) {
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

var _ eirini.Bifrost = new(FakeBifrost)