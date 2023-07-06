package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

var hasInitialized bool = false
var processors []gin.HandlerFunc

func processorExists(p interface{}) bool {
	for _, e := range processors {
		if &e == p {
			return true
		}
	}
	return false
}

// RegisterMiddleware takes a handler function and adds it to the
// list of processors to be initialized.  Any processors registered
// after Init() is called, are not added to the gin engine.
func RegisterMiddleware(m gin.HandlerFunc) error {

	if hasInitialized {
		return errors.New("cannot add middleware after initialization")
	}

	if !processorExists(m) {
		processors = append(processors, m)
		fmt.Println("Added processor")
	}

	return nil
}

// Init adds all registered processors to the
// `gin.Engine` for runtime
func Init(router *gin.Engine) *gin.Engine {

	hasInitialized = true
	for _, p := range processors {
		router.Use(p)
	}

	return router

}
