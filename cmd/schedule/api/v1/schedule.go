package v1

import (
	"fmt"
	"meteo/internal/config"
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/log"

	"github.com/go-co-op/gocron"
)

var started bool

type Options struct {
	job    entities.Jobs
	off    bool
	repeat bool
}

func (p scheduleAPI) jobFunc(o *Options) {
	log.Debugf("Job [%s], api [%s] running...", o.job.Note, o.job.Task.Api)
	log.Debugf("Params: %v", o.job.Params)
	_, err := kit.PutInt(o.job.Task.Api, o.job.Params)
	if err != nil {
		log.Errorf("Job: [%v] executed error: %v", o.job.Note, err)
	} else {
		if o.job.Verbose || config.Default.Schedule.LogLevel == "debug" {
			log.Infof("Job [%v] executed success", o.job.Note)
		}
	}
	if o.off {
		if val, ok := p.jobs[o.job.ID]; ok {
			p.cron.RemoveByReference(val)
			delete(p.jobs, o.job.ID)
		} else {
			log.Errorf("Cron ID: %d not found", o.job.ID)
			return
		}
		err = p.repo.DeactivateJob(o.job.ID, false)
		if err != nil {
			log.Error(err)
		}
	}
	if o.repeat {
		err = p.reloadJobs()
		if err != nil {
			log.Errorf("reload jobs error: %v", err)
		}
	}
}

func (p scheduleAPI) reloadJobs() error {
	p.cron.Clear()
	p.jobs = make(map[uint32]*gocron.Job)
	err := p.executeJobs()
	if err != nil {
		return fmt.Errorf("exeJobs error: %w", err)
	}
	p.cron.Update()
	p.LogActiveJobs()
	log.Debug("Jobs reloaded success")
	return nil
}

func (p scheduleAPI) executeJobs() error {

	jobs, err := p.repo.GetAllActiveJobs()
	if err != nil {
		return fmt.Errorf("read jobs error: %w", err)
	}
	if len(jobs) == 0 {
		return fmt.Errorf("job list is empty")
	}

	log.Warningf("Cron jobs available: %v", len(jobs))
	for _, job := range jobs {
		err = p.selectTask(&job)
		if err != nil {
			p.repo.DeactivateJob(job.ID, false)
			log.Errorf("can't start job [%s] error: %v", job.Note, err)
		}
	}
	return nil
}

func (p scheduleAPI) executeJob(job_id uint32) error {
	job, err := p.repo.GetJobByID(job_id)
	if err != nil {
		return fmt.Errorf("GetJob error: %w", err)
	}
	err = p.selectTask(job)
	if err != nil {
		return fmt.Errorf("selectTask error: %w", err)
	}
	return nil
}

func (p scheduleAPI) selectTask(job *entities.Jobs) error {
	if job.Executor.ID == "Master" && !kit.IsMain() {
		log.Warningf("can't run job [%s] as Slave, need Master", job.Note)
		return nil
	}
	if job.Executor.ID == "Slave" && kit.IsMain() {
		log.Warningf("can't run job [%s] as Master, need Slave", job.Note)
		return nil
	}
	if job.Executor.ID == "Leader" && !kit.IsLeader() {
		log.Warningf("can't run job [%s], need Leader", job.Note)
		return nil
	}

	cronJob, err := p.addJob(job)
	if err != nil || cronJob == nil {
		log.Errorf("addJob error: %v", err)
		return fmt.Errorf("addJob error: %w", err)
	}
	p.jobs[job.ID] = cronJob
	return nil
}

func (p scheduleAPI) StartCron() {

	err := p.executeJobs()
	if err != nil {
		log.Warningf("Can't execute jobs: %v", err)
	}

	go p.cron.StartAsync()
	log.Infof("%s: success started", config.Default.Schedule.Title)

	p.LogActiveJobs()

	started = true
}

func (p scheduleAPI) StopCron() {
	p.cron.Stop()
	log.Debugf("%s: success stoped", config.Default.Schedule.Title)
}
