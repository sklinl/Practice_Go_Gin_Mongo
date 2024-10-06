package router

import (
	"goGinPractice/internal/eventserver/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(router *gin.RouterGroup) {
	users := router.Group("/users")
	users.POST("/", service.CreateUser)
	users.GET("/", service.GetAllUsers)
	users.GET("/:username", service.GetUserByUserNanme)
	users.DELETE("/:username", service.DeleteUser)
	users.PUT("/:username", service.UpdateUser)
}
