package endpoint

import "github.com/gin-gonic/gin"

//GetMainEngine server & routes
func GetMainEngine() *gin.Engine {
	router := gin.Default()
	{
		router.POST("/tasks/:task_id/executions", ExecutionCreate)
		router.POST("tasks/:id", TaskCreate)
		router.GET("tasks/:id", TaskShow)
		router.GET("tasks/", TaskList)
		router.DELETE("tasks/:id", TaskDelete)
	}
	return router
}
