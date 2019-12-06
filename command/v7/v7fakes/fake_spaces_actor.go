// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeSpacesActor struct {
	GetOrganizationSpacesWithLabelSelectorStub        func(string, string) ([]v7action.Space, v7action.Warnings, error)
	getOrganizationSpacesWithLabelSelectorMutex       sync.RWMutex
	getOrganizationSpacesWithLabelSelectorArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getOrganizationSpacesWithLabelSelectorReturns struct {
		result1 []v7action.Space
		result2 v7action.Warnings
		result3 error
	}
	getOrganizationSpacesWithLabelSelectorReturnsOnCall map[int]struct {
		result1 []v7action.Space
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpacesActor) GetOrganizationSpacesWithLabelSelector(arg1 string, arg2 string) ([]v7action.Space, v7action.Warnings, error) {
	fake.getOrganizationSpacesWithLabelSelectorMutex.Lock()
	ret, specificReturn := fake.getOrganizationSpacesWithLabelSelectorReturnsOnCall[len(fake.getOrganizationSpacesWithLabelSelectorArgsForCall)]
	fake.getOrganizationSpacesWithLabelSelectorArgsForCall = append(fake.getOrganizationSpacesWithLabelSelectorArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetOrganizationSpacesWithLabelSelector", []interface{}{arg1, arg2})
	fake.getOrganizationSpacesWithLabelSelectorMutex.Unlock()
	if fake.GetOrganizationSpacesWithLabelSelectorStub != nil {
		return fake.GetOrganizationSpacesWithLabelSelectorStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationSpacesWithLabelSelectorReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSpacesActor) GetOrganizationSpacesWithLabelSelectorCallCount() int {
	fake.getOrganizationSpacesWithLabelSelectorMutex.RLock()
	defer fake.getOrganizationSpacesWithLabelSelectorMutex.RUnlock()
	return len(fake.getOrganizationSpacesWithLabelSelectorArgsForCall)
}

func (fake *FakeSpacesActor) GetOrganizationSpacesWithLabelSelectorCalls(stub func(string, string) ([]v7action.Space, v7action.Warnings, error)) {
	fake.getOrganizationSpacesWithLabelSelectorMutex.Lock()
	defer fake.getOrganizationSpacesWithLabelSelectorMutex.Unlock()
	fake.GetOrganizationSpacesWithLabelSelectorStub = stub
}

func (fake *FakeSpacesActor) GetOrganizationSpacesWithLabelSelectorArgsForCall(i int) (string, string) {
	fake.getOrganizationSpacesWithLabelSelectorMutex.RLock()
	defer fake.getOrganizationSpacesWithLabelSelectorMutex.RUnlock()
	argsForCall := fake.getOrganizationSpacesWithLabelSelectorArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSpacesActor) GetOrganizationSpacesWithLabelSelectorReturns(result1 []v7action.Space, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationSpacesWithLabelSelectorMutex.Lock()
	defer fake.getOrganizationSpacesWithLabelSelectorMutex.Unlock()
	fake.GetOrganizationSpacesWithLabelSelectorStub = nil
	fake.getOrganizationSpacesWithLabelSelectorReturns = struct {
		result1 []v7action.Space
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpacesActor) GetOrganizationSpacesWithLabelSelectorReturnsOnCall(i int, result1 []v7action.Space, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationSpacesWithLabelSelectorMutex.Lock()
	defer fake.getOrganizationSpacesWithLabelSelectorMutex.Unlock()
	fake.GetOrganizationSpacesWithLabelSelectorStub = nil
	if fake.getOrganizationSpacesWithLabelSelectorReturnsOnCall == nil {
		fake.getOrganizationSpacesWithLabelSelectorReturnsOnCall = make(map[int]struct {
			result1 []v7action.Space
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getOrganizationSpacesWithLabelSelectorReturnsOnCall[i] = struct {
		result1 []v7action.Space
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpacesActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getOrganizationSpacesWithLabelSelectorMutex.RLock()
	defer fake.getOrganizationSpacesWithLabelSelectorMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSpacesActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.SpacesActor = new(FakeSpacesActor)