// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/db"
)

type FakeTaskQueue struct {
	DequeueStub        func(string, lager.Logger)
	dequeueMutex       sync.RWMutex
	dequeueArgsForCall []struct {
		arg1 string
		arg2 lager.Logger
	}
	FindOrAppendStub        func(string, string, int, string, lager.Logger) (int, int, error)
	findOrAppendMutex       sync.RWMutex
	findOrAppendArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 int
		arg4 string
		arg5 lager.Logger
	}
	findOrAppendReturns struct {
		result1 int
		result2 int
		result3 error
	}
	findOrAppendReturnsOnCall map[int]struct {
		result1 int
		result2 int
		result3 error
	}
	FindQueueStub        func(string) (string, int, string, error)
	findQueueMutex       sync.RWMutex
	findQueueArgsForCall []struct {
		arg1 string
	}
	findQueueReturns struct {
		result1 string
		result2 int
		result3 string
		result4 error
	}
	findQueueReturnsOnCall map[int]struct {
		result1 string
		result2 int
		result3 string
		result4 error
	}
	LengthStub        func(string) (int, error)
	lengthMutex       sync.RWMutex
	lengthArgsForCall []struct {
		arg1 string
	}
	lengthReturns struct {
		result1 int
		result2 error
	}
	lengthReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	PositionStub        func(string) (int, error)
	positionMutex       sync.RWMutex
	positionArgsForCall []struct {
		arg1 string
	}
	positionReturns struct {
		result1 int
		result2 error
	}
	positionReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskQueue) Dequeue(arg1 string, arg2 lager.Logger) {
	fake.dequeueMutex.Lock()
	fake.dequeueArgsForCall = append(fake.dequeueArgsForCall, struct {
		arg1 string
		arg2 lager.Logger
	}{arg1, arg2})
	fake.recordInvocation("Dequeue", []interface{}{arg1, arg2})
	fake.dequeueMutex.Unlock()
	if fake.DequeueStub != nil {
		fake.DequeueStub(arg1, arg2)
	}
}

func (fake *FakeTaskQueue) DequeueCallCount() int {
	fake.dequeueMutex.RLock()
	defer fake.dequeueMutex.RUnlock()
	return len(fake.dequeueArgsForCall)
}

func (fake *FakeTaskQueue) DequeueCalls(stub func(string, lager.Logger)) {
	fake.dequeueMutex.Lock()
	defer fake.dequeueMutex.Unlock()
	fake.DequeueStub = stub
}

func (fake *FakeTaskQueue) DequeueArgsForCall(i int) (string, lager.Logger) {
	fake.dequeueMutex.RLock()
	defer fake.dequeueMutex.RUnlock()
	argsForCall := fake.dequeueArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTaskQueue) FindOrAppend(arg1 string, arg2 string, arg3 int, arg4 string, arg5 lager.Logger) (int, int, error) {
	fake.findOrAppendMutex.Lock()
	ret, specificReturn := fake.findOrAppendReturnsOnCall[len(fake.findOrAppendArgsForCall)]
	fake.findOrAppendArgsForCall = append(fake.findOrAppendArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 int
		arg4 string
		arg5 lager.Logger
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("FindOrAppend", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.findOrAppendMutex.Unlock()
	if fake.FindOrAppendStub != nil {
		return fake.FindOrAppendStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.findOrAppendReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeTaskQueue) FindOrAppendCallCount() int {
	fake.findOrAppendMutex.RLock()
	defer fake.findOrAppendMutex.RUnlock()
	return len(fake.findOrAppendArgsForCall)
}

func (fake *FakeTaskQueue) FindOrAppendCalls(stub func(string, string, int, string, lager.Logger) (int, int, error)) {
	fake.findOrAppendMutex.Lock()
	defer fake.findOrAppendMutex.Unlock()
	fake.FindOrAppendStub = stub
}

func (fake *FakeTaskQueue) FindOrAppendArgsForCall(i int) (string, string, int, string, lager.Logger) {
	fake.findOrAppendMutex.RLock()
	defer fake.findOrAppendMutex.RUnlock()
	argsForCall := fake.findOrAppendArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeTaskQueue) FindOrAppendReturns(result1 int, result2 int, result3 error) {
	fake.findOrAppendMutex.Lock()
	defer fake.findOrAppendMutex.Unlock()
	fake.FindOrAppendStub = nil
	fake.findOrAppendReturns = struct {
		result1 int
		result2 int
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTaskQueue) FindOrAppendReturnsOnCall(i int, result1 int, result2 int, result3 error) {
	fake.findOrAppendMutex.Lock()
	defer fake.findOrAppendMutex.Unlock()
	fake.FindOrAppendStub = nil
	if fake.findOrAppendReturnsOnCall == nil {
		fake.findOrAppendReturnsOnCall = make(map[int]struct {
			result1 int
			result2 int
			result3 error
		})
	}
	fake.findOrAppendReturnsOnCall[i] = struct {
		result1 int
		result2 int
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTaskQueue) FindQueue(arg1 string) (string, int, string, error) {
	fake.findQueueMutex.Lock()
	ret, specificReturn := fake.findQueueReturnsOnCall[len(fake.findQueueArgsForCall)]
	fake.findQueueArgsForCall = append(fake.findQueueArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FindQueue", []interface{}{arg1})
	fake.findQueueMutex.Unlock()
	if fake.FindQueueStub != nil {
		return fake.FindQueueStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	fakeReturns := fake.findQueueReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakeTaskQueue) FindQueueCallCount() int {
	fake.findQueueMutex.RLock()
	defer fake.findQueueMutex.RUnlock()
	return len(fake.findQueueArgsForCall)
}

func (fake *FakeTaskQueue) FindQueueCalls(stub func(string) (string, int, string, error)) {
	fake.findQueueMutex.Lock()
	defer fake.findQueueMutex.Unlock()
	fake.FindQueueStub = stub
}

func (fake *FakeTaskQueue) FindQueueArgsForCall(i int) string {
	fake.findQueueMutex.RLock()
	defer fake.findQueueMutex.RUnlock()
	argsForCall := fake.findQueueArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTaskQueue) FindQueueReturns(result1 string, result2 int, result3 string, result4 error) {
	fake.findQueueMutex.Lock()
	defer fake.findQueueMutex.Unlock()
	fake.FindQueueStub = nil
	fake.findQueueReturns = struct {
		result1 string
		result2 int
		result3 string
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeTaskQueue) FindQueueReturnsOnCall(i int, result1 string, result2 int, result3 string, result4 error) {
	fake.findQueueMutex.Lock()
	defer fake.findQueueMutex.Unlock()
	fake.FindQueueStub = nil
	if fake.findQueueReturnsOnCall == nil {
		fake.findQueueReturnsOnCall = make(map[int]struct {
			result1 string
			result2 int
			result3 string
			result4 error
		})
	}
	fake.findQueueReturnsOnCall[i] = struct {
		result1 string
		result2 int
		result3 string
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeTaskQueue) Length(arg1 string) (int, error) {
	fake.lengthMutex.Lock()
	ret, specificReturn := fake.lengthReturnsOnCall[len(fake.lengthArgsForCall)]
	fake.lengthArgsForCall = append(fake.lengthArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Length", []interface{}{arg1})
	fake.lengthMutex.Unlock()
	if fake.LengthStub != nil {
		return fake.LengthStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.lengthReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTaskQueue) LengthCallCount() int {
	fake.lengthMutex.RLock()
	defer fake.lengthMutex.RUnlock()
	return len(fake.lengthArgsForCall)
}

func (fake *FakeTaskQueue) LengthCalls(stub func(string) (int, error)) {
	fake.lengthMutex.Lock()
	defer fake.lengthMutex.Unlock()
	fake.LengthStub = stub
}

func (fake *FakeTaskQueue) LengthArgsForCall(i int) string {
	fake.lengthMutex.RLock()
	defer fake.lengthMutex.RUnlock()
	argsForCall := fake.lengthArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTaskQueue) LengthReturns(result1 int, result2 error) {
	fake.lengthMutex.Lock()
	defer fake.lengthMutex.Unlock()
	fake.LengthStub = nil
	fake.lengthReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskQueue) LengthReturnsOnCall(i int, result1 int, result2 error) {
	fake.lengthMutex.Lock()
	defer fake.lengthMutex.Unlock()
	fake.LengthStub = nil
	if fake.lengthReturnsOnCall == nil {
		fake.lengthReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.lengthReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskQueue) Position(arg1 string) (int, error) {
	fake.positionMutex.Lock()
	ret, specificReturn := fake.positionReturnsOnCall[len(fake.positionArgsForCall)]
	fake.positionArgsForCall = append(fake.positionArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Position", []interface{}{arg1})
	fake.positionMutex.Unlock()
	if fake.PositionStub != nil {
		return fake.PositionStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.positionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTaskQueue) PositionCallCount() int {
	fake.positionMutex.RLock()
	defer fake.positionMutex.RUnlock()
	return len(fake.positionArgsForCall)
}

func (fake *FakeTaskQueue) PositionCalls(stub func(string) (int, error)) {
	fake.positionMutex.Lock()
	defer fake.positionMutex.Unlock()
	fake.PositionStub = stub
}

func (fake *FakeTaskQueue) PositionArgsForCall(i int) string {
	fake.positionMutex.RLock()
	defer fake.positionMutex.RUnlock()
	argsForCall := fake.positionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTaskQueue) PositionReturns(result1 int, result2 error) {
	fake.positionMutex.Lock()
	defer fake.positionMutex.Unlock()
	fake.PositionStub = nil
	fake.positionReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskQueue) PositionReturnsOnCall(i int, result1 int, result2 error) {
	fake.positionMutex.Lock()
	defer fake.positionMutex.Unlock()
	fake.PositionStub = nil
	if fake.positionReturnsOnCall == nil {
		fake.positionReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.positionReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskQueue) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dequeueMutex.RLock()
	defer fake.dequeueMutex.RUnlock()
	fake.findOrAppendMutex.RLock()
	defer fake.findOrAppendMutex.RUnlock()
	fake.findQueueMutex.RLock()
	defer fake.findQueueMutex.RUnlock()
	fake.lengthMutex.RLock()
	defer fake.lengthMutex.RUnlock()
	fake.positionMutex.RLock()
	defer fake.positionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTaskQueue) recordInvocation(key string, args []interface{}) {
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

var _ db.TaskQueue = new(FakeTaskQueue)
