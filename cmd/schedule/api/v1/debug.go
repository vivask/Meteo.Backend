package v1

import (
	"meteo/internal/entities"
	"meteo/internal/log"
)

func (p scheduleAPI) LogActiveJobs() {
	for _, job := range p.cron.Jobs() {
		switch {
		case !job.ScheduledTime().IsZero():
			for _, tag := range job.Tags() {
				log.Debugf("JOB [%s] ScheduledTime: %v", tag, job.ScheduledTime())
			}
		case len(job.ScheduledAtTime()) != 0:
			for _, tag := range job.Tags() {
				log.Debugf("JOB [%s] ScheduledAtTime: %v", tag, job.ScheduledAtTime())
			}
		}
	}
}

func (p scheduleAPI) getCronJobs() (jobs []entities.CronJobs) {
	for _, job := range p.cron.Jobs() {
		for _, tag := range job.Tags() {
			jobs = append(jobs, entities.CronJobs{Note: tag, ScheduledTime: job.ScheduledTime()})
		}
	}
	return
}
