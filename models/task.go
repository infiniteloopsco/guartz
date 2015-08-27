package models

import (
	"fmt"
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
	CreatedAt   time.Time `json:"created_at"`
}

//BeforeCreate callback
func (t *Task) BeforeCreate() {
	if t.ID == "" {
		t.ID = helpers.PseudoUUID()
	}
}

//AfterCreate callback
func (t *Task) AfterCreate(txn *gorm.DB) error {
	return t.Start(txn)
}

//AfterUpdate callback
func (t *Task) AfterUpdate(txn *gorm.DB) error {
	if err := t.Stop(txn); err != nil {
		return err
	}
	return t.Start(txn)
}

//BeforeDelete callback
func (t *Task) BeforeDelete(txn *gorm.DB) error {
	return t.Stop(txn)
}

func (t *Task) Start(txn *gorm.DB) error {
	pid, err := MasterCron.AddFunc(t.Periodicity, func() {
		commandArr := strings.Split(t.Command, " ")
		command, args := commandArr[0], commandArr[1:]
		exec.Command(command, args...).Run()
	})
	if err != nil {
		return err
	}
	fmt.Println("Cron started")
	return txn.Model(t).UpdateColumn("cron_id", int(pid)).Error
}

func (t *Task) Stop(txn *gorm.DB) error {
	MasterCron.Remove(cron.EntryID(t.CronID))
	return txn.Model(t).UpdateColumn("cron_id", 0).Error
}
