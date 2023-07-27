package connection

import (
	"project/models"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func DatabaseConnect() {
	database, err := gorm.Open(postgres.Open("postgres://default:0vWRuhSlz7GN@ep-lingering-bonus-79890093-pooler.ap-southeast-1.postgres.vercel-storage.com:5432/verceldb"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Project{})

	DB = database
}