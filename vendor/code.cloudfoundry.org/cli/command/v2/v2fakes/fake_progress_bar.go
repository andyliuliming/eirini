// Code generated by counterfeiter. DO NOT EDIT.
package v2fakes

import (
	"io"
	"sync"

	"code.cloudfoundry.org/cli/command/v2"
)

type FakeProgressBar struct {
	NewProgressBarWrapperStub        func(reader io.Reader, sizeOfFile int64) io.Reader
	newProgressBarWrapperMutex       sync.RWMutex
	newProgressBarWrapperArgsForCall []struct {
		reader     io.Reader
		sizeOfFile int64
	}
	newProgressBarWrapperReturns struct {
		result1 io.Reader
	}
	newProgressBarWrapperReturnsOnCall map[int]struct {
		result1 io.Reader
	}
	CompleteStub        func()
	completeMutex       sync.RWMutex
	completeArgsForCall []struct{}
	ReadyStub           func()
	readyMutex          sync.RWMutex
	readyArgsForCall    []struct{}
	invocations         map[string][][]interface{}
	invocationsMutex    sync.RWMutex
}

func (fake *FakeProgressBar) NewProgressBarWrapper(reader io.Reader, sizeOfFile int64) io.Reader {
	fake.newProgressBarWrapperMutex.Lock()
	ret, specificReturn := fake.newProgressBarWrapperReturnsOnCall[len(fake.newProgressBarWrapperArgsForCall)]
	fake.newProgressBarWrapperArgsForCall = append(fake.newProgressBarWrapperArgsForCall, struct {
		reader     io.Reader
		sizeOfFile int64
	}{reader, sizeOfFile})
	fake.recordInvocation("NewProgressBarWrapper", []interface{}{reader, sizeOfFile})
	fake.newProgressBarWrapperMutex.Unlock()
	if fake.NewProgressBarWrapperStub != nil {
		return fake.NewProgressBarWrapperStub(reader, sizeOfFile)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newProgressBarWrapperReturns.result1
}

func (fake *FakeProgressBar) NewProgressBarWrapperCallCount() int {
	fake.newProgressBarWrapperMutex.RLock()
	defer fake.newProgressBarWrapperMutex.RUnlock()
	return len(fake.newProgressBarWrapperArgsForCall)
}

func (fake *FakeProgressBar) NewProgressBarWrapperArgsForCall(i int) (io.Reader, int64) {
	fake.newProgressBarWrapperMutex.RLock()
	defer fake.newProgressBarWrapperMutex.RUnlock()
	return fake.newProgressBarWrapperArgsForCall[i].reader, fake.newProgressBarWrapperArgsForCall[i].sizeOfFile
}

func (fake *FakeProgressBar) NewProgressBarWrapperReturns(result1 io.Reader) {
	fake.NewProgressBarWrapperStub = nil
	fake.newProgressBarWrapperReturns = struct {
		result1 io.Reader
	}{result1}
}

func (fake *FakeProgressBar) NewProgressBarWrapperReturnsOnCall(i int, result1 io.Reader) {
	fake.NewProgressBarWrapperStub = nil
	if fake.newProgressBarWrapperReturnsOnCall == nil {
		fake.newProgressBarWrapperReturnsOnCall = make(map[int]struct {
			result1 io.Reader
		})
	}
	fake.newProgressBarWrapperReturnsOnCall[i] = struct {
		result1 io.Reader
	}{result1}
}

func (fake *FakeProgressBar) Complete() {
	fake.completeMutex.Lock()
	fake.completeArgsForCall = append(fake.completeArgsForCall, struct{}{})
	fake.recordInvocation("Complete", []interface{}{})
	fake.completeMutex.Unlock()
	if fake.CompleteStub != nil {
		fake.CompleteStub()
	}
}

func (fake *FakeProgressBar) CompleteCallCount() int {
	fake.completeMutex.RLock()
	defer fake.completeMutex.RUnlock()
	return len(fake.completeArgsForCall)
}

func (fake *FakeProgressBar) Ready() {
	fake.readyMutex.Lock()
	fake.readyArgsForCall = append(fake.readyArgsForCall, struct{}{})
	fake.recordInvocation("Ready", []interface{}{})
	fake.readyMutex.Unlock()
	if fake.ReadyStub != nil {
		fake.ReadyStub()
	}
}

func (fake *FakeProgressBar) ReadyCallCount() int {
	fake.readyMutex.RLock()
	defer fake.readyMutex.RUnlock()
	return len(fake.readyArgsForCall)
}

func (fake *FakeProgressBar) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newProgressBarWrapperMutex.RLock()
	defer fake.newProgressBarWrapperMutex.RUnlock()
	fake.completeMutex.RLock()
	defer fake.completeMutex.RUnlock()
	fake.readyMutex.RLock()
	defer fake.readyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProgressBar) recordInvocation(key string, args []interface{}) {
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

var _ v2.ProgressBar = new(FakeProgressBar)
