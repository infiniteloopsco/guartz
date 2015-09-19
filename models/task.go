package models

import (
	"os/exec"
	"strings"
	"time"

	"github.com/infiniteloopsco/guartz/utils"

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
	utils.Log.Infof("task after create")
	return t.Start(txn)
}

//AfterUpdate callback
func (t *Task) AfterUpdate(txn *gorm.DB) error {
	utils.Log.Infof("task after update")
	if err := t.Stop(txn); err != nil {
		return err
	}
	return t.Start(txn)
}

//BeforeDelete callback
func (t *Task) BeforeDelete(txn *gorm.DB) error {
	return t.Stop(txn)
}

//Start the task
func (t *Task) Start(txn *gorm.DB) error {
	if t.Periodicity == "stop" {
		return txn.Model(t).UpdateColumn("cron_id", 0).Error
	}
	pid, err := MasterCron.AddFunc(t.Periodicity, func() {
		commandArr := strings.Split(t.Command, " ")
		command, args := commandArr[0], commandArr[1:]
		utils.Log.Infof("Running command %s with args: %v", command, args)
		resp, err := exec.Command(command, args...).Output()
		if err != nil {
			utils.Log.Error(err)
		}
		utils.Log.Infof("Output: %s", string(resp))
	})
	if err != nil {
		return err
	}
	utils.Log.Infof("The task %s has been started", t.ID)
	return txn.Model(t).UpdateColumn("cron_id", int(pid)).Error
}

//Stop the task
func (t *Task) Stop(txn *gorm.DB) error {
	if t.CronID != 0 {
		entryID := cron.EntryID(t.CronID)
		MasterCron.Remove(entryID)
		utils.Log.Infof("The task %s has been stopped", t.ID)
		return txn.Model(t).UpdateColumn("cron_id", 0).Error
	}
	return nil
}
