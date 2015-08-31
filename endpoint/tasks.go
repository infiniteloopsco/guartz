package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infiniteloopsco/guartz/models"
	"github.com/jinzhu/gorm"
)

//TaskList serves the route GET /tasks
func TaskList(c *gin.Context) {
	var tasks []models.Task
	models.Gdb.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

//TaskShow serves the route GET /tasks/:task_id
func TaskShow(c *gin.Context) {
	var task models.Task
	models.Gdb.Where("id like ?", c.Param("task_id")).First(&task)
	if task.ID == "" {
		c.JSON(http.StatusNotFound, "")
	} else {
		c.JSON(http.StatusOK, task)
	}
}

//TaskCreate serves the route POST /tasks
func TaskCreate(c *gin.Context) {
	models.InTx(func(txn *gorm.DB) bool {
		var task models.Task
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return false
		}
		if valid, errMap := models.ValidStruct(&task); !valid {
			c.JSON(http.StatusConflict, errMap)
			return false
		}
		var taskExistent models.Task
		models.Gdb.Where("id like ?", task.ID).First(&taskExistent)
		var err error
		if task.ID != "" && taskExistent.ID != "" {
			taskExistent.Periodicity = task.Periodicity
			taskExistent.Command = task.Command
			err = txn.Save(&taskExistent).Error
		} else {
			err = txn.Create(&task).Error
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, "Couldn't create the task")
			return false
		}
		c.JSON(http.StatusOK, task)
		return true
	})
}

//TaskDelete serves the route DELETE /tasks/:task_id
func TaskDelete(c *gin.Context) {
	models.InTx(func(txn *gorm.DB) bool {
		var task models.Task
		models.Gdb.Where("id like ?", c.Param("task_id")).First(&task)
		if task.ID == "" {
			c.JSON(http.StatusNotFound, "")
			return false
		}
		if err := txn.Delete(&task).Error; err != nil {
			c.JSON(http.StatusBadRequest, "Could not delete the task")
			return false
		}
		c.JSON(http.StatusOK, task)
		return true
	})
}
