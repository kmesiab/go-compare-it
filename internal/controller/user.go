package controller

import "github.com/gin-gonic/gin"

// POST /users/create
func CreateUser(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "not implemented",
	})
}
