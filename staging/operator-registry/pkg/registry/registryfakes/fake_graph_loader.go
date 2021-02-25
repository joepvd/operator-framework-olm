// Code generated by counterfeiter. DO NOT EDIT.
package registryfakes

import (
	"sync"

	"github.com/operator-framework/operator-registry/pkg/registry"
)

type FakeGraphLoader struct {
	GenerateStub        func(string) (*registry.Package, error)
	generateMutex       sync.RWMutex
	generateArgsForCall []struct {
		arg1 string
	}
	generateReturns struct {
		result1 *registry.Package
		result2 error
	}
	generateReturnsOnCall map[int]struct {
		result1 *registry.Package
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGraphLoader) Generate(arg1 string) (*registry.Package, error) {
	fake.generateMutex.Lock()
	ret, specificReturn := fake.generateReturnsOnCall[len(fake.generateArgsForCall)]
	fake.generateArgsForCall = append(fake.generateArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Generate", []interface{}{arg1})
	fake.generateMutex.Unlock()
	if fake.GenerateStub != nil {
		return fake.GenerateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.generateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGraphLoader) GenerateCallCount() int {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	return len(fake.generateArgsForCall)
}

func (fake *FakeGraphLoader) GenerateCalls(stub func(string) (*registry.Package, error)) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = stub
}

func (fake *FakeGraphLoader) GenerateArgsForCall(i int) string {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	argsForCall := fake.generateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGraphLoader) GenerateReturns(result1 *registry.Package, result2 error) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = nil
	fake.generateReturns = struct {
		result1 *registry.Package
		result2 error
	}{result1, result2}
}

func (fake *FakeGraphLoader) GenerateReturnsOnCall(i int, result1 *registry.Package, result2 error) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = nil
	if fake.generateReturnsOnCall == nil {
		fake.generateReturnsOnCall = make(map[int]struct {
			result1 *registry.Package
			result2 error
		})
	}
	fake.generateReturnsOnCall[i] = struct {
		result1 *registry.Package
		result2 error
	}{result1, result2}
}

func (fake *FakeGraphLoader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGraphLoader) recordInvocation(key string, args []interface{}) {
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

var _ registry.GraphLoader = new(FakeGraphLoader)