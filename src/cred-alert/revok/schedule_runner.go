package revok

import (
	"os"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/robfig/cron"
)

type ScheduleRunner struct {
	logger lager.Logger

	cron    *cron.Cron
	cronMut *sync.Mutex

	jobWg *sync.WaitGroup
}

func Schedule(logger lager.Logger, cron string, work func()) *ScheduleRunner {
	runner := NewScheduleRunner(logger)
	runner.ScheduleWork(cron, work)

	return runner
}

func NewScheduleRunner(logger lager.Logger) *ScheduleRunner {
	return &ScheduleRunner{
		logger:  logger,
		cron:    cron.New(),
		cronMut: &sync.Mutex{},
		jobWg:   &sync.WaitGroup{},
	}
}

func (s *ScheduleRunner) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	logger := s.logger.Session("schedule-runner")

	s.cron.Start()

	close(ready)

	logger.Info("started")
	defer logger.Info("done")

	select {
	case <-signals:
		logger.Info("signalled")

		s.cron.Stop()
		s.jobWg.Wait()
	}

	return nil
}

func (s *ScheduleRunner) ScheduleWork(cron string, work func()) {
	s.cronMut.Lock()
	defer s.cronMut.Unlock()

	s.cron.AddFunc(cron, func() {
		s.jobWg.Add(1)
		defer s.jobWg.Done()

		work()
	})
}
