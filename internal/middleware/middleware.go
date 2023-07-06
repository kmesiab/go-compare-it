package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

var hasInitialized bool = false
var middleware = map[string]gin.HandlerFunc{}

func GetAllMiddleware() map[string]gin.HandlerFunc {
	return middleware
}

func GetMiddleware(name string) gin.HandlerFunc {
	for key, fn := range middleware {
		if key == name {
			return fn
		}
	}
	return nil
}

// RegisterMiddleware takes a handler function and adds it to the
// list of middleware to be initialized.  Any middleware registered
// after Init() are not added to the gin engine.
func RegisterMiddleware(name string, fn gin.HandlerFunc) error {

	if hasInitialized {
		return errors.New("cannot add middleware after initialization")
	}

	if nil != GetMiddleware(name) {
		return errors.New(fmt.Sprintf("middleware '%s' has already been loaded", name))
	}

	middleware[name] = fn
	return nil
}

// Init adds all registered middleware to the gin.Engine for runtime
func Init(r *gin.Engine) {

	hasInitialized = true
	for _, p := range middleware {
		r.Use(p)
	}
}
