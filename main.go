package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	user_router "goGinPractice/internal/eventserver/router"
)

func main() {
	fmt.Println("Hello, World!")
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/post/:post_id", func(c *gin.Context) {
		post_id := c.Param("post_id")
		c.JSON(200, gin.H{
			"msg": "hi " + post_id,
		})
	})

	v1 := router.Group("/v1")
	user_router.AddUserRouter(v1)

	router.Run(":8000")
}
