package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kmesiab/go-compare-it/internal/controller"
	"net/http"
)

const currentVersionPath = "/api/v1/"
const staticRoute = "/a/1/"

// apiRoutes is the router group based on the root
// path defined in router.currentVersionPath
var apiRoutes *gin.RouterGroup

// Init attaches all the application routes to the
// gin engine
func Init(router *gin.Engine) {

	router.GET("/healthcheck", controller.Healthcheck)
	router.StaticFS(staticRoute, http.Dir("public"))

	if apiRoutes == nil {
		apiRoutes = router.Group(currentVersionPath)
	}

	InitApiRoutes(apiRoutes)
}

// InitApiRoutes attaches all API endpoints to the
// current api's router group
func InitApiRoutes(router *gin.RouterGroup) {
	router.POST("users", controller.CreateUser)
}
