package controllers

import (
	"api-echo/db"
	"api-echo/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetNews(c echo.Context) error {
	var news []models.News
	db.DB.Find(&news)
	return c.JSON(http.StatusOK, news)
}