package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/infiniteloopsco/guartz/models"
	"github.com/jinzhu/gorm"
)

func createTask() models.Task {
	task := models.Task{
		Periodicity: "@every 10s",
		Command:     "curl -X POST --data payload={\"channel\":\"#general\",\"text\":\"EOOO\"} https://hooks.slack.com/services/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",
	}
	models.InTx(func(txn *gorm.DB) bool {
		if txn.Create(&task).Error != nil {
			panic("error creating the task")
		}
		return true
	})
	return task
}

func createExecution(task models.Task) models.Execution {
	execution := models.Execution{TaskID: task.ID}
	models.InTx(func(txn *gorm.DB) bool {
		if txn.Create(&execution).Error != nil {
			panic("error creating the execution")
		}
		return true
	})
	return execution
}

func getBodyJSON(resp *http.Response, i interface{}) {
	if jsonDataFromHTTP, err := ioutil.ReadAll(resp.Body); err == nil {
		if err := json.Unmarshal([]byte(jsonDataFromHTTP), &i); err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func debugResponse(resp *http.Response) {
	contents, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("*****************")
	fmt.Println(string(contents))
	fmt.Println("*****************")
}
