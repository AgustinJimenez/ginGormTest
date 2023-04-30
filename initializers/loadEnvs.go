package initializers

import (
	"go_practice/utils"

	"github.com/joho/godotenv"
)

func getEnvPath() string {
	projectPath := utils.GetProjectPath()
	envFileStr := "/.env"
	
	return projectPath+envFileStr
}


func LoadEnvVariables() {
	envFilePath := getEnvPath() 
	err := godotenv.Load(envFilePath)
	utils.CheckError(err)
}