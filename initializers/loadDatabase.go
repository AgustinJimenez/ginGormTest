package initializers

import (
	"fmt"

	"go_practice/migrations"
	"go_practice/utils"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB * gorm.DB

func ConnectToDb() {
	var err error
	dbName := os.Getenv("DB_NAME")
	conn_url := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	connDbUrl := fmt.Sprintf("%s dbname=%s", conn_url, dbName)
	println("DB URL ===> "+conn_url)
    DB, err = gorm.Open(postgres.Open(conn_url), &gorm.Config{})
	utils.CheckError(err)
    count  := 0

	DB.Raw("SELECT count(*) FROM pg_database WHERE datname = ?", dbName).Scan(&count)
	if count == 0 {
		sql :=fmt.Sprintf("CREATE DATABASE %s", dbName)
		result := DB.Exec(sql)
		utils.CheckError(result.Error)
	}

	d, _ := DB.DB()
	d.Close()

	DB, err = gorm.Open(postgres.Open(connDbUrl), &gorm.Config{})
	utils.CheckError(err)


	DB.AutoMigrate(
		migrations.MigrationModelList...
	)
}