package tests

import (
	"go_practice/initializers"
	"go_practice/setup"

	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker"
)

var Fake = faker.New()
var TestApp *gin.Engine

func ResetApp() {
	TestApp = setup.GetApp()
}

func BeginDbTransaction() {
	initializers.DB = initializers.DB.Begin()
}

func RollbackDbTransaction() {
	initializers.DB.Rollback()
}

func init() {
    ResetApp()
	Fake = faker.New()
}