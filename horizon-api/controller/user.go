package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAdd(c *gin.Context) {
	detail := gin.H{
		"token": "",
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success",
		"detail": detail,
	})
}
