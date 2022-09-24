package entities

import (
	"time"

	"gorm.io/gorm"
)

type TaskParams struct {
	ID     uint32 `gorm:"column:id;not null;unique;index" json:"id"`
	Name   string `gorm:"column:name;not null;size:45" json:"name"`
	TaskID string `gorm:"column:task_id;not null;size:45" json:"task_id"`
}

func (TaskParams) TableName() string {
	return "task_params"
}

type Tasks struct {
	ID     string       `gorm:"column:id;not null;primaryKey;unique;index;size:45" json:"id"`
	Name   string       `gorm:"column:name;not null;unique;index;size:45" json:"name"`
	Note   string       `gorm:"column:note" json:"note"`
	Params []TaskParams `json:"params"`
}

func (Tasks) TableName() string {
	return "tasks"
}

func (t *Tasks) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Delete(&TaskParams{}, "task_id = ?", t.ID).Error
}

type Periods struct {
	ID   string `gorm:"column:id;not null;primaryKey;unique;index;size:16" json:"id"`
	Name string `gorm:"column:name;not null;unique;index;size:45" json:"name"`
	Idx  int    `gorm:"column:idx;not null;unique;index" json:"idx"`
}

func (Periods) TableName() string {
	return "periods"
}

type Days struct {
	ID   uint32 `gorm:"column:id;not null;primaryKey;unique;index" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
}

func (Days) TableName() string {
	return "days"
}

type Executors struct {
	ID string `gorm:"column:id;not null;primaryKey;unique;index;size:20" json:"id"`
}

func (Executors) TableName() string {
	return "executors"
}

type JobParams struct {
	ID    uint32 `gorm:"column:id;not null;primaryKey;unique;index" json:"id"`
	Name  string `gorm:"column:name;not null;size:45" json:"name"`
	Value string `gorm:"column:value;not null" json:"value"`
	JobID uint32 `gorm:"column:job_id;not null" json:"job_id"`
}

func (JobParams) TableName() string {
	return "job_params"
}

type Jobs struct {
	ID       uint32      `gorm:"column:id;not null;primaryKey;unique;index" json:"id"`
	Note     string      `gorm:"column:note;not null;unique" json:"note"`
	Active   int         `gorm:"column:active;not null" json:"active"`
	Value    int         `gorm:"column:value;not null" json:"value"`
	Time     string      `gorm:"column:time;size:45" json:"time"`
	Date     string      `gorm:"column:date;size:45" json:"date"`
	Verbose  int         `gorm:"column:verbose;not null" json:"verbose"`
	Executor string      `gorm:"column:executor_id;not null;size:20" json:"executor_id"`
	TaskID   string      `gorm:"colunm:task_id;not null;size:45" json:"task_id"`
	PeriodID string      `gorm:"column:period_id;not null;size:16" json:"period_id"`
	DayID    int         `gorm:"column:day_id" json:"day_id"`
	Task     Tasks       `json:"task"`
	Period   Periods     `json:"period"`
	Day      Days        `json:"day"`
	Params   []JobParams `json:"params"`
}

func (j *Jobs) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Delete(&JobParams{}, "job_id = ?", j.ID).Error
}

func (Jobs) TableName() string {
	return "jobs"
}

type CronJobs struct {
	Note          string
	ScheduledTime time.Time
}
