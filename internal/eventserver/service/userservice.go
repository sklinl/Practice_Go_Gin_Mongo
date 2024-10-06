package service

import (
	"goGinPractice/internal/eventserver/dto"

	"github.com/gin-gonic/gin"
)

var userList = []dto.User{}

func GetAllUsers(c *gin.Context) {
	c.JSON(200, userList)
}

func CreateUser(c *gin.Context) {

	var user dto.User
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userList = append(userList, user)
	c.JSON(200, user)
}
