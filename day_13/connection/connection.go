package connection

import (
	"project/models"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func DatabaseConnect() {
	database, err := gorm.Open(postgres.Open("postgresql://postgres:ydHzocUo6iAAIajoLJMv@containers-us-west-106.railway.app:5616/railway"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Project{})

	DB = database
}