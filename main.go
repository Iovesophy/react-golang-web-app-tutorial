package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("starting gin server")
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "world"})
		})
		api.GET("/world", func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "hello"})
		})
	}
	router.NoRoute(func(c *gin.Context) { c.JSON(http.StatusNotFound, gin.H{}) })
	router.Run(":8080")
}
