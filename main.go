package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kmesiab/go-compare-it/internal/middleware"
	"github.com/kmesiab/go-compare-it/internal/router"
)

// 0.0.0.0:8080
func main() {
	r := gin.Default()

	router.Init(r)

	err := middleware.RegisterMiddleware("logger", middleware.Logger())
	if err != nil {
		panic(err)
	}

	middleware.Init(r)

	err = r.Run()

	if err != nil {
		panic(err)
	}
}
