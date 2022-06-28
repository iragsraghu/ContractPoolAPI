package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/api/v1/pools", poolList)
	router.Run() // listen and serve on
}
