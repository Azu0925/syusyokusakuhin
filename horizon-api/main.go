package main

import (
	"horizon-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	apiR := r.Group("/api")
	{
		v1 := apiR.Group("/v1")
		{
			v1.POST("/user", controller.UserAdd)
			v1.GET("/user/:id")
			v1.PUT("/user/:id")
			v1.DELETE("/user/:id")
		}

	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
