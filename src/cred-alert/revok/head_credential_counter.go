package revok

import (
	"context"
	"os"
	"time"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"

	"cred-alert/db"
	credlog "cred-alert/log"
	"cred-alert/sniff"
)

//go:generate counterfeiter . GitBranchCredentialsCounterClient

type GitBranchCredentialsCounterClient interface {
	BranchTargets(repoPath string) (map[string]string, error)
	BranchCredentialCounts(lager.Logger, string, sniff.Sniffer) (map[string]uint, error)
}

type HeadCredentialCounter struct {
	logger               lager.Logger
	branchRepository     db.BranchRepository
	repositoryRepository db.RepositoryRepository
	clock                clock.Clock
	interval             time.Duration
	gitClient            GitBranchCredentialsCounterClient
	sniffer              sniff.Sniffer
}

func NewHeadCredentialCounter(
	logger lager.Logger,
	branchRepository db.BranchRepository,
	repositoryRepository db.RepositoryRepository,
	clock clock.Clock,
	interval time.Duration,
	gitClient GitBranchCredentialsCounterClient,
	sniffer sniff.Sniffer,
) *HeadCredentialCounter {
	return &HeadCredentialCounter{
		logger:               logger,
		branchRepository:     branchRepository,
		repositoryRepository: repositoryRepository,
		clock:                clock,
		interval:             interval,
		gitClient:            gitClient,
		sniffer:              sniffer,
	}
}

func (c *HeadCredentialCounter) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	logger := c.logger.Session("head-credential-counter")
	logger.Info("starting")

	close(ready)

	timer := c.clock.NewTicker(c.interval)

	ctx, cancel := context.WithCancel(context.Background())

	quietLogger := credlog.NewNullLogger()

	c.work(cancel, signals, logger, quietLogger)

	for {
		select {
		case <-timer.C():
			c.work(cancel, signals, logger, quietLogger)
		case <-signals:
			cancel()
			return nil
		case <-ctx.Done():
			return nil
		}
	}
}

func (c *HeadCredentialCounter) work(
	cancel context.CancelFunc,
	signals <-chan os.Signal,
	logger lager.Logger,
	quietLogger lager.Logger,
) {
	repositories, err := c.repositoryRepository.All()
	if err != nil {
		logger.Error("failed-getting-all-repositories", err)
	}

	for _, repository := range repositories {
		select {
		case <-signals:
			cancel()
			return
		default:
			repoLogger := logger.WithData(lager.Data{
				"ref":  repository.DefaultBranch,
				"path": repository.Path,
			})

			if !repository.Cloned {
				repoLogger.Debug("skipping-uncloned-repository")
				continue
			}

			credentialCounts, err := c.gitClient.BranchCredentialCounts(quietLogger, repository.Path, c.sniffer)
			if err != nil {
				repoLogger.Error("failed-to-get-credential-counts", err)
				continue
			}

			branches := make([]db.Branch, 0, len(credentialCounts))
			for branchName, credentialCount := range credentialCounts {
				branches = append(branches, db.Branch{
					Name:            branchName,
					CredentialCount: credentialCount,
				})
			}

			err = c.branchRepository.UpdateBranches(repository, branches)
			if err != nil {
				repoLogger.Error("failed-to-update-credential-count", err)
				continue
			}

			repoLogger.Debug("updated-credential-count")
		}
	}
}
