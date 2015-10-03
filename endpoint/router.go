package endpoint

import "github.com/gin-gonic/gin"

//GetMainEngine server & routes
func GetMainEngine() *gin.Engine {
	router := gin.Default()
	root := router.Group("")
	{
		root.POST("tasks/:task_id/executions", ExecutionCreate)
		root.GET("tasks/:task_id/executions", ExecutionList)
		root.POST("tasks", TaskCreate)
		root.GET("tasks/:task_id", TaskShow)
		root.GET("tasks/", TaskList)
		root.DELETE("tasks/:task_id", TaskDelete)
	}
	return router
}
