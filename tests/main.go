package tests

import (
	"go_practice/initializers"
	"go_practice/setup"
	"go_practice/utils"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker"
)

var Fake = faker.New()
var TestApp *gin.Engine

func ResetApp() {
	TestApp = setup.GetApp()
}

func TestHttpRequest(method string, url string, body io.Reader) (*httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, body)
	utils.CheckError(err)
	req.Header.Set("Content-Type", "application/json")
	TestApp.ServeHTTP(w, req)
	return w
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