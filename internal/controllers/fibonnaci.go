package controllers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var cache = &sync.Map{}

func FibonacciHandler(c *gin.Context) {
	n := c.DefaultQuery("n", "0")
	nInt, err := strconv.Atoi(n)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if cacheResult, ok := cache.Load(nInt); ok {
		c.JSON(http.StatusOK, gin.H{"result": cacheResult})
		return
	}

	result := sumFibonacci(nInt)
	cache.Store(nInt, result)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func sumFibonacci(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	var result int

	for i := 0; i <= n; i++ {
		result += a
		a, b = b, a+b
	}

	return result
}
