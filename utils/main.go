package utils

import (
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

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