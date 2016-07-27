// This file was generated by counterfeiter
package githubclientfakes

import (
	"cred-alert/githubclient"
	"net/url"
	"sync"

	"github.com/pivotal-golang/lager"
)

type FakeClient struct {
	CompareRefsStub        func(logger lager.Logger, owner, repo, base, head string) (string, error)
	compareRefsMutex       sync.RWMutex
	compareRefsArgsForCall []struct {
		logger lager.Logger
		owner  string
		repo   string
		base   string
		head   string
	}
	compareRefsReturns struct {
		result1 string
		result2 error
	}
	ArchiveLinkStub        func(owner, repo, ref string) (*url.URL, error)
	archiveLinkMutex       sync.RWMutex
	archiveLinkArgsForCall []struct {
		owner string
		repo  string
		ref   string
	}
	archiveLinkReturns struct {
		result1 *url.URL
		result2 error
	}
	ParentsStub        func(logger lager.Logger, owner, repo, sha string) ([]string, error)
	parentsMutex       sync.RWMutex
	parentsArgsForCall []struct {
		logger lager.Logger
		owner  string
		repo   string
		sha    string
	}
	parentsReturns struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) CompareRefs(logger lager.Logger, owner string, repo string, base string, head string) (string, error) {
	fake.compareRefsMutex.Lock()
	fake.compareRefsArgsForCall = append(fake.compareRefsArgsForCall, struct {
		logger lager.Logger
		owner  string
		repo   string
		base   string
		head   string
	}{logger, owner, repo, base, head})
	fake.recordInvocation("CompareRefs", []interface{}{logger, owner, repo, base, head})
	fake.compareRefsMutex.Unlock()
	if fake.CompareRefsStub != nil {
		return fake.CompareRefsStub(logger, owner, repo, base, head)
	} else {
		return fake.compareRefsReturns.result1, fake.compareRefsReturns.result2
	}
}

func (fake *FakeClient) CompareRefsCallCount() int {
	fake.compareRefsMutex.RLock()
	defer fake.compareRefsMutex.RUnlock()
	return len(fake.compareRefsArgsForCall)
}

func (fake *FakeClient) CompareRefsArgsForCall(i int) (lager.Logger, string, string, string, string) {
	fake.compareRefsMutex.RLock()
	defer fake.compareRefsMutex.RUnlock()
	return fake.compareRefsArgsForCall[i].logger, fake.compareRefsArgsForCall[i].owner, fake.compareRefsArgsForCall[i].repo, fake.compareRefsArgsForCall[i].base, fake.compareRefsArgsForCall[i].head
}

func (fake *FakeClient) CompareRefsReturns(result1 string, result2 error) {
	fake.CompareRefsStub = nil
	fake.compareRefsReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ArchiveLink(owner string, repo string, ref string) (*url.URL, error) {
	fake.archiveLinkMutex.Lock()
	fake.archiveLinkArgsForCall = append(fake.archiveLinkArgsForCall, struct {
		owner string
		repo  string
		ref   string
	}{owner, repo, ref})
	fake.recordInvocation("ArchiveLink", []interface{}{owner, repo, ref})
	fake.archiveLinkMutex.Unlock()
	if fake.ArchiveLinkStub != nil {
		return fake.ArchiveLinkStub(owner, repo, ref)
	} else {
		return fake.archiveLinkReturns.result1, fake.archiveLinkReturns.result2
	}
}

func (fake *FakeClient) ArchiveLinkCallCount() int {
	fake.archiveLinkMutex.RLock()
	defer fake.archiveLinkMutex.RUnlock()
	return len(fake.archiveLinkArgsForCall)
}

func (fake *FakeClient) ArchiveLinkArgsForCall(i int) (string, string, string) {
	fake.archiveLinkMutex.RLock()
	defer fake.archiveLinkMutex.RUnlock()
	return fake.archiveLinkArgsForCall[i].owner, fake.archiveLinkArgsForCall[i].repo, fake.archiveLinkArgsForCall[i].ref
}

func (fake *FakeClient) ArchiveLinkReturns(result1 *url.URL, result2 error) {
	fake.ArchiveLinkStub = nil
	fake.archiveLinkReturns = struct {
		result1 *url.URL
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Parents(logger lager.Logger, owner string, repo string, sha string) ([]string, error) {
	fake.parentsMutex.Lock()
	fake.parentsArgsForCall = append(fake.parentsArgsForCall, struct {
		logger lager.Logger
		owner  string
		repo   string
		sha    string
	}{logger, owner, repo, sha})
	fake.recordInvocation("Parents", []interface{}{logger, owner, repo, sha})
	fake.parentsMutex.Unlock()
	if fake.ParentsStub != nil {
		return fake.ParentsStub(logger, owner, repo, sha)
	} else {
		return fake.parentsReturns.result1, fake.parentsReturns.result2
	}
}

func (fake *FakeClient) ParentsCallCount() int {
	fake.parentsMutex.RLock()
	defer fake.parentsMutex.RUnlock()
	return len(fake.parentsArgsForCall)
}

func (fake *FakeClient) ParentsArgsForCall(i int) (lager.Logger, string, string, string) {
	fake.parentsMutex.RLock()
	defer fake.parentsMutex.RUnlock()
	return fake.parentsArgsForCall[i].logger, fake.parentsArgsForCall[i].owner, fake.parentsArgsForCall[i].repo, fake.parentsArgsForCall[i].sha
}

func (fake *FakeClient) ParentsReturns(result1 []string, result2 error) {
	fake.ParentsStub = nil
	fake.parentsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.compareRefsMutex.RLock()
	defer fake.compareRefsMutex.RUnlock()
	fake.archiveLinkMutex.RLock()
	defer fake.archiveLinkMutex.RUnlock()
	fake.parentsMutex.RLock()
	defer fake.parentsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ githubclient.Client = new(FakeClient)