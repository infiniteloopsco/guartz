package models

import (
	"github.com/infiniteloopsco/guartz/utils"

	"github.com/jinzhu/gorm"
	"gopkg.in/robfig/cron.v2"
)

//MasterCron system
var MasterCron *cron.Cron

//InitCron system
func InitCron() {
	MasterCron = cron.New()
	MasterCron.Start()

	//Loads the tasks on the cron system
	var tasks []Task
	Gdb.Find(&tasks)
	utils.Log.Infof("***STARTING %d PROJECTS***\n", len(tasks))
	InTx(func(txn *gorm.DB) bool {
		for _, task := range tasks {
			if err := task.Start(txn); err != nil {
				panic(err)
			}
		}
		return true
	})
}
