// This file was generated by counterfeiter
package teamstrfakes

import (
	"context"
	"sync"
	"teamstr"

	"github.com/google/go-github/github"
)

type FakeRepoRepo struct {
	ListByOrgStub        func(ctx context.Context, org string, opt *github.RepositoryListByOrgOptions) ([]*github.Repository, *github.Response, error)
	listByOrgMutex       sync.RWMutex
	listByOrgArgsForCall []struct {
		ctx context.Context
		org string
		opt *github.RepositoryListByOrgOptions
	}
	listByOrgReturns struct {
		result1 []*github.Repository
		result2 *github.Response
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRepoRepo) ListByOrg(ctx context.Context, org string, opt *github.RepositoryListByOrgOptions) ([]*github.Repository, *github.Response, error) {
	fake.listByOrgMutex.Lock()
	fake.listByOrgArgsForCall = append(fake.listByOrgArgsForCall, struct {
		ctx context.Context
		org string
		opt *github.RepositoryListByOrgOptions
	}{ctx, org, opt})
	fake.recordInvocation("ListByOrg", []interface{}{ctx, org, opt})
	fake.listByOrgMutex.Unlock()
	if fake.ListByOrgStub != nil {
		return fake.ListByOrgStub(ctx, org, opt)
	}
	return fake.listByOrgReturns.result1, fake.listByOrgReturns.result2, fake.listByOrgReturns.result3
}

func (fake *FakeRepoRepo) ListByOrgCallCount() int {
	fake.listByOrgMutex.RLock()
	defer fake.listByOrgMutex.RUnlock()
	return len(fake.listByOrgArgsForCall)
}

func (fake *FakeRepoRepo) ListByOrgArgsForCall(i int) (context.Context, string, *github.RepositoryListByOrgOptions) {
	fake.listByOrgMutex.RLock()
	defer fake.listByOrgMutex.RUnlock()
	return fake.listByOrgArgsForCall[i].ctx, fake.listByOrgArgsForCall[i].org, fake.listByOrgArgsForCall[i].opt
}

func (fake *FakeRepoRepo) ListByOrgReturns(result1 []*github.Repository, result2 *github.Response, result3 error) {
	fake.ListByOrgStub = nil
	fake.listByOrgReturns = struct {
		result1 []*github.Repository
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepoRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listByOrgMutex.RLock()
	defer fake.listByOrgMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRepoRepo) recordInvocation(key string, args []interface{}) {
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

var _ teamstr.RepoRepo = new(FakeRepoRepo)
