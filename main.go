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
	routes.Init(e)
	e.Logger.Fatal(e.Start(":8080"))

	_, err := c.AddFunc("0 8,14,20 * * *", func() {
		log.Println("🔁 Executando job de notícias...")
		services.FetchAndStoreNews()
	})
	if err != nil {
		log.Fatal("Erro ao agendar o job:", err)
	}

	c.Start()
	
}
