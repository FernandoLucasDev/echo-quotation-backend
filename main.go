package main

import (
	"github.com/labstack/echo/v4"
	"api-echo/db"
	"api-echo/routes"
)

func main() {
	e := echo.New()
	db.Connect()
	routes.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
