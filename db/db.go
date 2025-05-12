package db

import (
	"fmt"
	"api-echo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=postgresql.uhserver.com user=quotation_admin password=b4t4tinha@ dbname=quotation_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Erro ao conectar com banco: %v", err))
	}
	database.AutoMigrate(&models.User{})
	DB = database
}
