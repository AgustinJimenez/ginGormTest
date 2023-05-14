package main

import (
	"fmt"
	"go_practice/initializers"
	"go_practice/models"
)

var modelList = []interface{}{&models.Post{}, &models.User{}}

func main(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	// Migrate the schema
	initializers.DB.AutoMigrate(modelList...)
	d, _ := initializers.DB.DB()
	d.Close()


	fmt.Println("Migration Completed...")
}