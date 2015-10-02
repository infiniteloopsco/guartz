package endpoint

import "github.com/gin-gonic/gin"

func AngularFilter(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS, HEAD")
	c.Header("Access-Control-Allow-Headers", "*,x-requested-with,__setXHR_,Content-Type,If-Modified-Since,If-None-Match")
}
