// Code generated by counterfeiter. DO NOT EDIT.
package revokfakes

import (
	"cred-alert/revok"
	"sync"
)

type FakeGitBranchCloneClient struct {
	CloneStub        func(string, string) error
	cloneMutex       sync.RWMutex
	cloneArgsForCall []struct {
		arg1 string
		arg2 string
	}
	cloneReturns struct {
		result1 error
	}
	cloneReturnsOnCall map[int]struct {
		result1 error
	}
	BranchTargetsStub        func(repoPath string) (map[string]string, error)
	branchTargetsMutex       sync.RWMutex
	branchTargetsArgsForCall []struct {
		repoPath string
	}
	branchTargetsReturns struct {
		result1 map[string]string
		result2 error
	}
	branchTargetsReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGitBranchCloneClient) Clone(arg1 string, arg2 string) error {
	fake.cloneMutex.Lock()
	ret, specificReturn := fake.cloneReturnsOnCall[len(fake.cloneArgsForCall)]
	fake.cloneArgsForCall = append(fake.cloneArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Clone", []interface{}{arg1, arg2})
	fake.cloneMutex.Unlock()
	if fake.CloneStub != nil {
		return fake.CloneStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloneReturns.result1
}

func (fake *FakeGitBranchCloneClient) CloneCallCount() int {
	fake.cloneMutex.RLock()
	defer fake.cloneMutex.RUnlock()
	return len(fake.cloneArgsForCall)
}

func (fake *FakeGitBranchCloneClient) CloneArgsForCall(i int) (string, string) {
	fake.cloneMutex.RLock()
	defer fake.cloneMutex.RUnlock()
	return fake.cloneArgsForCall[i].arg1, fake.cloneArgsForCall[i].arg2
}

func (fake *FakeGitBranchCloneClient) CloneReturns(result1 error) {
	fake.CloneStub = nil
	fake.cloneReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitBranchCloneClient) CloneReturnsOnCall(i int, result1 error) {
	fake.CloneStub = nil
	if fake.cloneReturnsOnCall == nil {
		fake.cloneReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cloneReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitBranchCloneClient) BranchTargets(repoPath string) (map[string]string, error) {
	fake.branchTargetsMutex.Lock()
	ret, specificReturn := fake.branchTargetsReturnsOnCall[len(fake.branchTargetsArgsForCall)]
	fake.branchTargetsArgsForCall = append(fake.branchTargetsArgsForCall, struct {
		repoPath string
	}{repoPath})
	fake.recordInvocation("BranchTargets", []interface{}{repoPath})
	fake.branchTargetsMutex.Unlock()
	if fake.BranchTargetsStub != nil {
		return fake.BranchTargetsStub(repoPath)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.branchTargetsReturns.result1, fake.branchTargetsReturns.result2
}

func (fake *FakeGitBranchCloneClient) BranchTargetsCallCount() int {
	fake.branchTargetsMutex.RLock()
	defer fake.branchTargetsMutex.RUnlock()
	return len(fake.branchTargetsArgsForCall)
}

func (fake *FakeGitBranchCloneClient) BranchTargetsArgsForCall(i int) string {
	fake.branchTargetsMutex.RLock()
	defer fake.branchTargetsMutex.RUnlock()
	return fake.branchTargetsArgsForCall[i].repoPath
}

func (fake *FakeGitBranchCloneClient) BranchTargetsReturns(result1 map[string]string, result2 error) {
	fake.BranchTargetsStub = nil
	fake.branchTargetsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeGitBranchCloneClient) BranchTargetsReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.BranchTargetsStub = nil
	if fake.branchTargetsReturnsOnCall == nil {
		fake.branchTargetsReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.branchTargetsReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeGitBranchCloneClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloneMutex.RLock()
	defer fake.cloneMutex.RUnlock()
	fake.branchTargetsMutex.RLock()
	defer fake.branchTargetsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGitBranchCloneClient) recordInvocation(key string, args []interface{}) {
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

var _ revok.GitBranchCloneClient = new(FakeGitBranchCloneClient)
