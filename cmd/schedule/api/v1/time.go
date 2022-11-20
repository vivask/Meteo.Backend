package v1

import (
	"fmt"
	"meteo/internal/log"
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

func (d DateTime) Day() (int, error) {
	stamp, err := d.Parse()
	if err != nil {
		return -1, fmt.Errorf("parse error: %w", err)
	}
	return stamp.Day(), nil
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

func concatDateTime(dt time.Time, ts string) (time.Time, error) {
	ds := dt.Format("2006-01-02")
	dts := fmt.Sprintf("%s %s", ds, ts)
	dt, err := time.ParseInLocation("2006-01-02 15:04:05", dts, time.Local)
	if err != nil {
		return dt, fmt.Errorf("parse date time error: %w", err)
	}
	return dt, nil
}

func getWeekDateStartJob(repeat int, dt *DateTime) (time.Time, error) {

	stamp, err := dt.Parse()
	if err != nil {
		return stamp, fmt.Errorf("stamp error: %w", err)
	}

	now := time.Now()
	if now.Unix() < stamp.Unix() {
		return stamp, nil
	}

	addDays := repeat * 7
	next := stamp.AddDate(0, 0, addDays)
	for {
		if next.Unix() > now.Unix() {
			break
		}
		next = next.AddDate(0, 0, addDays)
	}
	target, err := concatDateTime(next, dt.ts)
	if err != nil {
		return target, fmt.Errorf("concatDateTime error: %w", err)
	}
	return target, nil
}

func getYearDateStartJob(repeat int, dt *DateTime) (time.Time, error) {

	stamp, err := dt.Parse()
	if err != nil {
		return stamp, fmt.Errorf("stamp error: %w", err)
	}

	now := time.Now()
	if now.Unix() < stamp.Unix() {
		log.Infof("Start At stamp: %v", stamp)
		return stamp, nil
	}

	next := stamp.AddDate(repeat, 0, 0)
	for {
		if next.Unix() > now.Unix() {
			break
		}
		next = next.AddDate(repeat, 0, 0)
	}
	target, err := concatDateTime(next, dt.ts)
	if err != nil {
		return target, fmt.Errorf("concatDateTime error: %w", err)
	}
	log.Infof("Start At target: %v", target)
	return target, nil
}
