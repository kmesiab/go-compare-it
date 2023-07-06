package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kmesiab/go-compare-it/internal/controller"
)

const currentVersionPath = "/api/v1/"

// apiRoutes is the router group based on the root
// path defined in router.currentVersionPath
var apiRoutes *gin.RouterGroup

// Init attaches all the application routes to the
// gin engine
func Init(router *gin.Engine) {

	// Infrastructure Routes
	router.GET("/healthcheck", controller.Healthcheck)

	if apiRoutes == nil {
		apiRoutes = router.Group(currentVersionPath)
	}

	InitApiRoutes(apiRoutes)
}

// InitApiRoutes attaches all API endpoints to the
// current api's router group
func InitApiRoutes(router *gin.RouterGroup) {
	router.POST("/api/v1/users/create", controller.CreateUser)

}
