package main

import (
	"api-echo/db"
	"api-echo/routes"
	"api-echo/services"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
)

func main() {

	e := echo.New()
	c := cron.New()

	db.Connect()

	_, err := c.AddFunc("*/5 * * * *", func() {
		log.Println("🔁 Calling news job...")
		services.FetchAndStoreNews()
	})
	if err != nil {
		log.Fatal("Error scheduling Job:", err)
	}

	c.Start()

	routes.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
	
}
