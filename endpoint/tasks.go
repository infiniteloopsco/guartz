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
	models.Gdb.Where("cron_id == 0").Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

//TaskShow serves the route GET /tasks/:id
func TaskShow(c *gin.Context) {
	var task models.Task
	models.Gdb.First(&task, c.Param("id"))
	if task.ID != "" {
		c.JSON(http.StatusOK, task)
	} else {
		c.JSON(http.StatusNotFound, "")
	}
}

//TaskCreate serves the route POST /tasks
func TaskCreate(c *gin.Context) {
	models.InTx(func(txn *gorm.DB) bool {
		var task models.Task
		if err := c.BindJSON(&task); err == nil {
			if valid, errMap := models.ValidStruct(&task); valid {
				if err := txn.Save(&task); err == nil {
					c.JSON(http.StatusOK, task)
					return true
				} else {
					c.JSON(http.StatusBadRequest, "Couldn't create the task")
				}
			} else {
				c.JSON(http.StatusConflict, errMap)
			}
		}
		return false
	})
}

//TaskDelete serves the route DELETE /tasks/:id
func TaskDelete(c *gin.Context) {
	models.InTx(func(txn *gorm.DB) bool {
		var task models.Task
		models.Gdb.First(&task, c.Param("id"))
		if task.ID != "" {
			if err := txn.Delete(&task).Error; err == nil {
				c.JSON(http.StatusOK, task)
				return true
			} else {
				c.JSON(http.StatusBadRequest, "Could not delete the task")
			}
		} else {
			c.JSON(http.StatusNotFound, "")
		}
		return false
	})
}
