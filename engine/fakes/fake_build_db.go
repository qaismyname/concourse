// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/engine"
	"github.com/pivotal-golang/lager"
)

type FakeBuildDB struct {
	GetBuildStub        func(int) (db.Build, bool, error)
	getBuildMutex       sync.RWMutex
	getBuildArgsForCall []struct {
		arg1 int
	}
	getBuildReturns struct {
		result1 db.Build
		result2 bool
		result3 error
	}
	StartBuildStub        func(int, string, string) (bool, error)
	startBuildMutex       sync.RWMutex
	startBuildArgsForCall []struct {
		arg1 int
		arg2 string
		arg3 string
	}
	startBuildReturns struct {
		result1 bool
		result2 error
	}
	AbortBuildStub        func(int) error
	abortBuildMutex       sync.RWMutex
	abortBuildArgsForCall []struct {
		arg1 int
	}
	abortBuildReturns struct {
		result1 error
	}
	AbortNotifierStub        func(int) (db.Notifier, error)
	abortNotifierMutex       sync.RWMutex
	abortNotifierArgsForCall []struct {
		arg1 int
	}
	abortNotifierReturns struct {
		result1 db.Notifier
		result2 error
	}
	LeaseBuildTrackingStub        func(logger lager.Logger, buildID int, interval time.Duration) (db.Lease, bool, error)
	leaseBuildTrackingMutex       sync.RWMutex
	leaseBuildTrackingArgsForCall []struct {
		logger   lager.Logger
		buildID  int
		interval time.Duration
	}
	leaseBuildTrackingReturns struct {
		result1 db.Lease
		result2 bool
		result3 error
	}
	FinishBuildStub        func(int, db.Status) error
	finishBuildMutex       sync.RWMutex
	finishBuildArgsForCall []struct {
		arg1 int
		arg2 db.Status
	}
	finishBuildReturns struct {
		result1 error
	}
}

func (fake *FakeBuildDB) GetBuild(arg1 int) (db.Build, bool, error) {
	fake.getBuildMutex.Lock()
	fake.getBuildArgsForCall = append(fake.getBuildArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.getBuildMutex.Unlock()
	if fake.GetBuildStub != nil {
		return fake.GetBuildStub(arg1)
	} else {
		return fake.getBuildReturns.result1, fake.getBuildReturns.result2, fake.getBuildReturns.result3
	}
}

func (fake *FakeBuildDB) GetBuildCallCount() int {
	fake.getBuildMutex.RLock()
	defer fake.getBuildMutex.RUnlock()
	return len(fake.getBuildArgsForCall)
}

func (fake *FakeBuildDB) GetBuildArgsForCall(i int) int {
	fake.getBuildMutex.RLock()
	defer fake.getBuildMutex.RUnlock()
	return fake.getBuildArgsForCall[i].arg1
}

func (fake *FakeBuildDB) GetBuildReturns(result1 db.Build, result2 bool, result3 error) {
	fake.GetBuildStub = nil
	fake.getBuildReturns = struct {
		result1 db.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuildDB) StartBuild(arg1 int, arg2 string, arg3 string) (bool, error) {
	fake.startBuildMutex.Lock()
	fake.startBuildArgsForCall = append(fake.startBuildArgsForCall, struct {
		arg1 int
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.startBuildMutex.Unlock()
	if fake.StartBuildStub != nil {
		return fake.StartBuildStub(arg1, arg2, arg3)
	} else {
		return fake.startBuildReturns.result1, fake.startBuildReturns.result2
	}
}

func (fake *FakeBuildDB) StartBuildCallCount() int {
	fake.startBuildMutex.RLock()
	defer fake.startBuildMutex.RUnlock()
	return len(fake.startBuildArgsForCall)
}

func (fake *FakeBuildDB) StartBuildArgsForCall(i int) (int, string, string) {
	fake.startBuildMutex.RLock()
	defer fake.startBuildMutex.RUnlock()
	return fake.startBuildArgsForCall[i].arg1, fake.startBuildArgsForCall[i].arg2, fake.startBuildArgsForCall[i].arg3
}

func (fake *FakeBuildDB) StartBuildReturns(result1 bool, result2 error) {
	fake.StartBuildStub = nil
	fake.startBuildReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildDB) AbortBuild(arg1 int) error {
	fake.abortBuildMutex.Lock()
	fake.abortBuildArgsForCall = append(fake.abortBuildArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.abortBuildMutex.Unlock()
	if fake.AbortBuildStub != nil {
		return fake.AbortBuildStub(arg1)
	} else {
		return fake.abortBuildReturns.result1
	}
}

func (fake *FakeBuildDB) AbortBuildCallCount() int {
	fake.abortBuildMutex.RLock()
	defer fake.abortBuildMutex.RUnlock()
	return len(fake.abortBuildArgsForCall)
}

func (fake *FakeBuildDB) AbortBuildArgsForCall(i int) int {
	fake.abortBuildMutex.RLock()
	defer fake.abortBuildMutex.RUnlock()
	return fake.abortBuildArgsForCall[i].arg1
}

func (fake *FakeBuildDB) AbortBuildReturns(result1 error) {
	fake.AbortBuildStub = nil
	fake.abortBuildReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildDB) AbortNotifier(arg1 int) (db.Notifier, error) {
	fake.abortNotifierMutex.Lock()
	fake.abortNotifierArgsForCall = append(fake.abortNotifierArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.abortNotifierMutex.Unlock()
	if fake.AbortNotifierStub != nil {
		return fake.AbortNotifierStub(arg1)
	} else {
		return fake.abortNotifierReturns.result1, fake.abortNotifierReturns.result2
	}
}

func (fake *FakeBuildDB) AbortNotifierCallCount() int {
	fake.abortNotifierMutex.RLock()
	defer fake.abortNotifierMutex.RUnlock()
	return len(fake.abortNotifierArgsForCall)
}

func (fake *FakeBuildDB) AbortNotifierArgsForCall(i int) int {
	fake.abortNotifierMutex.RLock()
	defer fake.abortNotifierMutex.RUnlock()
	return fake.abortNotifierArgsForCall[i].arg1
}

func (fake *FakeBuildDB) AbortNotifierReturns(result1 db.Notifier, result2 error) {
	fake.AbortNotifierStub = nil
	fake.abortNotifierReturns = struct {
		result1 db.Notifier
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildDB) LeaseBuildTracking(logger lager.Logger, buildID int, interval time.Duration) (db.Lease, bool, error) {
	fake.leaseBuildTrackingMutex.Lock()
	fake.leaseBuildTrackingArgsForCall = append(fake.leaseBuildTrackingArgsForCall, struct {
		logger   lager.Logger
		buildID  int
		interval time.Duration
	}{logger, buildID, interval})
	fake.leaseBuildTrackingMutex.Unlock()
	if fake.LeaseBuildTrackingStub != nil {
		return fake.LeaseBuildTrackingStub(logger, buildID, interval)
	} else {
		return fake.leaseBuildTrackingReturns.result1, fake.leaseBuildTrackingReturns.result2, fake.leaseBuildTrackingReturns.result3
	}
}

func (fake *FakeBuildDB) LeaseBuildTrackingCallCount() int {
	fake.leaseBuildTrackingMutex.RLock()
	defer fake.leaseBuildTrackingMutex.RUnlock()
	return len(fake.leaseBuildTrackingArgsForCall)
}

func (fake *FakeBuildDB) LeaseBuildTrackingArgsForCall(i int) (lager.Logger, int, time.Duration) {
	fake.leaseBuildTrackingMutex.RLock()
	defer fake.leaseBuildTrackingMutex.RUnlock()
	return fake.leaseBuildTrackingArgsForCall[i].logger, fake.leaseBuildTrackingArgsForCall[i].buildID, fake.leaseBuildTrackingArgsForCall[i].interval
}

func (fake *FakeBuildDB) LeaseBuildTrackingReturns(result1 db.Lease, result2 bool, result3 error) {
	fake.LeaseBuildTrackingStub = nil
	fake.leaseBuildTrackingReturns = struct {
		result1 db.Lease
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuildDB) FinishBuild(arg1 int, arg2 db.Status) error {
	fake.finishBuildMutex.Lock()
	fake.finishBuildArgsForCall = append(fake.finishBuildArgsForCall, struct {
		arg1 int
		arg2 db.Status
	}{arg1, arg2})
	fake.finishBuildMutex.Unlock()
	if fake.FinishBuildStub != nil {
		return fake.FinishBuildStub(arg1, arg2)
	} else {
		return fake.finishBuildReturns.result1
	}
}

func (fake *FakeBuildDB) FinishBuildCallCount() int {
	fake.finishBuildMutex.RLock()
	defer fake.finishBuildMutex.RUnlock()
	return len(fake.finishBuildArgsForCall)
}

func (fake *FakeBuildDB) FinishBuildArgsForCall(i int) (int, db.Status) {
	fake.finishBuildMutex.RLock()
	defer fake.finishBuildMutex.RUnlock()
	return fake.finishBuildArgsForCall[i].arg1, fake.finishBuildArgsForCall[i].arg2
}

func (fake *FakeBuildDB) FinishBuildReturns(result1 error) {
	fake.FinishBuildStub = nil
	fake.finishBuildReturns = struct {
		result1 error
	}{result1}
}

var _ engine.BuildDB = new(FakeBuildDB)
