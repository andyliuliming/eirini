// Code generated by counterfeiter. DO NOT EDIT.
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/command/v2"
)

type FakeDownloader struct {
	DownloadStub        func(string) (string, error)
	downloadMutex       sync.RWMutex
	downloadArgsForCall []struct {
		arg1 string
	}
	downloadReturns struct {
		result1 string
		result2 error
	}
	downloadReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDownloader) Download(arg1 string) (string, error) {
	fake.downloadMutex.Lock()
	ret, specificReturn := fake.downloadReturnsOnCall[len(fake.downloadArgsForCall)]
	fake.downloadArgsForCall = append(fake.downloadArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Download", []interface{}{arg1})
	fake.downloadMutex.Unlock()
	if fake.DownloadStub != nil {
		return fake.DownloadStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.downloadReturns.result1, fake.downloadReturns.result2
}

func (fake *FakeDownloader) DownloadCallCount() int {
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	return len(fake.downloadArgsForCall)
}

func (fake *FakeDownloader) DownloadArgsForCall(i int) string {
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	return fake.downloadArgsForCall[i].arg1
}

func (fake *FakeDownloader) DownloadReturns(result1 string, result2 error) {
	fake.DownloadStub = nil
	fake.downloadReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDownloader) DownloadReturnsOnCall(i int, result1 string, result2 error) {
	fake.DownloadStub = nil
	if fake.downloadReturnsOnCall == nil {
		fake.downloadReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.downloadReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDownloader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDownloader) recordInvocation(key string, args []interface{}) {
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

var _ v2.Downloader = new(FakeDownloader)
