// Code generated by counterfeiter. DO NOT EDIT.
package revokfakes

import (
	"cred-alert/revok"
	"cred-alert/sniff"
	"sync"

	"code.cloudfoundry.org/lager"
)

type FakeScannerSniffer struct {
	SniffStub        func(lager.Logger, sniff.Scanner, sniff.ViolationHandlerFunc) error
	sniffMutex       sync.RWMutex
	sniffArgsForCall []struct {
		arg1 lager.Logger
		arg2 sniff.Scanner
		arg3 sniff.ViolationHandlerFunc
	}
	sniffReturns struct {
		result1 error
	}
	sniffReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeScannerSniffer) Sniff(arg1 lager.Logger, arg2 sniff.Scanner, arg3 sniff.ViolationHandlerFunc) error {
	fake.sniffMutex.Lock()
	ret, specificReturn := fake.sniffReturnsOnCall[len(fake.sniffArgsForCall)]
	fake.sniffArgsForCall = append(fake.sniffArgsForCall, struct {
		arg1 lager.Logger
		arg2 sniff.Scanner
		arg3 sniff.ViolationHandlerFunc
	}{arg1, arg2, arg3})
	fake.recordInvocation("Sniff", []interface{}{arg1, arg2, arg3})
	fake.sniffMutex.Unlock()
	if fake.SniffStub != nil {
		return fake.SniffStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sniffReturns.result1
}

func (fake *FakeScannerSniffer) SniffCallCount() int {
	fake.sniffMutex.RLock()
	defer fake.sniffMutex.RUnlock()
	return len(fake.sniffArgsForCall)
}

func (fake *FakeScannerSniffer) SniffArgsForCall(i int) (lager.Logger, sniff.Scanner, sniff.ViolationHandlerFunc) {
	fake.sniffMutex.RLock()
	defer fake.sniffMutex.RUnlock()
	return fake.sniffArgsForCall[i].arg1, fake.sniffArgsForCall[i].arg2, fake.sniffArgsForCall[i].arg3
}

func (fake *FakeScannerSniffer) SniffReturns(result1 error) {
	fake.SniffStub = nil
	fake.sniffReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScannerSniffer) SniffReturnsOnCall(i int, result1 error) {
	fake.SniffStub = nil
	if fake.sniffReturnsOnCall == nil {
		fake.sniffReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sniffReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeScannerSniffer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sniffMutex.RLock()
	defer fake.sniffMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeScannerSniffer) recordInvocation(key string, args []interface{}) {
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

var _ revok.ScannerSniffer = new(FakeScannerSniffer)
