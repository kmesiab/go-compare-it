package main

import "github.com/gin-gonic/gin"
import "github.com/kmesiab/go-compare-it/internal/controller"

func main() {
	r := gin.Default()
	initRoutes(r)
	initMiddleware(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func initRoutes(router *gin.Engine) *gin.Engine {

	// Infrastructure Routes
	router.GET("/ping", controller.PingController)
	router.GET("/healthcheck", controller.HealthcheckController)

	// Application Routes

	// API Routes
	// Current root: /api/v1/
	//
	// r.GET("/api/v1/docs")
	//

	// Create a new user
	// POST /users/create
	router.POST("/api/v1/users/create", controller.CreateUser)

	return router
}

func initMiddleware(router *gin.Engine) *gin.Engine {

	/*
		for _, r := range middleware.GetAll() {
			r.Use(m)
		}
	*/

	return router

}
