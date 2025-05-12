package controllers

import (
	"api-echo/db"
	"api-echo/models"
)

func CreateActivationCode(email string, code string) bool {
	activation := models.Activation{
		Email:      email,
		Activation: code,
	}

	if err := db.DB.Create(&activation).Error; err != nil {
		return false
	}

	return true
}
