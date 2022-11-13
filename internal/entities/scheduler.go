package entities

import (
	"time"

	"gorm.io/gorm"
)

type TaskParams struct {
	ID     uint32 `gorm:"column:id;primaryKey" json:"id"`
	Name   string `gorm:"column:name;not null;size:45" json:"name"`
	Note   string `gorm:"column:note" json:"note"`
	TaskID string `gorm:"column:task_id;not null;size:45" json:"task_id"`
}

func (TaskParams) TableName() string {
	return "task_params"
}

type Tasks struct {
	ID     string       `gorm:"column:id;primaryKey;size:45" json:"id"`
	Name   string       `gorm:"column:name;not null;unique;size:45" json:"name"`
	Api    string       `gorm:"column:api;not null" json:"api"`
	Note   string       `gorm:"column:note" json:"note"`
	Params []TaskParams `gorm:"foreignKey:TaskID;references:ID" json:"params"`
}

func (Tasks) TableName() string {
	return "tasks"
}

func (t *Tasks) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Delete(&TaskParams{}, "task_id = ?", t.ID).Error
}

type Periods struct {
	ID   string `gorm:"column:id;primaryKey;size:16" json:"id"`
	Name string `gorm:"column:name;not null;unique;size:45" json:"name"`
	Idx  int    `gorm:"column:idx;not null;unique" json:"idx"`
}

func (Periods) TableName() string {
	return "periods"
}

type Days struct {
	ID   uint32 `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
}

func (Days) TableName() string {
	return "days"
}

type Executors struct {
	ID string `gorm:"column:id;primaryKey;size:20" json:"id"`
}

func (Executors) TableName() string {
	return "executors"
}

type JobParams struct {
	ID    uint32 `gorm:"column:id;primaryKey" json:"id"`
	Name  string `gorm:"column:name;not null;size:45" json:"name"`
	Value string `gorm:"column:value;not null" json:"value"`
	JobID uint32 `gorm:"column:job_id;not null" json:"job_id"`
}

func (JobParams) TableName() string {
	return "job_params"
}

type Jobs struct {
	ID         uint32      `gorm:"column:id;primaryKey" json:"id"`
	Note       string      `gorm:"column:note;not null;unique" json:"note"`
	Active     bool        `gorm:"column:active;not null" json:"active"`
	Value      int         `gorm:"column:value;not null" json:"value"`
	Time       string      `gorm:"column:time;size:45" json:"time"`
	Date       string      `gorm:"column:date;size:45" json:"date"`
	Verbose    bool        `gorm:"column:verbose;not null" json:"verbose"`
	ExecutorID string      `json:"-"`
	Executor   Executors   `gorm:"foreignkey:ExecutorID" json:"executor"`
	TaskID     string      `json:"-"`
	Task       Tasks       `gorm:"foreignkey:TaskID" json:"task"`
	PeriodID   string      `json:"-"`
	Period     Periods     `gorm:"foreignkey:PeriodID" json:"period"`
	DayID      int         `json:"-"`
	Day        Days        `gorm:"foreignkey:DayID" json:"day"`
	Params     []JobParams `gorm:"foreignKey:JobID;references:ID" json:"params"`
}

func (j *Jobs) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Delete(&JobParams{}, "job_id = ?", j.ID).Error
}

func (Jobs) TableName() string {
	return "jobs"
}

type CronJobs struct {
	Note          string    `json:"note"`
	ScheduledTime time.Time `json:"time"`
}
