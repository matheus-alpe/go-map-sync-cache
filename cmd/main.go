package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matheus-alpe/go-map-sync-cache/internal/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/fibonacci", controllers.FibonacciHandler)

	r.Run(":8080")
}
