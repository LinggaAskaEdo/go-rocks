package scheduler

import (
	"sync"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/rs/zerolog"
)

var once = &sync.Once{}

type Scheduler interface{}

type scheduler struct {
	log       zerolog.Logger
	opt       Options
	scheduler gocron.Scheduler
}

type Options struct {
	JobDummy1 JobDummy1Options
	JobDummy2 JobDummy2Options
}

type JobDummy1Options struct {
	Enable bool
	Time   int
}

type JobDummy2Options struct {
	Enable bool
	Time   int
}

func Init(log zerolog.Logger, opt Options, s gocron.Scheduler) {
	once.Do(func() {
		sched := &scheduler{
			log:       log,
			opt:       opt,
			scheduler: s,
		}

		sched.Serve()
	})
}

func (s *scheduler) Serve() {
	var (
		jobDummy1ID string
		jobDummy2ID string
	)

	if s.opt.JobDummy1.Enable {
		jobDummy, err := s.scheduler.NewJob(
			gocron.DurationJob(
				time.Duration(s.opt.JobDummy1.Time*int(time.Second)),
			),
			gocron.NewTask(func() {
				s.DummyJob("Hello From Job Dummy")
			}),
		)
		if err != nil {
			s.log.Err(err).Send()
		}

		jobDummy1ID = jobDummy.ID().String()
		s.log.Debug().Msg("JOB - Name: Dummy, ID: " + jobDummy1ID)
	} else {
		s.log.Debug().Msg("JOB - Name: Dummy, ID: -")
	}

	if s.opt.JobDummy2.Enable {
		jobDummy2, err := s.scheduler.NewJob(
			gocron.DurationJob(
				time.Duration(s.opt.JobDummy2.Time*int(time.Second)),
			),
			gocron.NewTask(func() {
				s.DummyJob2("Hello From Job Dummy 2")
			}),
		)
		if err != nil {
			s.log.Err(err).Send()
		}

		jobDummy2ID = jobDummy2.ID().String()
		s.log.Debug().Msg("JOB - Name: Dummy2, ID: " + jobDummy2ID)
	} else {
		s.log.Debug().Msg("JOB - Name: Dummy2, ID: -")
	}

	s.scheduler.Start()
}
