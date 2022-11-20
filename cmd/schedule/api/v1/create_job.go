package v1

import (
	"fmt"

	"meteo/internal/entities"
	"meteo/internal/log"

	"github.com/go-co-op/gocron"
)

func (p scheduleAPI) addJob(job entities.Jobs) (*gocron.Job, error) {

	if !job.Active {
		log.Warningf("Can't run inactive job")
		return nil, nil
	}

	repeat := job.Value

	options := Options{
		job:    job,
		off:    false,
		repeat: false,
	}

	dt := NewTime(job.Date, job.Time)

	switch job.Period.ID {
	case "once":
		options.off = true
		return p.createJob(dt, options, p.cron.Every(1).Days())
	case "second":
		return p.createSimpleJob(dt, options, p.cron.Every(repeat).Seconds())
	case "minute":
		return p.createSimpleJob(dt, options, p.cron.Every(repeat).Minutes())
	case "hour":
		return p.createSimpleJob(dt, options, p.cron.Every(repeat).Hours())
	case "day":
		return p.createJob(dt, options, p.cron.Every(repeat).Days())
	case "week":
		return p.createWeeksJob(dt, repeat, options, p.cron.Every(repeat).Weeks())
	case "month":
		return p.createMonthJob(dt, options, p.cron.Every(repeat))
	case "year":
		return p.createYearJob(dt, repeat, options, p.cron.Every(1))
	default:
		return nil, fmt.Errorf("unknown period id: %s", job.Period.ID)
	}

}

func (p scheduleAPI) createSimpleJob(dt *DateTime, options Options, fn *gocron.Scheduler) (*gocron.Job, error) {
	if dt.IsZero() {
		return fn.Tag(options.job.Note).Do(p.jobFunc, options)
	} else {
		stamp, err := dt.Stamp()
		if err != nil {
			return nil, fmt.Errorf("Stamp error: %w", err)
		}
		return fn.StartAt(stamp).Tag(options.job.Note).Do(p.jobFunc, options)
	}
}

func (p scheduleAPI) createJob(dt *DateTime, options Options, fn *gocron.Scheduler) (*gocron.Job, error) {
	if dt.IsZero() {
		return fn.Tag(options.job.Note).Do(p.jobFunc, options)
	} else {
		if dt.TimeOnly() {
			return fn.At(dt.Time()).Tag(options.job.Note).Do(p.jobFunc, options)
		} else {
			stamp, err := dt.Stamp()
			if err != nil {
				return nil, fmt.Errorf("Stamp error: %w", err)
			}
			return fn.StartAt(stamp).Tag(options.job.Note).Do(p.jobFunc, options)
		}
	}
}

func (p scheduleAPI) createWeeksJob(dt *DateTime, repeat int, options Options, fn *gocron.Scheduler) (*gocron.Job, error) {
	if dt.IsZero() {
		return fn.Tag(options.job.Note).Do(p.jobFunc, options)
	} else {
		if dt.TimeOnly() {
			return fn.At(dt.Time()).Tag(options.job.Note).Do(p.jobFunc, options)
		} else {
			if repeat > 0 {
				stamp, err := getWeekDateStartJob(repeat, dt)
				if err != nil {
					return nil, fmt.Errorf("Stamp error: %w", err)
				}
				return fn.StartAt(stamp).Tag(options.job.Note).Do(p.jobFunc, options)
			} else {
				stamp, err := dt.Stamp()
				if err != nil {
					return nil, fmt.Errorf("Stamp error: %w", err)
				}
				return fn.StartAt(stamp).Tag(options.job.Note).Do(p.jobFunc, options)
			}
		}
	}
}

func (p scheduleAPI) createMonthJob(dt *DateTime, options Options, fn *gocron.Scheduler) (*gocron.Job, error) {
	day, err := dt.Day()
	if err != nil {
		return nil, fmt.Errorf("get day error: %w", err)
	}
	if day < 29 {
		return fn.Month(day).Tag(options.job.Note).Do(p.jobFunc, options)
	} else {
		return fn.MonthLastDay().Tag(options.job.Note).Do(p.jobFunc, options)
	}
}

func (p scheduleAPI) createYearJob(dt *DateTime, repeat int, options Options, fn *gocron.Scheduler) (*gocron.Job, error) {
	if dt.IsZero() {
		return fn.Tag(options.job.Note).Do(p.jobFunc, options)
	} else {
		if dt.TimeOnly() {
			return fn.At(dt.Time()).Tag(options.job.Note).Do(p.jobFunc, options)
		} else {
			if repeat > 0 {
				stamp, err := getYearDateStartJob(repeat, dt)
				if err != nil {
					return nil, fmt.Errorf("Stamp error: %w", err)
				}
				return fn.StartAt(stamp).Tag(options.job.Note).Do(p.jobFunc, options)
			} else {
				stamp, err := dt.Stamp()
				if err != nil {
					return nil, fmt.Errorf("Stamp error: %w", err)
				}
				return fn.StartAt(stamp).Tag(options.job.Note).Do(p.jobFunc, options)
			}
		}
	}
}
