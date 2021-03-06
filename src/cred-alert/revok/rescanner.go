package revok

import (
	"os"

	"code.cloudfoundry.org/lager"

	"context"
	"cred-alert/db"
	"cred-alert/metrics"
	"cred-alert/notifications"
	"cred-alert/sniff"
)

//go:generate counterfeiter . RescannerScanner

type RescannerScanner interface {
	Scan(lager.Logger, string, string, map[string]struct{}, string, string, string) ([]db.Credential, error)
}

type Rescanner struct {
	logger         lager.Logger
	scanRepo       db.ScanRepository
	credRepo       db.CredentialRepository
	scanner        RescannerScanner
	router         notifications.Router
	successCounter metrics.Counter
	failedCounter  metrics.Counter
}

func NewRescanner(
	logger lager.Logger,
	scanRepo db.ScanRepository,
	credRepo db.CredentialRepository,
	scanner RescannerScanner,
	router notifications.Router,
	emitter metrics.Emitter,
) *Rescanner {
	return &Rescanner{
		logger:         logger,
		scanRepo:       scanRepo,
		credRepo:       credRepo,
		scanner:        scanner,
		router:         router,
		successCounter: emitter.Counter("revok.rescanner.success"),
		failedCounter:  emitter.Counter("revok.rescanner.failed"),
	}
}

func (r *Rescanner) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	logger := r.logger.Session("rescanner")
	logger.Info("started")

	defer logger.Info("done")

	close(ready)

	priorScans, err := r.scanRepo.ScansNotYetRunWithVersion(logger, sniff.RulesVersion)
	if err != nil {
		logger.Error("failed-getting-prior-scans", err)
	}

	for _, priorScan := range priorScans {
		select {
		case <-signals:
			return nil
		default:
			err := r.work(logger, priorScan)
			if err != nil {
				r.failedCounter.Inc(logger)
				logger.Error("failed-to-rescan", err, lager.Data{
					"scan-id": priorScan.ID,
				})
			}
		}
	}

	logger.Info("all-scans-up-to-date")
	<-signals
	return nil
}

func (r *Rescanner) work(logger lager.Logger, priorScan db.PriorScan) error {
	logger.Info("rescanning", lager.Data{
		"owner":   priorScan.Owner,
		"repo":    priorScan.Repository,
		"scan-id": priorScan.ID,
	})

	oldCredentials, err := r.credRepo.ForScanWithID(priorScan.ID)
	if err != nil {
		logger.Error("failed-getting-prior-credentials", err)
		return err
	}

	credMap := make(map[string]db.Credential, len(oldCredentials))
	for _, cred := range oldCredentials {
		credMap[cred.Hash()] = cred
	}

	newCredentials, err := r.scanner.Scan(
		logger,
		priorScan.Owner,
		priorScan.Repository,
		map[string]struct{}{},
		priorScan.Branch,
		priorScan.StartSHA,
		priorScan.StopSHA,
	)
	if err != nil {
		return err
	}

	r.successCounter.Inc(logger)

	var first error
	var batch []notifications.Notification
	for _, cred := range newCredentials {
		if _, ok := credMap[cred.Hash()]; ok {
			continue
		}
		credReported, err := r.credRepo.CredentialReported(&cred)

		if err != nil {
			if first == nil {
				first = err
			}
			logger.Error("failed-checking-for-existing-credential", err)
		}

		if !credReported {
			batch = append(batch, notifications.Notification{
				Owner:      cred.Owner,
				Repository: cred.Repository,
				SHA:        cred.SHA,
				Path:       cred.Path,
				LineNumber: cred.LineNumber,
				Private:    cred.Private,
			})
		}
	}

	if len(batch) > 0 {
		err = r.router.Deliver(context.TODO(), logger, batch)
		if err != nil {
			logger.Error("failed-to-notify", err)
		}
	}

	return first
}
