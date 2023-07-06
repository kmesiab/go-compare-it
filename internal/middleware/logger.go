package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.
			// before request
			c.Next()
		// after request

		latency := time.Since(t)
		log.Printf("latency: %d", latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
