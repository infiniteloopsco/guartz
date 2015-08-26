package endpoint

import "github.com/gin-gonic/gin"

//GetMainEngine server & routes
func GetMainEngine() *gin.Engine {
	router := gin.Default()
	{
		router.POST("tasks/:task_id/executions", ExecutionCreate)
		router.GET("tasks/:task_id/executions", ExecutionList)
		router.POST("tasks", TaskCreate)
		router.GET("tasks/:task_id", TaskShow)
		router.GET("tasks/", TaskList)
		router.DELETE("tasks/:task_id", TaskDelete)
	}
	return router
}
