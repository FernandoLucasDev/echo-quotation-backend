package controllers

import (
	"api-echo/db"
	"api-echo/models"
	"api-echo/services"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateActivationCode(c echo.Context) error {

	email := c.Param("email")

	rand.Seed(time.Now().UnixNano())
	activationCode := fmt.Sprintf("%06d", rand.Intn(1000000))

	activation := models.Activation{
		Email:      email,
		Activation: activationCode,
	}

	if err := db.DB.Create(&activation).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	body := "Use the following code to authenticate: " + activationCode

	sent := services.SendEmail(email, body)

	if sent {
		return c.JSON(http.StatusAccepted, echo.Map{
			"message": "Activation code was sent",
		})
	}

	return c.JSON(http.StatusInternalServerError, echo.Map{
		"message": "Was not possible to generate or send activation code",
	})

}