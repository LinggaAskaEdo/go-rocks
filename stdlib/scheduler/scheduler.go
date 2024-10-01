package scheduler

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/rs/zerolog"
)

type Options struct {
}

func Init(log zerolog.Logger) gocron.Scheduler {
	sched, err := gocron.NewScheduler()
	if err != nil {
		log.Panic().Err(err).Msg("Job init failed !!!")
	}

	return sched
}
