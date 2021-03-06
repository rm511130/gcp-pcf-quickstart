// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type ProductUnstager struct {
	UnstageStub        func(api.UnstageProductInput) error
	unstageMutex       sync.RWMutex
	unstageArgsForCall []struct {
		arg1 api.UnstageProductInput
	}
	unstageReturns struct {
		result1 error
	}
	unstageReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ProductUnstager) Unstage(arg1 api.UnstageProductInput) error {
	fake.unstageMutex.Lock()
	ret, specificReturn := fake.unstageReturnsOnCall[len(fake.unstageArgsForCall)]
	fake.unstageArgsForCall = append(fake.unstageArgsForCall, struct {
		arg1 api.UnstageProductInput
	}{arg1})
	fake.recordInvocation("Unstage", []interface{}{arg1})
	fake.unstageMutex.Unlock()
	if fake.UnstageStub != nil {
		return fake.UnstageStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unstageReturns.result1
}

func (fake *ProductUnstager) UnstageCallCount() int {
	fake.unstageMutex.RLock()
	defer fake.unstageMutex.RUnlock()
	return len(fake.unstageArgsForCall)
}

func (fake *ProductUnstager) UnstageArgsForCall(i int) api.UnstageProductInput {
	fake.unstageMutex.RLock()
	defer fake.unstageMutex.RUnlock()
	return fake.unstageArgsForCall[i].arg1
}

func (fake *ProductUnstager) UnstageReturns(result1 error) {
	fake.UnstageStub = nil
	fake.unstageReturns = struct {
		result1 error
	}{result1}
}

func (fake *ProductUnstager) UnstageReturnsOnCall(i int, result1 error) {
	fake.UnstageStub = nil
	if fake.unstageReturnsOnCall == nil {
		fake.unstageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unstageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ProductUnstager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.unstageMutex.RLock()
	defer fake.unstageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ProductUnstager) recordInvocation(key string, args []interface{}) {
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
