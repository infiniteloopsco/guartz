package models

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/robfig/cron.v2"
)

//MasterCron system
var MasterCron *cron.Cron

//InitCron system
func InitCron() {
	MasterCron = cron.New()
	MasterCron.Start()
	Gdb.Model(Task{}).Where("cron_id != 0").Updates(Task{CronID: 0})

	//Loads the tasks on the cron system
	var tasks []Task
	Gdb.Where("cron_id == 0").Find(&tasks)
	InTx(func(txn *gorm.DB) bool {
		for _, task := range tasks {
			if err := task.Start(txn); err != nil {
				return false
			}
		}
		return true
	})
}
