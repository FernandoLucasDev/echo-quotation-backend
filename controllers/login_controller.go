package controllers

import (
	"net/http"
	"time"
	"api-echo/db"
	"api-echo/models"
	"github.com/labstack/echo/v4"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("supersecreta")

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var loginReq LoginRequest

	if err := c.Bind(&loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request",
		})
	}

	var user models.User

	if result := db.DB.Where("email = ?", loginReq.Email).First(&user); result.Error != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Invalid email or password",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Invalid email or password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 2160).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error generating token",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": tokenString,
	})
}
