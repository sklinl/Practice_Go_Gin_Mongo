package service

import (
	"goGinPractice/internal/eventserver/dto"
	// "strconv"

	"github.com/gin-gonic/gin"
)

var userList = []dto.User{}

func GetAllUsers(c *gin.Context) {
	// c.JSON(200, userList)
	userList := dto.FindAllUser()
	c.JSON(200, userList)

}

func GetUserByUserNanme(c *gin.Context) {
	user_name := c.Params.ByName("username")

	user := dto.FindUserByUsername(user_name)
	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {

	var user dto.User
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// userList = append(userList, user)
	user = dto.CreateUser(user)

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	// userIdx, _ := strconv.Atoi(c.Param("user_id"))
	// if userIdx < 0 || userIdx >= len(userList) {
	// 	c.JSON(400, gin.H{"error": "user id not found"})
	// 	return
	// }
	// userList = append(userList[:userIdx], userList[userIdx+1:]...)

	user_name := c.Params.ByName("username")
	ok := dto.DeleteUser(user_name)

	if !ok {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	c.JSON(200, gin.H{"message": "deleted"})

}

func UpdateUser(c *gin.Context) {
	// userIdx, _ := strconv.Atoi(c.Param("user_id"))
	// if userIdx < 0 || userIdx >= len(userList) {
	// 	c.JSON(400, gin.H{"error": "user id not found"})
	// 	return
	// }
	var user dto.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// userList[userIdx] = user

	user = dto.UpdateUser(user)

	c.JSON(200, user)
}
