// Code generated by counterfeiter. DO NOT EDIT.
package pluginsfakes

import (
	"sync"

	"code.cloudfoundry.org/cpu-entitlement-plugin/plugins"
	"code.cloudfoundry.org/cpu-entitlement-plugin/reporter"
	"code.cloudfoundry.org/lager"
)

type FakeReporter struct {
	CreateApplicationReportStub        func(lager.Logger, string) (reporter.ApplicationReport, error)
	createApplicationReportMutex       sync.RWMutex
	createApplicationReportArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	createApplicationReportReturns struct {
		result1 reporter.ApplicationReport
		result2 error
	}
	createApplicationReportReturnsOnCall map[int]struct {
		result1 reporter.ApplicationReport
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReporter) CreateApplicationReport(arg1 lager.Logger, arg2 string) (reporter.ApplicationReport, error) {
	fake.createApplicationReportMutex.Lock()
	ret, specificReturn := fake.createApplicationReportReturnsOnCall[len(fake.createApplicationReportArgsForCall)]
	fake.createApplicationReportArgsForCall = append(fake.createApplicationReportArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	stub := fake.CreateApplicationReportStub
	fakeReturns := fake.createApplicationReportReturns
	fake.recordInvocation("CreateApplicationReport", []interface{}{arg1, arg2})
	fake.createApplicationReportMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReporter) CreateApplicationReportCallCount() int {
	fake.createApplicationReportMutex.RLock()
	defer fake.createApplicationReportMutex.RUnlock()
	return len(fake.createApplicationReportArgsForCall)
}

func (fake *FakeReporter) CreateApplicationReportCalls(stub func(lager.Logger, string) (reporter.ApplicationReport, error)) {
	fake.createApplicationReportMutex.Lock()
	defer fake.createApplicationReportMutex.Unlock()
	fake.CreateApplicationReportStub = stub
}

func (fake *FakeReporter) CreateApplicationReportArgsForCall(i int) (lager.Logger, string) {
	fake.createApplicationReportMutex.RLock()
	defer fake.createApplicationReportMutex.RUnlock()
	argsForCall := fake.createApplicationReportArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeReporter) CreateApplicationReportReturns(result1 reporter.ApplicationReport, result2 error) {
	fake.createApplicationReportMutex.Lock()
	defer fake.createApplicationReportMutex.Unlock()
	fake.CreateApplicationReportStub = nil
	fake.createApplicationReportReturns = struct {
		result1 reporter.ApplicationReport
		result2 error
	}{result1, result2}
}

func (fake *FakeReporter) CreateApplicationReportReturnsOnCall(i int, result1 reporter.ApplicationReport, result2 error) {
	fake.createApplicationReportMutex.Lock()
	defer fake.createApplicationReportMutex.Unlock()
	fake.CreateApplicationReportStub = nil
	if fake.createApplicationReportReturnsOnCall == nil {
		fake.createApplicationReportReturnsOnCall = make(map[int]struct {
			result1 reporter.ApplicationReport
			result2 error
		})
	}
	fake.createApplicationReportReturnsOnCall[i] = struct {
		result1 reporter.ApplicationReport
		result2 error
	}{result1, result2}
}

func (fake *FakeReporter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createApplicationReportMutex.RLock()
	defer fake.createApplicationReportMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReporter) recordInvocation(key string, args []interface{}) {
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

var _ plugins.Reporter = new(FakeReporter)
