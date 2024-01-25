package controllers

import (
	"go_practice/initializers"
	"go_practice/libs"
	"go_practice/models"
	repositories "go_practice/repositories/users"
	"go_practice/types"
	"go_practice/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterUserRequestType struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UsersIndex(c *gin.Context){
	var users []models.User

	
	initializers.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"posts": users,
	  })
}

func RegisterUser(c *gin.Context) {
	var body RegisterUserRequestType
	c.Bind(&body)
	user, record := repositories.CreateUser(repositories.CreateUserDataType{
		Name: body.Name,
		Username: body.Username,
		Email: body.Email,
		Password: body.Password,
	}) 
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}
	response := types.RegisterUserResponseType{
		Email: user.Email,
		Id: user.ID,
		Username: user.Username,
	}
	c.JSON(http.StatusCreated, response)
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token    string `json:"token"`
}

func LoginUser(context *gin.Context) {
	var request LoginRequest
	var user models.User

	if err := utils.BindJsonReq(context, &request); err != nil {
        return 
    }


	record := initializers.DB.First(&user,  "email = ?", request.Email)

	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err:= libs.GenerateJWT(user.Email, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, LoginResponse{Token: tokenString})
}