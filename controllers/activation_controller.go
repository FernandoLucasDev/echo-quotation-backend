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

type ActivationRequest struct {
	Email string `json:"email"`
}

func CreateActivationCode(c echo.Context) error {

	var req ActivationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request",
		})
	}

	email := req.Email

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