package controllers

import (
	"api-echo/db"
	"api-echo/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-contrib/session"
)

type LoginRequest struct {
	Email      string `json:"email"`
	Activation string `json:"activation"`
}

func Login(c echo.Context) error {
	var loginReq LoginRequest

	if err := c.Bind(&loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Requisição inválida",
		})
	}

	var user models.User
	var activation models.Activation

	if result := db.DB.Where("email = ?", loginReq.Email).First(&user); result.Error != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Email ou código inválido",
		})
	}

	if result := db.DB.Where("email = ?", loginReq.Email).First(&activation); result.Error != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Email ou código inválido",
		})
	}

	if activation.Activation != loginReq.Activation {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Código de ativação incorreto",
		})
	}

	sess, _ := session.Get("session", c)
	sess.Values["authenticated"] = true
	sess.Values["user_id"] = user.ID
	sess.Values["email"] = user.Email
	sess.Save(c.Request(), c.Response())

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login realizado com sucesso",
	})
}
