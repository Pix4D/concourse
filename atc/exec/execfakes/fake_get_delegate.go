// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	"context"
	"io"
	"sync"
	"time"

	lager "code.cloudfoundry.org/lager/v3"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/exec"
	"github.com/concourse/concourse/atc/resource"
	"github.com/concourse/concourse/atc/runtime"
	"github.com/concourse/concourse/tracing"
	"go.opentelemetry.io/otel/trace"
)

type FakeGetDelegate struct {
	BeforeSelectWorkerStub        func(lager.Logger) error
	beforeSelectWorkerMutex       sync.RWMutex
	beforeSelectWorkerArgsForCall []struct {
		arg1 lager.Logger
	}
	beforeSelectWorkerReturns struct {
		result1 error
	}
	beforeSelectWorkerReturnsOnCall map[int]struct {
		result1 error
	}
	BuildStartTimeStub        func() time.Time
	buildStartTimeMutex       sync.RWMutex
	buildStartTimeArgsForCall []struct {
	}
	buildStartTimeReturns struct {
		result1 time.Time
	}
	buildStartTimeReturnsOnCall map[int]struct {
		result1 time.Time
	}
	ContainerOwnerStub        func(atc.PlanID) db.ContainerOwner
	containerOwnerMutex       sync.RWMutex
	containerOwnerArgsForCall []struct {
		arg1 atc.PlanID
	}
	containerOwnerReturns struct {
		result1 db.ContainerOwner
	}
	containerOwnerReturnsOnCall map[int]struct {
		result1 db.ContainerOwner
	}
	ErroredStub        func(lager.Logger, string)
	erroredMutex       sync.RWMutex
	erroredArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	FetchImageStub        func(context.Context, atc.Plan, *atc.Plan, bool) (runtime.ImageSpec, db.ResourceCache, error)
	fetchImageMutex       sync.RWMutex
	fetchImageArgsForCall []struct {
		arg1 context.Context
		arg2 atc.Plan
		arg3 *atc.Plan
		arg4 bool
	}
	fetchImageReturns struct {
		result1 runtime.ImageSpec
		result2 db.ResourceCache
		result3 error
	}
	fetchImageReturnsOnCall map[int]struct {
		result1 runtime.ImageSpec
		result2 db.ResourceCache
		result3 error
	}
	FinishedStub        func(lager.Logger, exec.ExitStatus, resource.VersionResult)
	finishedMutex       sync.RWMutex
	finishedArgsForCall []struct {
		arg1 lager.Logger
		arg2 exec.ExitStatus
		arg3 resource.VersionResult
	}
	InitializingStub        func(lager.Logger)
	initializingMutex       sync.RWMutex
	initializingArgsForCall []struct {
		arg1 lager.Logger
	}
	ResourceCacheUserStub        func() db.ResourceCacheUser
	resourceCacheUserMutex       sync.RWMutex
	resourceCacheUserArgsForCall []struct {
	}
	resourceCacheUserReturns struct {
		result1 db.ResourceCacheUser
	}
	resourceCacheUserReturnsOnCall map[int]struct {
		result1 db.ResourceCacheUser
	}
	SelectedWorkerStub        func(lager.Logger, string)
	selectedWorkerMutex       sync.RWMutex
	selectedWorkerArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	StartSpanStub        func(context.Context, string, tracing.Attrs) (context.Context, trace.Span)
	startSpanMutex       sync.RWMutex
	startSpanArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 tracing.Attrs
	}
	startSpanReturns struct {
		result1 context.Context
		result2 trace.Span
	}
	startSpanReturnsOnCall map[int]struct {
		result1 context.Context
		result2 trace.Span
	}
	StartingStub        func(lager.Logger)
	startingMutex       sync.RWMutex
	startingArgsForCall []struct {
		arg1 lager.Logger
	}
	StderrStub        func() io.Writer
	stderrMutex       sync.RWMutex
	stderrArgsForCall []struct {
	}
	stderrReturns struct {
		result1 io.Writer
	}
	stderrReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	StdoutStub        func() io.Writer
	stdoutMutex       sync.RWMutex
	stdoutArgsForCall []struct {
	}
	stdoutReturns struct {
		result1 io.Writer
	}
	stdoutReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	StreamingVolumeStub        func(lager.Logger, string, string, string)
	streamingVolumeMutex       sync.RWMutex
	streamingVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 string
	}
	UpdateResourceVersionStub        func(lager.Logger, string, resource.VersionResult)
	updateResourceVersionMutex       sync.RWMutex
	updateResourceVersionArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 resource.VersionResult
	}
	WaitingForStreamedVolumeStub        func(lager.Logger, string, string)
	waitingForStreamedVolumeMutex       sync.RWMutex
	waitingForStreamedVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}
	WaitingForWorkerStub        func(lager.Logger)
	waitingForWorkerMutex       sync.RWMutex
	waitingForWorkerArgsForCall []struct {
		arg1 lager.Logger
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGetDelegate) BeforeSelectWorker(arg1 lager.Logger) error {
	fake.beforeSelectWorkerMutex.Lock()
	ret, specificReturn := fake.beforeSelectWorkerReturnsOnCall[len(fake.beforeSelectWorkerArgsForCall)]
	fake.beforeSelectWorkerArgsForCall = append(fake.beforeSelectWorkerArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.BeforeSelectWorkerStub
	fakeReturns := fake.beforeSelectWorkerReturns
	fake.recordInvocation("BeforeSelectWorker", []interface{}{arg1})
	fake.beforeSelectWorkerMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGetDelegate) BeforeSelectWorkerCallCount() int {
	fake.beforeSelectWorkerMutex.RLock()
	defer fake.beforeSelectWorkerMutex.RUnlock()
	return len(fake.beforeSelectWorkerArgsForCall)
}

func (fake *FakeGetDelegate) BeforeSelectWorkerCalls(stub func(lager.Logger) error) {
	fake.beforeSelectWorkerMutex.Lock()
	defer fake.beforeSelectWorkerMutex.Unlock()
	fake.BeforeSelectWorkerStub = stub
}

func (fake *FakeGetDelegate) BeforeSelectWorkerArgsForCall(i int) lager.Logger {
	fake.beforeSelectWorkerMutex.RLock()
	defer fake.beforeSelectWorkerMutex.RUnlock()
	argsForCall := fake.beforeSelectWorkerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGetDelegate) BeforeSelectWorkerReturns(result1 error) {
	fake.beforeSelectWorkerMutex.Lock()
	defer fake.beforeSelectWorkerMutex.Unlock()
	fake.BeforeSelectWorkerStub = nil
	fake.beforeSelectWorkerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGetDelegate) BeforeSelectWorkerReturnsOnCall(i int, result1 error) {
	fake.beforeSelectWorkerMutex.Lock()
	defer fake.beforeSelectWorkerMutex.Unlock()
	fake.BeforeSelectWorkerStub = nil
	if fake.beforeSelectWorkerReturnsOnCall == nil {
		fake.beforeSelectWorkerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.beforeSelectWorkerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGetDelegate) BuildStartTime() time.Time {
	fake.buildStartTimeMutex.Lock()
	ret, specificReturn := fake.buildStartTimeReturnsOnCall[len(fake.buildStartTimeArgsForCall)]
	fake.buildStartTimeArgsForCall = append(fake.buildStartTimeArgsForCall, struct {
	}{})
	stub := fake.BuildStartTimeStub
	fakeReturns := fake.buildStartTimeReturns
	fake.recordInvocation("BuildStartTime", []interface{}{})
	fake.buildStartTimeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGetDelegate) BuildStartTimeCallCount() int {
	fake.buildStartTimeMutex.RLock()
	defer fake.buildStartTimeMutex.RUnlock()
	return len(fake.buildStartTimeArgsForCall)
}

func (fake *FakeGetDelegate) BuildStartTimeCalls(stub func() time.Time) {
	fake.buildStartTimeMutex.Lock()
	defer fake.buildStartTimeMutex.Unlock()
	fake.BuildStartTimeStub = stub
}

func (fake *FakeGetDelegate) BuildStartTimeReturns(result1 time.Time) {
	fake.buildStartTimeMutex.Lock()
	defer fake.buildStartTimeMutex.Unlock()
	fake.BuildStartTimeStub = nil
	fake.buildStartTimeReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeGetDelegate) BuildStartTimeReturnsOnCall(i int, result1 time.Time) {
	fake.buildStartTimeMutex.Lock()
	defer fake.buildStartTimeMutex.Unlock()
	fake.BuildStartTimeStub = nil
	if fake.buildStartTimeReturnsOnCall == nil {
		fake.buildStartTimeReturnsOnCall = make(map[int]struct {
			result1 time.Time
		})
	}
	fake.buildStartTimeReturnsOnCall[i] = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeGetDelegate) ContainerOwner(arg1 atc.PlanID) db.ContainerOwner {
	fake.containerOwnerMutex.Lock()
	ret, specificReturn := fake.containerOwnerReturnsOnCall[len(fake.containerOwnerArgsForCall)]
	fake.containerOwnerArgsForCall = append(fake.containerOwnerArgsForCall, struct {
		arg1 atc.PlanID
	}{arg1})
	stub := fake.ContainerOwnerStub
	fakeReturns := fake.containerOwnerReturns
	fake.recordInvocation("ContainerOwner", []interface{}{arg1})
	fake.containerOwnerMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGetDelegate) ContainerOwnerCallCount() int {
	fake.containerOwnerMutex.RLock()
	defer fake.containerOwnerMutex.RUnlock()
	return len(fake.containerOwnerArgsForCall)
}

func (fake *FakeGetDelegate) ContainerOwnerCalls(stub func(atc.PlanID) db.ContainerOwner) {
	fake.containerOwnerMutex.Lock()
	defer fake.containerOwnerMutex.Unlock()
	fake.ContainerOwnerStub = stub
}

func (fake *FakeGetDelegate) ContainerOwnerArgsForCall(i int) atc.PlanID {
	fake.containerOwnerMutex.RLock()
	defer fake.containerOwnerMutex.RUnlock()
	argsForCall := fake.containerOwnerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGetDelegate) ContainerOwnerReturns(result1 db.ContainerOwner) {
	fake.containerOwnerMutex.Lock()
	defer fake.containerOwnerMutex.Unlock()
	fake.ContainerOwnerStub = nil
	fake.containerOwnerReturns = struct {
		result1 db.ContainerOwner
	}{result1}
}

func (fake *FakeGetDelegate) ContainerOwnerReturnsOnCall(i int, result1 db.ContainerOwner) {
	fake.containerOwnerMutex.Lock()
	defer fake.containerOwnerMutex.Unlock()
	fake.ContainerOwnerStub = nil
	if fake.containerOwnerReturnsOnCall == nil {
		fake.containerOwnerReturnsOnCall = make(map[int]struct {
			result1 db.ContainerOwner
		})
	}
	fake.containerOwnerReturnsOnCall[i] = struct {
		result1 db.ContainerOwner
	}{result1}
}

func (fake *FakeGetDelegate) Errored(arg1 lager.Logger, arg2 string) {
	fake.erroredMutex.Lock()
	fake.erroredArgsForCall = append(fake.erroredArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	stub := fake.ErroredStub
	fake.recordInvocation("Errored", []interface{}{arg1, arg2})
	fake.erroredMutex.Unlock()
	if stub != nil {
		fake.ErroredStub(arg1, arg2)
	}
}

func (fake *FakeGetDelegate) ErroredCallCount() int {
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	return len(fake.erroredArgsForCall)
}

func (fake *FakeGetDelegate) ErroredCalls(stub func(lager.Logger, string)) {
	fake.erroredMutex.Lock()
	defer fake.erroredMutex.Unlock()
	fake.ErroredStub = stub
}

func (fake *FakeGetDelegate) ErroredArgsForCall(i int) (lager.Logger, string) {
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	argsForCall := fake.erroredArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGetDelegate) FetchImage(arg1 context.Context, arg2 atc.Plan, arg3 *atc.Plan, arg4 bool) (runtime.ImageSpec, db.ResourceCache, error) {
	fake.fetchImageMutex.Lock()
	ret, specificReturn := fake.fetchImageReturnsOnCall[len(fake.fetchImageArgsForCall)]
	fake.fetchImageArgsForCall = append(fake.fetchImageArgsForCall, struct {
		arg1 context.Context
		arg2 atc.Plan
		arg3 *atc.Plan
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	stub := fake.FetchImageStub
	fakeReturns := fake.fetchImageReturns
	fake.recordInvocation("FetchImage", []interface{}{arg1, arg2, arg3, arg4})
	fake.fetchImageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeGetDelegate) FetchImageCallCount() int {
	fake.fetchImageMutex.RLock()
	defer fake.fetchImageMutex.RUnlock()
	return len(fake.fetchImageArgsForCall)
}

func (fake *FakeGetDelegate) FetchImageCalls(stub func(context.Context, atc.Plan, *atc.Plan, bool) (runtime.ImageSpec, db.ResourceCache, error)) {
	fake.fetchImageMutex.Lock()
	defer fake.fetchImageMutex.Unlock()
	fake.FetchImageStub = stub
}

func (fake *FakeGetDelegate) FetchImageArgsForCall(i int) (context.Context, atc.Plan, *atc.Plan, bool) {
	fake.fetchImageMutex.RLock()
	defer fake.fetchImageMutex.RUnlock()
	argsForCall := fake.fetchImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeGetDelegate) FetchImageReturns(result1 runtime.ImageSpec, result2 db.ResourceCache, result3 error) {
	fake.fetchImageMutex.Lock()
	defer fake.fetchImageMutex.Unlock()
	fake.FetchImageStub = nil
	fake.fetchImageReturns = struct {
		result1 runtime.ImageSpec
		result2 db.ResourceCache
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGetDelegate) FetchImageReturnsOnCall(i int, result1 runtime.ImageSpec, result2 db.ResourceCache, result3 error) {
	fake.fetchImageMutex.Lock()
	defer fake.fetchImageMutex.Unlock()
	fake.FetchImageStub = nil
	if fake.fetchImageReturnsOnCall == nil {
		fake.fetchImageReturnsOnCall = make(map[int]struct {
			result1 runtime.ImageSpec
			result2 db.ResourceCache
			result3 error
		})
	}
	fake.fetchImageReturnsOnCall[i] = struct {
		result1 runtime.ImageSpec
		result2 db.ResourceCache
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGetDelegate) Finished(arg1 lager.Logger, arg2 exec.ExitStatus, arg3 resource.VersionResult) {
	fake.finishedMutex.Lock()
	fake.finishedArgsForCall = append(fake.finishedArgsForCall, struct {
		arg1 lager.Logger
		arg2 exec.ExitStatus
		arg3 resource.VersionResult
	}{arg1, arg2, arg3})
	stub := fake.FinishedStub
	fake.recordInvocation("Finished", []interface{}{arg1, arg2, arg3})
	fake.finishedMutex.Unlock()
	if stub != nil {
		fake.FinishedStub(arg1, arg2, arg3)
	}
}

func (fake *FakeGetDelegate) FinishedCallCount() int {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	return len(fake.finishedArgsForCall)
}

func (fake *FakeGetDelegate) FinishedCalls(stub func(lager.Logger, exec.ExitStatus, resource.VersionResult)) {
	fake.finishedMutex.Lock()
	defer fake.finishedMutex.Unlock()
	fake.FinishedStub = stub
}

func (fake *FakeGetDelegate) FinishedArgsForCall(i int) (lager.Logger, exec.ExitStatus, resource.VersionResult) {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	argsForCall := fake.finishedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeGetDelegate) Initializing(arg1 lager.Logger) {
	fake.initializingMutex.Lock()
	fake.initializingArgsForCall = append(fake.initializingArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.InitializingStub
	fake.recordInvocation("Initializing", []interface{}{arg1})
	fake.initializingMutex.Unlock()
	if stub != nil {
		fake.InitializingStub(arg1)
	}
}

func (fake *FakeGetDelegate) InitializingCallCount() int {
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	return len(fake.initializingArgsForCall)
}

func (fake *FakeGetDelegate) InitializingCalls(stub func(lager.Logger)) {
	fake.initializingMutex.Lock()
	defer fake.initializingMutex.Unlock()
	fake.InitializingStub = stub
}

func (fake *FakeGetDelegate) InitializingArgsForCall(i int) lager.Logger {
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	argsForCall := fake.initializingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGetDelegate) ResourceCacheUser() db.ResourceCacheUser {
	fake.resourceCacheUserMutex.Lock()
	ret, specificReturn := fake.resourceCacheUserReturnsOnCall[len(fake.resourceCacheUserArgsForCall)]
	fake.resourceCacheUserArgsForCall = append(fake.resourceCacheUserArgsForCall, struct {
	}{})
	stub := fake.ResourceCacheUserStub
	fakeReturns := fake.resourceCacheUserReturns
	fake.recordInvocation("ResourceCacheUser", []interface{}{})
	fake.resourceCacheUserMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGetDelegate) ResourceCacheUserCallCount() int {
	fake.resourceCacheUserMutex.RLock()
	defer fake.resourceCacheUserMutex.RUnlock()
	return len(fake.resourceCacheUserArgsForCall)
}

func (fake *FakeGetDelegate) ResourceCacheUserCalls(stub func() db.ResourceCacheUser) {
	fake.resourceCacheUserMutex.Lock()
	defer fake.resourceCacheUserMutex.Unlock()
	fake.ResourceCacheUserStub = stub
}

func (fake *FakeGetDelegate) ResourceCacheUserReturns(result1 db.ResourceCacheUser) {
	fake.resourceCacheUserMutex.Lock()
	defer fake.resourceCacheUserMutex.Unlock()
	fake.ResourceCacheUserStub = nil
	fake.resourceCacheUserReturns = struct {
		result1 db.ResourceCacheUser
	}{result1}
}

func (fake *FakeGetDelegate) ResourceCacheUserReturnsOnCall(i int, result1 db.ResourceCacheUser) {
	fake.resourceCacheUserMutex.Lock()
	defer fake.resourceCacheUserMutex.Unlock()
	fake.ResourceCacheUserStub = nil
	if fake.resourceCacheUserReturnsOnCall == nil {
		fake.resourceCacheUserReturnsOnCall = make(map[int]struct {
			result1 db.ResourceCacheUser
		})
	}
	fake.resourceCacheUserReturnsOnCall[i] = struct {
		result1 db.ResourceCacheUser
	}{result1}
}

func (fake *FakeGetDelegate) SelectedWorker(arg1 lager.Logger, arg2 string) {
	fake.selectedWorkerMutex.Lock()
	fake.selectedWorkerArgsForCall = append(fake.selectedWorkerArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	stub := fake.SelectedWorkerStub
	fake.recordInvocation("SelectedWorker", []interface{}{arg1, arg2})
	fake.selectedWorkerMutex.Unlock()
	if stub != nil {
		fake.SelectedWorkerStub(arg1, arg2)
	}
}

func (fake *FakeGetDelegate) SelectedWorkerCallCount() int {
	fake.selectedWorkerMutex.RLock()
	defer fake.selectedWorkerMutex.RUnlock()
	return len(fake.selectedWorkerArgsForCall)
}

func (fake *FakeGetDelegate) SelectedWorkerCalls(stub func(lager.Logger, string)) {
	fake.selectedWorkerMutex.Lock()
	defer fake.selectedWorkerMutex.Unlock()
	fake.SelectedWorkerStub = stub
}

func (fake *FakeGetDelegate) SelectedWorkerArgsForCall(i int) (lager.Logger, string) {
	fake.selectedWorkerMutex.RLock()
	defer fake.selectedWorkerMutex.RUnlock()
	argsForCall := fake.selectedWorkerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGetDelegate) StartSpan(arg1 context.Context, arg2 string, arg3 tracing.Attrs) (context.Context, trace.Span) {
	fake.startSpanMutex.Lock()
	ret, specificReturn := fake.startSpanReturnsOnCall[len(fake.startSpanArgsForCall)]
	fake.startSpanArgsForCall = append(fake.startSpanArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 tracing.Attrs
	}{arg1, arg2, arg3})
	stub := fake.StartSpanStub
	fakeReturns := fake.startSpanReturns
	fake.recordInvocation("StartSpan", []interface{}{arg1, arg2, arg3})
	fake.startSpanMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGetDelegate) StartSpanCallCount() int {
	fake.startSpanMutex.RLock()
	defer fake.startSpanMutex.RUnlock()
	return len(fake.startSpanArgsForCall)
}

func (fake *FakeGetDelegate) StartSpanCalls(stub func(context.Context, string, tracing.Attrs) (context.Context, trace.Span)) {
	fake.startSpanMutex.Lock()
	defer fake.startSpanMutex.Unlock()
	fake.StartSpanStub = stub
}

func (fake *FakeGetDelegate) StartSpanArgsForCall(i int) (context.Context, string, tracing.Attrs) {
	fake.startSpanMutex.RLock()
	defer fake.startSpanMutex.RUnlock()
	argsForCall := fake.startSpanArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeGetDelegate) StartSpanReturns(result1 context.Context, result2 trace.Span) {
	fake.startSpanMutex.Lock()
	defer fake.startSpanMutex.Unlock()
	fake.StartSpanStub = nil
	fake.startSpanReturns = struct {
		result1 context.Context
		result2 trace.Span
	}{result1, result2}
}

func (fake *FakeGetDelegate) StartSpanReturnsOnCall(i int, result1 context.Context, result2 trace.Span) {
	fake.startSpanMutex.Lock()
	defer fake.startSpanMutex.Unlock()
	fake.StartSpanStub = nil
	if fake.startSpanReturnsOnCall == nil {
		fake.startSpanReturnsOnCall = make(map[int]struct {
			result1 context.Context
			result2 trace.Span
		})
	}
	fake.startSpanReturnsOnCall[i] = struct {
		result1 context.Context
		result2 trace.Span
	}{result1, result2}
}

func (fake *FakeGetDelegate) Starting(arg1 lager.Logger) {
	fake.startingMutex.Lock()
	fake.startingArgsForCall = append(fake.startingArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.StartingStub
	fake.recordInvocation("Starting", []interface{}{arg1})
	fake.startingMutex.Unlock()
	if stub != nil {
		fake.StartingStub(arg1)
	}
}

func (fake *FakeGetDelegate) StartingCallCount() int {
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	return len(fake.startingArgsForCall)
}

func (fake *FakeGetDelegate) StartingCalls(stub func(lager.Logger)) {
	fake.startingMutex.Lock()
	defer fake.startingMutex.Unlock()
	fake.StartingStub = stub
}

func (fake *FakeGetDelegate) StartingArgsForCall(i int) lager.Logger {
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	argsForCall := fake.startingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGetDelegate) Stderr() io.Writer {
	fake.stderrMutex.Lock()
	ret, specificReturn := fake.stderrReturnsOnCall[len(fake.stderrArgsForCall)]
	fake.stderrArgsForCall = append(fake.stderrArgsForCall, struct {
	}{})
	stub := fake.StderrStub
	fakeReturns := fake.stderrReturns
	fake.recordInvocation("Stderr", []interface{}{})
	fake.stderrMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGetDelegate) StderrCallCount() int {
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	return len(fake.stderrArgsForCall)
}

func (fake *FakeGetDelegate) StderrCalls(stub func() io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = stub
}

func (fake *FakeGetDelegate) StderrReturns(result1 io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = nil
	fake.stderrReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeGetDelegate) StderrReturnsOnCall(i int, result1 io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = nil
	if fake.stderrReturnsOnCall == nil {
		fake.stderrReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stderrReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeGetDelegate) Stdout() io.Writer {
	fake.stdoutMutex.Lock()
	ret, specificReturn := fake.stdoutReturnsOnCall[len(fake.stdoutArgsForCall)]
	fake.stdoutArgsForCall = append(fake.stdoutArgsForCall, struct {
	}{})
	stub := fake.StdoutStub
	fakeReturns := fake.stdoutReturns
	fake.recordInvocation("Stdout", []interface{}{})
	fake.stdoutMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGetDelegate) StdoutCallCount() int {
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	return len(fake.stdoutArgsForCall)
}

func (fake *FakeGetDelegate) StdoutCalls(stub func() io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = stub
}

func (fake *FakeGetDelegate) StdoutReturns(result1 io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	fake.stdoutReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeGetDelegate) StdoutReturnsOnCall(i int, result1 io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	if fake.stdoutReturnsOnCall == nil {
		fake.stdoutReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stdoutReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeGetDelegate) StreamingVolume(arg1 lager.Logger, arg2 string, arg3 string, arg4 string) {
	fake.streamingVolumeMutex.Lock()
	fake.streamingVolumeArgsForCall = append(fake.streamingVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.StreamingVolumeStub
	fake.recordInvocation("StreamingVolume", []interface{}{arg1, arg2, arg3, arg4})
	fake.streamingVolumeMutex.Unlock()
	if stub != nil {
		fake.StreamingVolumeStub(arg1, arg2, arg3, arg4)
	}
}

func (fake *FakeGetDelegate) StreamingVolumeCallCount() int {
	fake.streamingVolumeMutex.RLock()
	defer fake.streamingVolumeMutex.RUnlock()
	return len(fake.streamingVolumeArgsForCall)
}

func (fake *FakeGetDelegate) StreamingVolumeCalls(stub func(lager.Logger, string, string, string)) {
	fake.streamingVolumeMutex.Lock()
	defer fake.streamingVolumeMutex.Unlock()
	fake.StreamingVolumeStub = stub
}

func (fake *FakeGetDelegate) StreamingVolumeArgsForCall(i int) (lager.Logger, string, string, string) {
	fake.streamingVolumeMutex.RLock()
	defer fake.streamingVolumeMutex.RUnlock()
	argsForCall := fake.streamingVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeGetDelegate) UpdateResourceVersion(arg1 lager.Logger, arg2 string, arg3 resource.VersionResult) {
	fake.updateResourceVersionMutex.Lock()
	fake.updateResourceVersionArgsForCall = append(fake.updateResourceVersionArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 resource.VersionResult
	}{arg1, arg2, arg3})
	stub := fake.UpdateResourceVersionStub
	fake.recordInvocation("UpdateResourceVersion", []interface{}{arg1, arg2, arg3})
	fake.updateResourceVersionMutex.Unlock()
	if stub != nil {
		fake.UpdateResourceVersionStub(arg1, arg2, arg3)
	}
}

func (fake *FakeGetDelegate) UpdateResourceVersionCallCount() int {
	fake.updateResourceVersionMutex.RLock()
	defer fake.updateResourceVersionMutex.RUnlock()
	return len(fake.updateResourceVersionArgsForCall)
}

func (fake *FakeGetDelegate) UpdateResourceVersionCalls(stub func(lager.Logger, string, resource.VersionResult)) {
	fake.updateResourceVersionMutex.Lock()
	defer fake.updateResourceVersionMutex.Unlock()
	fake.UpdateResourceVersionStub = stub
}

func (fake *FakeGetDelegate) UpdateResourceVersionArgsForCall(i int) (lager.Logger, string, resource.VersionResult) {
	fake.updateResourceVersionMutex.RLock()
	defer fake.updateResourceVersionMutex.RUnlock()
	argsForCall := fake.updateResourceVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeGetDelegate) WaitingForStreamedVolume(arg1 lager.Logger, arg2 string, arg3 string) {
	fake.waitingForStreamedVolumeMutex.Lock()
	fake.waitingForStreamedVolumeArgsForCall = append(fake.waitingForStreamedVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.WaitingForStreamedVolumeStub
	fake.recordInvocation("WaitingForStreamedVolume", []interface{}{arg1, arg2, arg3})
	fake.waitingForStreamedVolumeMutex.Unlock()
	if stub != nil {
		fake.WaitingForStreamedVolumeStub(arg1, arg2, arg3)
	}
}

func (fake *FakeGetDelegate) WaitingForStreamedVolumeCallCount() int {
	fake.waitingForStreamedVolumeMutex.RLock()
	defer fake.waitingForStreamedVolumeMutex.RUnlock()
	return len(fake.waitingForStreamedVolumeArgsForCall)
}

func (fake *FakeGetDelegate) WaitingForStreamedVolumeCalls(stub func(lager.Logger, string, string)) {
	fake.waitingForStreamedVolumeMutex.Lock()
	defer fake.waitingForStreamedVolumeMutex.Unlock()
	fake.WaitingForStreamedVolumeStub = stub
}

func (fake *FakeGetDelegate) WaitingForStreamedVolumeArgsForCall(i int) (lager.Logger, string, string) {
	fake.waitingForStreamedVolumeMutex.RLock()
	defer fake.waitingForStreamedVolumeMutex.RUnlock()
	argsForCall := fake.waitingForStreamedVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeGetDelegate) WaitingForWorker(arg1 lager.Logger) {
	fake.waitingForWorkerMutex.Lock()
	fake.waitingForWorkerArgsForCall = append(fake.waitingForWorkerArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.WaitingForWorkerStub
	fake.recordInvocation("WaitingForWorker", []interface{}{arg1})
	fake.waitingForWorkerMutex.Unlock()
	if stub != nil {
		fake.WaitingForWorkerStub(arg1)
	}
}

func (fake *FakeGetDelegate) WaitingForWorkerCallCount() int {
	fake.waitingForWorkerMutex.RLock()
	defer fake.waitingForWorkerMutex.RUnlock()
	return len(fake.waitingForWorkerArgsForCall)
}

func (fake *FakeGetDelegate) WaitingForWorkerCalls(stub func(lager.Logger)) {
	fake.waitingForWorkerMutex.Lock()
	defer fake.waitingForWorkerMutex.Unlock()
	fake.WaitingForWorkerStub = stub
}

func (fake *FakeGetDelegate) WaitingForWorkerArgsForCall(i int) lager.Logger {
	fake.waitingForWorkerMutex.RLock()
	defer fake.waitingForWorkerMutex.RUnlock()
	argsForCall := fake.waitingForWorkerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGetDelegate) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.beforeSelectWorkerMutex.RLock()
	defer fake.beforeSelectWorkerMutex.RUnlock()
	fake.buildStartTimeMutex.RLock()
	defer fake.buildStartTimeMutex.RUnlock()
	fake.containerOwnerMutex.RLock()
	defer fake.containerOwnerMutex.RUnlock()
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	fake.fetchImageMutex.RLock()
	defer fake.fetchImageMutex.RUnlock()
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	fake.resourceCacheUserMutex.RLock()
	defer fake.resourceCacheUserMutex.RUnlock()
	fake.selectedWorkerMutex.RLock()
	defer fake.selectedWorkerMutex.RUnlock()
	fake.startSpanMutex.RLock()
	defer fake.startSpanMutex.RUnlock()
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	fake.streamingVolumeMutex.RLock()
	defer fake.streamingVolumeMutex.RUnlock()
	fake.updateResourceVersionMutex.RLock()
	defer fake.updateResourceVersionMutex.RUnlock()
	fake.waitingForStreamedVolumeMutex.RLock()
	defer fake.waitingForStreamedVolumeMutex.RUnlock()
	fake.waitingForWorkerMutex.RLock()
	defer fake.waitingForWorkerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGetDelegate) recordInvocation(key string, args []interface{}) {
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

var _ exec.GetDelegate = new(FakeGetDelegate)
