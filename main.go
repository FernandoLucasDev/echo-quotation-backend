package main

import (
	"api-echo/db"
	"api-echo/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db.Connect()
	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}
