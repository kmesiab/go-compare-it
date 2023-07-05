package controller

import "github.com/gin-gonic/gin"

func HealthcheckController(c *gin.Context) {
	c.JSON(200, gin.H{
		"service": "online",
	})
}
