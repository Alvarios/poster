package _setup

import (
	"github.com/gin-gonic/gin"
)

func ping (c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func postJSON(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "posted",
	})
}

func SetupServer() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", ping)
	r.POST("/post_json", postJSON)
	r.GET("/get_json", postJSON)

	return r
}
