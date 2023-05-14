package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func GetProjectPath() string {
    _, filename, _, _ := runtime.Caller(0)
    projectPath := filepath.Dir(filepath.Dir(filename))
    return projectPath
}

func GetStructFields(s interface{}) map[string]interface{} {
    fields := make(map[string]interface{})
    value := reflect.ValueOf(s).Elem()

    for i := 0; i < value.NumField(); i++ {
        fieldName := value.Type().Field(i).Name
        fieldValue := value.Field(i).Interface()
        fields[fieldName] = fieldValue
    }

    return fields
}

func CountFields(s interface{}) int {
    t := reflect.TypeOf(s)
    return t.NumField()
}

func CheckTestError(t *testing.T, e error) bool {
    if e != nil {
        t.Errorf("Error: %s", e.Error())
        return true;
    }
    return false;
}

func CheckError(e error) bool {
    if e != nil {
        log.Fatal("Error: %s"+ e.Error())
        return true;
    }
    return false;
}

func CheckDbTestError(t *testing.T, result *gorm.DB) bool {
    if result.Error != nil {
        t.Errorf("Error creating data: %s", result.Error.Error())
        return true;
    }
    return false;
}

func GenPayload(data any)*bytes.Buffer {
	jsonData, err := json.Marshal(data)
	CheckError(err)
	return bytes.NewBuffer(jsonData)
}

func BindJsonReq(c *gin.Context, request interface{}) error {
    if err := c.ShouldBindJSON(request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        c.Abort()
        return err
    }
    return nil
}

func ReverseStr(str string)(result string){
    for _,v := range str {
        result = string(v) + result
      }
    return result
}