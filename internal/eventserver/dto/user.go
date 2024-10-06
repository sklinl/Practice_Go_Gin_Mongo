package dto

import (
	"goGinPractice/config/database"
	"log"
)

type User struct {
	Username string `json:"username" binding:"gt=2"`
	Email    string `json:"email" `
	Password string `json:"password" binding:"user_pwd"`
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func FindAllUser() []UserResponse {
	var users []User
	var usersResponse []UserResponse

	database.DB_CONNECTOR.Find(&users)

	// users transfer to userResponse
	for _, user := range users {
		var userResponse UserResponse
		userResponse.Username = user.Username
		userResponse.Email = user.Email
		usersResponse = append(usersResponse, userResponse)

	}
	return usersResponse
}

func FindUserByUsername(user_name string) User {
	var user User
	database.DB_CONNECTOR.Where("username = ?", user_name).First(&user)
	if user.Username == "" {
		log.Fatal("user not found")
		return user
	}
	return user
}

func CreateUser(user User) User {
	database.DB_CONNECTOR.Create(&user)
	return user
}

func DeleteUser(user_name string) bool {
	var user User
	result := database.DB_CONNECTOR.Where("username = ?", user_name).Delete(&user)
	return result.RowsAffected > 0
}

func UpdateUser(user User) User {
	database.DB_CONNECTOR.Save(&user)
	return user
}
