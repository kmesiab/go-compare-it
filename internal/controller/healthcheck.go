package controller

import "github.com/gin-gonic/gin"

func Healthcheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"service": "online",
	})
}
