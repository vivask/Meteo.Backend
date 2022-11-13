package v1

import (
	"fmt"

	"meteo/internal/entities"
	"meteo/internal/log"

	"github.com/go-co-op/gocron"
)

func (p scheduleAPI) addJob(job *entities.Jobs) (*gocron.Job, error) {

	if !job.Active {
		log.Warningf("Can't run inactive job")
		return nil, nil
	}

	repeat := job.Value

	options := Options{
		job:    *job,
		off:    false,
		repeat: false,
	}

	dt := NewTime(job.Date, job.Time)

	switch job.Period.ID {
	case "once":
		options.off = true
		return p.createJob(dt, job, &options, p.cron.Days())
	case "second":
		return p.createSimpleJob(dt, job, &options, p.cron.Every(repeat).Seconds())
	case "minute":
		return p.createSimpleJob(dt, job, &options, p.cron.Every(repeat).Minutes())
	case "hour":
		return p.createSimpleJob(dt, job, &options, p.cron.Every(repeat).Hours())
	case "day":
		return p.createJob(dt, job, &options, p.cron.Every(repeat).Days())
	case "week":
		return p.createJob(dt, job, &options, p.cron.Every(repeat).Weeks())
	case "month":
		return p.createJob(dt, job, &options, p.cron.Every(repeat).Month())
	case "year":
		return p.createYearJob(dt, repeat, job, &options, p.cron.Every(1))
	case "day_of_week":
		switch job.Day.ID {
		case 1:
			return p.createDayOfWeekJob(dt, repeat, job, &options, p.cron.Every(1).Monday())
		case 2:
			return p.createDayOfWeekJob(dt, repeat, job, &options, p.cron.Every(1).Tuesday())
		case 3:
			return p.createDayOfWeekJob(dt, repeat, job, &options, p.cron.Every(1).Wednesday())
		case 4:
			return p.createDayOfWeekJob(dt, repeat, job, &options, p.cron.Every(1).Thursday())
		case 5:
			return p.createDayOfWeekJob(dt, repeat, job, &options, p.cron.Every(1).Friday())
		case 6:
			return p.createDayOfWeekJob(dt, repeat, job, &options, p.cron.Every(1).Saturday())
		case 7:
			return p.createDayOfWeekJob(dt, repeat, job, &options, p.cron.Every(1).Sunday())
		default:
			return nil, fmt.Errorf("invalid day of the week: %v", job.Day.ID)
		}
	default:
		return nil, fmt.Errorf("unknown period id: %s", job.Period.ID)
	}

}

func (p scheduleAPI) createSimpleJob(dt *DateTime, job *entities.Jobs, options *Options, fn *gocron.Scheduler) (*gocron.Job, error) {
	if dt.IsZero() {
		return fn.Tag(job.Note).Do(p.jobFunc, options)
	} else {
		stamp, err := dt.Stamp()
		if err != nil {
			return nil, fmt.Errorf("Stamp error: %w", err)
		}
		return fn.StartAt(stamp).Tag(job.Note).Do(p.jobFunc, options)
	}
}

func (p scheduleAPI) createJob(dt *DateTime, job *entities.Jobs, options *Options, fn *gocron.Scheduler) (*gocron.Job, error) {
	if dt.IsZero() {
		return fn.Tag(job.Note).Do(p.jobFunc, &options)
	} else {
		if dt.TimeOnly() {
			return fn.At(dt.Time()).Tag(job.Note).Do(p.jobFunc, options)
		} else {
			stamp, err := dt.Stamp()
			if err != nil {
				return nil, fmt.Errorf("Stamp error: %w", err)
			}
			return fn.StartAt(stamp).Tag(job.Note).Do(p.jobFunc, options)
		}
	}
}

func (p scheduleAPI) createDayOfWeekJob(dt *DateTime, repeat int, job *entities.Jobs, options *Options, fn *gocron.Scheduler) (*gocron.Job, error) {

	if repeat < 0 || repeat > 4 {
		return nil, fmt.Errorf("number [%d] week of month incorrect, expected 0-4", repeat)
	}

	if dt.IsZero() {
		dt.ts = "00:00:00"
	}

	options.repeat = false

	if dt.TimeOnly() {
		if repeat < 1 {
			return fn.At(dt.Time()).Tag(job.Note).Do(p.jobFunc, options)
			//weekDay, _ := cronJob.Weekday()
			//p.logger.Debugf("Job [%s] Start On: %v At: %v", job.Note, weekDay, dt.Time())
		} else {
			stamp, err := getNextDateStartJob(int(job.Day.ID), repeat, dt.Time(), job.Note)
			if err != nil {
				return nil, fmt.Errorf("getNextDateStartJob error: %w", err)
			}
			options.repeat = true
			return fn.StartAt(stamp).Tag(job.Note).Do(p.jobFunc, options)
		}
	} else {
		if repeat < 1 {
			stamp, err := dt.Stamp()
			if err != nil {
				return nil, fmt.Errorf("stamp error: %w", err)
			}
			//p.logger.Debugf("Job [%s] Start At: %v", job.Note, stamp)
			return fn.StartAt(stamp).Tag(job.Note).Do(p.jobFunc, options)
		} else {
			stamp, err := getNextDateStartJob(int(job.Day.ID), repeat, dt.Time(), job.Note)
			if err != nil {
				return nil, fmt.Errorf("getNextDateStartJob error: %w", err)
			}
			options.repeat = true
			//p.logger.Debugf("Job [%s] Start At: %v", job.Note, stamp)
			return fn.StartAt(stamp).Tag(job.Note).Do(p.jobFunc, options)
		}
	}
}

func (p scheduleAPI) createYearJob(dt *DateTime, repeat int, job *entities.Jobs, options *Options, fn *gocron.Scheduler) (cronJob *gocron.Job, err error) {
	if dt.IsZero() || dt.IsDateZero() {
		return nil, fmt.Errorf("undefined date start job")
	}
	stamp, err := getNextYearStartJob(repeat, dt)
	if err != nil {
		return nil, fmt.Errorf("getNextYearStartJob error: %w", err)
	}
	options.repeat = true
	cronJob, err = fn.StartAt(stamp).Tag(job.Note).Do(p.jobFunc, options)

	return
}
