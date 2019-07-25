// Code generated by counterfeiter. DO NOT EDIT.
package pluginfakes

import (
	"sync"

	"code.cloudfoundry.org/cpu-entitlement-plugin/metadata"
	"code.cloudfoundry.org/cpu-entitlement-plugin/plugin"
)

type FakeCFAppInfoGetter struct {
	GetCFAppInfoStub        func(string) (metadata.CFAppInfo, error)
	getCFAppInfoMutex       sync.RWMutex
	getCFAppInfoArgsForCall []struct {
		arg1 string
	}
	getCFAppInfoReturns struct {
		result1 metadata.CFAppInfo
		result2 error
	}
	getCFAppInfoReturnsOnCall map[int]struct {
		result1 metadata.CFAppInfo
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFAppInfoGetter) GetCFAppInfo(arg1 string) (metadata.CFAppInfo, error) {
	fake.getCFAppInfoMutex.Lock()
	ret, specificReturn := fake.getCFAppInfoReturnsOnCall[len(fake.getCFAppInfoArgsForCall)]
	fake.getCFAppInfoArgsForCall = append(fake.getCFAppInfoArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetCFAppInfo", []interface{}{arg1})
	fake.getCFAppInfoMutex.Unlock()
	if fake.GetCFAppInfoStub != nil {
		return fake.GetCFAppInfoStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getCFAppInfoReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFAppInfoGetter) GetCFAppInfoCallCount() int {
	fake.getCFAppInfoMutex.RLock()
	defer fake.getCFAppInfoMutex.RUnlock()
	return len(fake.getCFAppInfoArgsForCall)
}

func (fake *FakeCFAppInfoGetter) GetCFAppInfoCalls(stub func(string) (metadata.CFAppInfo, error)) {
	fake.getCFAppInfoMutex.Lock()
	defer fake.getCFAppInfoMutex.Unlock()
	fake.GetCFAppInfoStub = stub
}

func (fake *FakeCFAppInfoGetter) GetCFAppInfoArgsForCall(i int) string {
	fake.getCFAppInfoMutex.RLock()
	defer fake.getCFAppInfoMutex.RUnlock()
	argsForCall := fake.getCFAppInfoArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFAppInfoGetter) GetCFAppInfoReturns(result1 metadata.CFAppInfo, result2 error) {
	fake.getCFAppInfoMutex.Lock()
	defer fake.getCFAppInfoMutex.Unlock()
	fake.GetCFAppInfoStub = nil
	fake.getCFAppInfoReturns = struct {
		result1 metadata.CFAppInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeCFAppInfoGetter) GetCFAppInfoReturnsOnCall(i int, result1 metadata.CFAppInfo, result2 error) {
	fake.getCFAppInfoMutex.Lock()
	defer fake.getCFAppInfoMutex.Unlock()
	fake.GetCFAppInfoStub = nil
	if fake.getCFAppInfoReturnsOnCall == nil {
		fake.getCFAppInfoReturnsOnCall = make(map[int]struct {
			result1 metadata.CFAppInfo
			result2 error
		})
	}
	fake.getCFAppInfoReturnsOnCall[i] = struct {
		result1 metadata.CFAppInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeCFAppInfoGetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getCFAppInfoMutex.RLock()
	defer fake.getCFAppInfoMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCFAppInfoGetter) recordInvocation(key string, args []interface{}) {
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

var _ plugin.CFAppInfoGetter = new(FakeCFAppInfoGetter)
