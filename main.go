package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	pwd_validator "goGinPractice/cmd/validator"
	"goGinPractice/config/database"
	"goGinPractice/config/logger"
	user_router "goGinPractice/internal/eventserver/router"
)

func setupLogging() {
	f, _ := os.Create("log.txt")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func main() {

	setupLogging()

	fmt.Println("Hello, World!")
	router := gin.Default()
	router.Use(gin.BasicAuth(gin.Accounts{"sk": "123456"}), logger.Logger())
	binding.Validator.Engine().(*validator.Validate).RegisterValidation("user_pwd", pwd_validator.UserPwd)

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

	go func() {
		database.CreateDBConnection()
	}()

	router.Run(":8000")
}
