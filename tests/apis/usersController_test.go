package tests

import (
	"encoding/json"
	"go_practice/controllers"
	"go_practice/models"
	repositories "go_practice/repositories/users"
	"go_practice/tests"
	"go_practice/types"
	"go_practice/utils"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)
var defaultTestPassword = "TestTing123."
var testUser = models.User{
	Name: "Rand Name",
	Username: "Rand Username",
	Email: "RandEmail@mail.com",
	Password: defaultTestPassword,
}

func TestRegisterUser(t *testing.T){
	tests.ResetApp()
	tests.BeginDbTransaction()
	testCases := []models.User{testUser}
	for _, newPostPayload := range testCases {
		res := tests.TestHttpRequest("POST", "/register", utils.GenPayload(newPostPayload))
		assert.Equal(t, http.StatusCreated, res.Code)

		var response types.RegisterUserResponseType
		err := json.Unmarshal(res.Body.Bytes(), &response)
		utils.CheckTestError(t, err)

		assert.Equal(t, newPostPayload.Username, response.Username )
		assert.Equal(t, newPostPayload.Email, response.Email )
		assert.NotEmpty(t, response.Id)
		assert.Equal(t, 3, utils.CountFields(response) )
	}
	tests.RollbackDbTransaction()
}

func TestLoginUser(t *testing.T){
	tests.ResetApp()
	tests.BeginDbTransaction()

	_, record := repositories.CreateUser(
		repositories.CreateUserDataType{
			Name: testUser.Name,
			Username: testUser.Username,
			Email: testUser.Email,
			Password: testUser.Password,
		},
	)
	utils.CheckTestError(t, record.Error)

	testCases := []gin.H{
		{
			"email": testUser.Email,
			"password": testUser.Password,
		},
		{
			"email": testUser.Email,
			"password": nil,
		},
		{
			"email": nil,
			"password": testUser.Password,
		},
		{
			"email": nil,
			"password": nil,
		},
	}

	for _, userPayload := range testCases {
		println("\n\n TESTING EMAILS!!! EMAIL=", userPayload.get("email"), ", PASSWORD=", userPayload["password"])
		if userPayload["email"] != nil && userPayload["password"] != nil {
			res := tests.TestHttpRequest("POST", "/login", utils.GenPayload(userPayload))
			println("\n\n FIRST =>  EMAIL=", userPayload["email"], ", PASSWORD=", userPayload["password"])
			assert.Equal(t, http.StatusOK, res.Code)
			var response controllers.LoginResponse
			err := json.Unmarshal(res.Body.Bytes(), &response)
			utils.CheckTestError(t, err)
			assert.NotEmpty(t, response.Token )
		}else if userPayload["email"] == nil || userPayload["password"] == nil { 
			res := tests.TestHttpRequest("POST", "/login", utils.GenPayload(userPayload))
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	}

	tests.RollbackDbTransaction()
} 