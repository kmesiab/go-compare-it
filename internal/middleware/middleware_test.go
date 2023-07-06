package middleware_test

import (
	"github.com/gin-gonic/gin"
	"github.com/kmesiab/go-compare-it/internal/middleware"
	"github.com/stretchr/testify/assert"
	"testing"
)

func mockMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("test", "1")
	}
}
func mockMiddleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("test", "2")
	}
}

func TestRegisterMiddleware(t *testing.T) {

	err := middleware.RegisterMiddleware("mock", mockMiddleware())

	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, len(middleware.GetAllMiddleware()), 1,
		"failed to register middleware")

	// Add a second one
	err = middleware.RegisterMiddleware("mock2", mockMiddleware2())

	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, len(middleware.GetAllMiddleware()), 2,
		"failed to register second middleware")
}

func TestRegisterMiddlewareTwice(t *testing.T) {

	err := middleware.RegisterMiddleware("duplicate mock", mockMiddleware())

	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.GreaterOrEqualf(t, len(middleware.GetAllMiddleware()), 3,
		"failed to add first middleware")

	if len(middleware.GetAllMiddleware()) != 3 {
		t.Fatalf("failed to add first middleware mock")
	}

	err = middleware.RegisterMiddleware("duplicate mock", mockMiddleware())

	assert.NotNil(t, err, "duplicate middleware detected")

	if err.Error() != "middleware 'duplicate mock' has already been loaded" {
		t.Fatalf("failed to report middleware has been added twice")
	}

	assert.GreaterOrEqualf(t, len(middleware.GetAllMiddleware()), 3,
		"failed to prevent middleware from being added twice")
}

func TestAddMiddlewareAfterInit(t *testing.T) {

	middleware.Init(&gin.Engine{})
	err := middleware.RegisterMiddleware("mock", mockMiddleware())

	assert.NotNil(t, err, "failed to return an error when registering middleware late")

	if err.Error() != "cannot add middleware after initialization" {
		t.Fatalf("failed to report error when registering middleware late")
	}
}
