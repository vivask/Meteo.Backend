package v1

import (
	"fmt"
	"time"
)

type DateTime struct {
	ds string
	ts string
}

func NewTime(d, t string) *DateTime {
	return &DateTime{
		ds: d,
		ts: t,
	}
}

func (d DateTime) IsZero() bool {
	if len(d.ds) == 0 && len(d.ts) == 0 {
		return true
	}
	return false
}

func (d DateTime) IsTimeZero() bool {
	return len(d.ts) == 0
}

func (d DateTime) IsDateZero() bool {
	return len(d.ds) == 0
}

func (d DateTime) DateOnly() bool {
	if len(d.ds) != 0 && len(d.ts) == 0 {
		return true
	}
	return false
}

func (d DateTime) TimeOnly() bool {
	if len(d.ds) == 0 && len(d.ts) != 0 {
		return true
	}
	return false
}

func (d DateTime) IsFull() bool {
	if len(d.ds) != 0 && len(d.ts) != 0 {
		return true
	}
	return false
}

func (d DateTime) Date() string {
	return d.ds
}

func (d DateTime) Time() string {
	return d.ts
}

func (d DateTime) Stamp() (dt time.Time, err error) {
	ds := d.ds
	if len(d.ds) == 0 {
		ds = time.Now().Format("2006-01-02")
	}
	ts := "00:00:00"
	if len(d.ts) != 0 {
		ts = d.ts
	}
	dts := fmt.Sprintf("%s %s", ds, ts)
	dt, err = time.ParseInLocation("2006-01-02 15:04:05", dts, time.Local)
	if err != nil {
		return dt, fmt.Errorf("parse date time error: %w", err)
	}
	if dt.Unix() < time.Now().Unix() {
		return dt, fmt.Errorf("unable to start the job at the specified time")
	}
	return dt, nil
}

func (d DateTime) Parse() (dt time.Time, err error) {
	ds := d.ds
	if len(d.ds) == 0 {
		ds = time.Now().Format("2006-01-02")
	}
	ts := "00:00:00"
	if len(d.ts) != 0 {
		ts = d.ts
	}
	dts := fmt.Sprintf("%s %s", ds, ts)
	dt, err = time.ParseInLocation("2006-01-02 15:04:05", dts, time.Local)
	if err != nil {
		return dt, fmt.Errorf("parse date time error: %w", err)
	}
	return dt, nil
}

func getNumWeekOfMonth(dt time.Time) int {
	beginningOfTheMonth := time.Date(dt.Year(), dt.Month(), 1, 1, 1, 1, 1, time.UTC)
	_, thisWeek := dt.ISOWeek()
	_, beginningWeek := beginningOfTheMonth.ISOWeek()
	return 1 + thisWeek - beginningWeek
}

func concatDateTime(dt time.Time, ts string) (time.Time, error) {
	ds := dt.Format("2006-01-02")
	dts := fmt.Sprintf("%s %s", ds, ts)
	dt, err := time.ParseInLocation("2006-01-02 15:04:05", dts, time.Local)
	if err != nil {
		return dt, fmt.Errorf("parse date time error: %w", err)
	}
	return dt, nil
}

func getNumDayOfWeek(dt time.Time) int {
	num := int(dt.Weekday())
	if num == 0 {
		num = 7
	}
	return num
}

func getNextDateStartJob(day_of_week, repeat int, ts, tag string) (time.Time, error) {

	now := time.Now()
	target, err := concatDateTime(now, ts)
	if err != nil {
		return target, fmt.Errorf("concatDateTime error: %w", err)
	}

	numDayOfWeek := getNumDayOfWeek(now)
	if numDayOfWeek == day_of_week && getNumWeekOfMonth(now) == repeat && now.Unix() < target.Unix() {
		return target, nil
	}
	//s.logger.Debugf("[%s]: %v, DayOfWeek: %d, NumWeek: %d", tag, now, numDayOfWeek, getNumWeekOfMonth(now))
	for {
		now = now.AddDate(0, 0, 1)
		numDayOfWeek := getNumDayOfWeek(now)
		//s.logger.Debugf("[%s]: %v, DayOfWeek: %d, NumWeek: %d", tag, now, numDayOfWeek, getNumWeekOfMonth(now))
		if numDayOfWeek == day_of_week && getNumWeekOfMonth(now) == repeat {
			target, err := concatDateTime(now, ts)
			if err != nil {
				return target, fmt.Errorf("concatDateTime error: %w", err)
			}
			return target, nil
		}
	}
}

func getNextYearStartJob(repeat int, dt *DateTime) (time.Time, error) {
	stamp, err := dt.Parse()
	now := time.Now()
	if err != nil {
		return now, fmt.Errorf("date time parse error: %w", err)
	}

	tn := time.Date(stamp.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), stamp.Location())
	delta := (now.Year() - stamp.Year()) % repeat
	if tn.Unix() > stamp.Unix() {
		delta++
	}
	target := time.Date(now.Year()+delta, stamp.Month(), stamp.Day(), stamp.Hour(), stamp.Minute(), stamp.Second(), 0, stamp.Location())

	return target, nil
}
