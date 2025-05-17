package main

import (
	"api-echo/db"
	"api-echo/routes"
	"api-echo/services"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db.Connect()
	routes.Init(e)
	services.FetchAndStoreNews()
	e.Logger.Fatal(e.Start(":8080"))
}
