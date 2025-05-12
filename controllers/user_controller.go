package controllers

import (
	"net/http"
	"api-echo/db"
	"api-echo/models"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	var users []models.User
	db.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	db.DB.Create(user)
	return c.JSON(http.StatusCreated, user)
}

func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	if result := db.DB.First(&user, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	if result := db.DB.First(&user, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	if result := db.DB.Delete(&user); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error deleting user",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User deleted successfully",
	})
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	if result := db.DB.First(&user, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
		})
	}

	if result := db.DB.Save(&user); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error updating user",
		})
	}

	return c.JSON(http.StatusOK, user)
}
