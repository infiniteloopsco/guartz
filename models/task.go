package models

import (
	"os/exec"
	"strings"
	"time"

	"gopkg.in/robfig/cron.v2"

	"github.com/jinzhu/gorm"
	"github.com/mrkaspa/go-helpers"
)

//Task executed recurrently
type Task struct {
	ID          string    `sql:"type:varchar(100)" gorm:"primary_key" json:"id"`
	Periodicity string    `json:"periodicity" validate:"required"`
	CronID      int       `json:"-"`
	Command     string    `json:"command" validate:"required"`
	Args        string    `json:"args"`
	CreatedAt   time.Time `json:"created_at"`
}

//BeforeCreate callback
func (t *Task) BeforeCreate(txn *gorm.DB) {
	if t.ID == "" {
		t.ID = helpers.PseudoUUID()
	}
}

//AfterCreate callback
func (t *Task) AfterCreate(txn *gorm.DB) error {
	return t.Start(txn)
}

//BeforeDelete callback
func (t *Task) BeforeDelete(txn *gorm.DB) error {
	return t.Stop(txn)
}

func (t *Task) Start(txn *gorm.DB) error {
	pid, err := MasterCron.AddFunc(t.Periodicity, func() {
		args := strings.Split(t.Args, " ")
		exec.Command(t.Command, args...).Run()
	})
	if err != nil {
		return err
	}
	return txn.Model(t).UpdateColumn("cron_id", int(pid)).Error
}

func (t *Task) Stop(txn *gorm.DB) error {
	MasterCron.Remove(cron.EntryID(t.CronID))
	return txn.Model(t).UpdateColumn("cron_id", 0).Error
}
