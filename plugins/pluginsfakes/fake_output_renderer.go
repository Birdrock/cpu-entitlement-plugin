// Code generated by counterfeiter. DO NOT EDIT.
package pluginsfakes

import (
	"sync"

	"code.cloudfoundry.org/cpu-entitlement-plugin/plugins"
	"code.cloudfoundry.org/cpu-entitlement-plugin/reporter"
)

type FakeOutputRenderer struct {
	ShowApplicationReportStub        func(reporter.ApplicationReport) error
	showApplicationReportMutex       sync.RWMutex
	showApplicationReportArgsForCall []struct {
		arg1 reporter.ApplicationReport
	}
	showApplicationReportReturns struct {
		result1 error
	}
	showApplicationReportReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOutputRenderer) ShowApplicationReport(arg1 reporter.ApplicationReport) error {
	fake.showApplicationReportMutex.Lock()
	ret, specificReturn := fake.showApplicationReportReturnsOnCall[len(fake.showApplicationReportArgsForCall)]
	fake.showApplicationReportArgsForCall = append(fake.showApplicationReportArgsForCall, struct {
		arg1 reporter.ApplicationReport
	}{arg1})
	fake.recordInvocation("ShowApplicationReport", []interface{}{arg1})
	fake.showApplicationReportMutex.Unlock()
	if fake.ShowApplicationReportStub != nil {
		return fake.ShowApplicationReportStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.showApplicationReportReturns
	return fakeReturns.result1
}

func (fake *FakeOutputRenderer) ShowApplicationReportCallCount() int {
	fake.showApplicationReportMutex.RLock()
	defer fake.showApplicationReportMutex.RUnlock()
	return len(fake.showApplicationReportArgsForCall)
}

func (fake *FakeOutputRenderer) ShowApplicationReportCalls(stub func(reporter.ApplicationReport) error) {
	fake.showApplicationReportMutex.Lock()
	defer fake.showApplicationReportMutex.Unlock()
	fake.ShowApplicationReportStub = stub
}

func (fake *FakeOutputRenderer) ShowApplicationReportArgsForCall(i int) reporter.ApplicationReport {
	fake.showApplicationReportMutex.RLock()
	defer fake.showApplicationReportMutex.RUnlock()
	argsForCall := fake.showApplicationReportArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOutputRenderer) ShowApplicationReportReturns(result1 error) {
	fake.showApplicationReportMutex.Lock()
	defer fake.showApplicationReportMutex.Unlock()
	fake.ShowApplicationReportStub = nil
	fake.showApplicationReportReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutputRenderer) ShowApplicationReportReturnsOnCall(i int, result1 error) {
	fake.showApplicationReportMutex.Lock()
	defer fake.showApplicationReportMutex.Unlock()
	fake.ShowApplicationReportStub = nil
	if fake.showApplicationReportReturnsOnCall == nil {
		fake.showApplicationReportReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.showApplicationReportReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutputRenderer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.showApplicationReportMutex.RLock()
	defer fake.showApplicationReportMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOutputRenderer) recordInvocation(key string, args []interface{}) {
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

var _ plugins.OutputRenderer = new(FakeOutputRenderer)
