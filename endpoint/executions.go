package endpoint

import "github.com/gin-gonic/gin"

//ExecutionCreate serves the route POST /tasks/:task_id/executions
func ExecutionCreate(c *gin.Context) {
	// models.InTx(func(txn *gorm.DB) bool {
	// 	projectID, _ := strconv.Atoi(c.Param("project_id"))
	// 	if _, err := models.FindProject(projectID); err == nil {
	// 		var execution models.Execution
	// 		if err := c.BindJSON(&execution); err == nil {
	// 			if err := validator.Validate(&execution); err == nil {
	// 				execution.ProjectID = projectID
	// 				if txn.Save(&execution).Error == nil {
	// 					c.JSON(http.StatusOK, "")
	// 					return true
	// 				} else {
	// 					c.JSON(http.StatusBadRequest, "Execution can't be saved")
	// 				}
	// 			} else {
	// 				c.JSON(http.StatusConflict, err.(validator.ErrorMap))
	// 			}
	// 		}
	// 	} else {
	// 		c.JSON(http.StatusNotFound, "")
	// 	}
	// 	return false
	// })
}
