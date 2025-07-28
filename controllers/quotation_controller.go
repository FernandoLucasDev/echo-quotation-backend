package controllers

import (
	"api-echo/db"
	"api-echo/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetQuotations(c echo.Context) error {
	var quotations []models.Quotation
	if err := db.DB.Find(&quotations).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error fetching quotations",
		})
	}
	return c.JSON(http.StatusOK, quotations)
}
