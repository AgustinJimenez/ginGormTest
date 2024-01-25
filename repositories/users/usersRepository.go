package repositories

import (
	"go_practice/initializers"
	"go_practice/models"
	"go_practice/utils"

	"gorm.io/gorm"
)

type CreateUserDataType struct {
	Name string `binding:"required"`
	Username string `binding:"required"`
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func CreateUser(data CreateUserDataType)(models.User, *gorm.DB){
	var user models.User
	user.Name = data.Name
	user.Username = data.Username
	user.Email = data.Email
	err := user.HashPassword(data.Password)
	utils.CheckError(err)
	result := initializers.DB.Create(&user)
	utils.CheckError(result.Error)
	return user, result
}