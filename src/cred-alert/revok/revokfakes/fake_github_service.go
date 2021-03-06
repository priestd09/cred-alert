// Code generated by counterfeiter. DO NOT EDIT.
package revokfakes

import (
	"cred-alert/revok"
	"sync"

	"code.cloudfoundry.org/lager"
)

type FakeGithubService struct {
	ListRepositoriesByOrgStub        func(logger lager.Logger, orgName string) ([]revok.GitHubRepository, error)
	listRepositoriesByOrgMutex       sync.RWMutex
	listRepositoriesByOrgArgsForCall []struct {
		logger  lager.Logger
		orgName string
	}
	listRepositoriesByOrgReturns struct {
		result1 []revok.GitHubRepository
		result2 error
	}
	listRepositoriesByOrgReturnsOnCall map[int]struct {
		result1 []revok.GitHubRepository
		result2 error
	}
	ListRepositoriesByUserStub        func(logger lager.Logger, userName string) ([]revok.GitHubRepository, error)
	listRepositoriesByUserMutex       sync.RWMutex
	listRepositoriesByUserArgsForCall []struct {
		logger   lager.Logger
		userName string
	}
	listRepositoriesByUserReturns struct {
		result1 []revok.GitHubRepository
		result2 error
	}
	listRepositoriesByUserReturnsOnCall map[int]struct {
		result1 []revok.GitHubRepository
		result2 error
	}
	GetRepoStub        func(logger lager.Logger, owner, repoName string) (*revok.GitHubRepository, error)
	getRepoMutex       sync.RWMutex
	getRepoArgsForCall []struct {
		logger   lager.Logger
		owner    string
		repoName string
	}
	getRepoReturns struct {
		result1 *revok.GitHubRepository
		result2 error
	}
	getRepoReturnsOnCall map[int]struct {
		result1 *revok.GitHubRepository
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGithubService) ListRepositoriesByOrg(logger lager.Logger, orgName string) ([]revok.GitHubRepository, error) {
	fake.listRepositoriesByOrgMutex.Lock()
	ret, specificReturn := fake.listRepositoriesByOrgReturnsOnCall[len(fake.listRepositoriesByOrgArgsForCall)]
	fake.listRepositoriesByOrgArgsForCall = append(fake.listRepositoriesByOrgArgsForCall, struct {
		logger  lager.Logger
		orgName string
	}{logger, orgName})
	fake.recordInvocation("ListRepositoriesByOrg", []interface{}{logger, orgName})
	fake.listRepositoriesByOrgMutex.Unlock()
	if fake.ListRepositoriesByOrgStub != nil {
		return fake.ListRepositoriesByOrgStub(logger, orgName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listRepositoriesByOrgReturns.result1, fake.listRepositoriesByOrgReturns.result2
}

func (fake *FakeGithubService) ListRepositoriesByOrgCallCount() int {
	fake.listRepositoriesByOrgMutex.RLock()
	defer fake.listRepositoriesByOrgMutex.RUnlock()
	return len(fake.listRepositoriesByOrgArgsForCall)
}

func (fake *FakeGithubService) ListRepositoriesByOrgArgsForCall(i int) (lager.Logger, string) {
	fake.listRepositoriesByOrgMutex.RLock()
	defer fake.listRepositoriesByOrgMutex.RUnlock()
	return fake.listRepositoriesByOrgArgsForCall[i].logger, fake.listRepositoriesByOrgArgsForCall[i].orgName
}

func (fake *FakeGithubService) ListRepositoriesByOrgReturns(result1 []revok.GitHubRepository, result2 error) {
	fake.ListRepositoriesByOrgStub = nil
	fake.listRepositoriesByOrgReturns = struct {
		result1 []revok.GitHubRepository
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubService) ListRepositoriesByOrgReturnsOnCall(i int, result1 []revok.GitHubRepository, result2 error) {
	fake.ListRepositoriesByOrgStub = nil
	if fake.listRepositoriesByOrgReturnsOnCall == nil {
		fake.listRepositoriesByOrgReturnsOnCall = make(map[int]struct {
			result1 []revok.GitHubRepository
			result2 error
		})
	}
	fake.listRepositoriesByOrgReturnsOnCall[i] = struct {
		result1 []revok.GitHubRepository
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubService) ListRepositoriesByUser(logger lager.Logger, userName string) ([]revok.GitHubRepository, error) {
	fake.listRepositoriesByUserMutex.Lock()
	ret, specificReturn := fake.listRepositoriesByUserReturnsOnCall[len(fake.listRepositoriesByUserArgsForCall)]
	fake.listRepositoriesByUserArgsForCall = append(fake.listRepositoriesByUserArgsForCall, struct {
		logger   lager.Logger
		userName string
	}{logger, userName})
	fake.recordInvocation("ListRepositoriesByUser", []interface{}{logger, userName})
	fake.listRepositoriesByUserMutex.Unlock()
	if fake.ListRepositoriesByUserStub != nil {
		return fake.ListRepositoriesByUserStub(logger, userName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listRepositoriesByUserReturns.result1, fake.listRepositoriesByUserReturns.result2
}

func (fake *FakeGithubService) ListRepositoriesByUserCallCount() int {
	fake.listRepositoriesByUserMutex.RLock()
	defer fake.listRepositoriesByUserMutex.RUnlock()
	return len(fake.listRepositoriesByUserArgsForCall)
}

func (fake *FakeGithubService) ListRepositoriesByUserArgsForCall(i int) (lager.Logger, string) {
	fake.listRepositoriesByUserMutex.RLock()
	defer fake.listRepositoriesByUserMutex.RUnlock()
	return fake.listRepositoriesByUserArgsForCall[i].logger, fake.listRepositoriesByUserArgsForCall[i].userName
}

func (fake *FakeGithubService) ListRepositoriesByUserReturns(result1 []revok.GitHubRepository, result2 error) {
	fake.ListRepositoriesByUserStub = nil
	fake.listRepositoriesByUserReturns = struct {
		result1 []revok.GitHubRepository
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubService) ListRepositoriesByUserReturnsOnCall(i int, result1 []revok.GitHubRepository, result2 error) {
	fake.ListRepositoriesByUserStub = nil
	if fake.listRepositoriesByUserReturnsOnCall == nil {
		fake.listRepositoriesByUserReturnsOnCall = make(map[int]struct {
			result1 []revok.GitHubRepository
			result2 error
		})
	}
	fake.listRepositoriesByUserReturnsOnCall[i] = struct {
		result1 []revok.GitHubRepository
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubService) GetRepo(logger lager.Logger, owner string, repoName string) (*revok.GitHubRepository, error) {
	fake.getRepoMutex.Lock()
	ret, specificReturn := fake.getRepoReturnsOnCall[len(fake.getRepoArgsForCall)]
	fake.getRepoArgsForCall = append(fake.getRepoArgsForCall, struct {
		logger   lager.Logger
		owner    string
		repoName string
	}{logger, owner, repoName})
	fake.recordInvocation("GetRepo", []interface{}{logger, owner, repoName})
	fake.getRepoMutex.Unlock()
	if fake.GetRepoStub != nil {
		return fake.GetRepoStub(logger, owner, repoName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getRepoReturns.result1, fake.getRepoReturns.result2
}

func (fake *FakeGithubService) GetRepoCallCount() int {
	fake.getRepoMutex.RLock()
	defer fake.getRepoMutex.RUnlock()
	return len(fake.getRepoArgsForCall)
}

func (fake *FakeGithubService) GetRepoArgsForCall(i int) (lager.Logger, string, string) {
	fake.getRepoMutex.RLock()
	defer fake.getRepoMutex.RUnlock()
	return fake.getRepoArgsForCall[i].logger, fake.getRepoArgsForCall[i].owner, fake.getRepoArgsForCall[i].repoName
}

func (fake *FakeGithubService) GetRepoReturns(result1 *revok.GitHubRepository, result2 error) {
	fake.GetRepoStub = nil
	fake.getRepoReturns = struct {
		result1 *revok.GitHubRepository
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubService) GetRepoReturnsOnCall(i int, result1 *revok.GitHubRepository, result2 error) {
	fake.GetRepoStub = nil
	if fake.getRepoReturnsOnCall == nil {
		fake.getRepoReturnsOnCall = make(map[int]struct {
			result1 *revok.GitHubRepository
			result2 error
		})
	}
	fake.getRepoReturnsOnCall[i] = struct {
		result1 *revok.GitHubRepository
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listRepositoriesByOrgMutex.RLock()
	defer fake.listRepositoriesByOrgMutex.RUnlock()
	fake.listRepositoriesByUserMutex.RLock()
	defer fake.listRepositoriesByUserMutex.RUnlock()
	fake.getRepoMutex.RLock()
	defer fake.getRepoMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGithubService) recordInvocation(key string, args []interface{}) {
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

var _ revok.GithubService = new(FakeGithubService)
