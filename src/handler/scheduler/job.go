package scheduler

func (s *scheduler) DummyJob(param string) {
	s.log.Debug().Any("JOB_DUMMY", param).Send()
}

func (s *scheduler) DummyJob2(param string) {
	s.log.Debug().Any("JOB_DUMMY_2", param).Send()
}
