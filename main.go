package main

import (
	"api-echo/db"
	"api-echo/routes"
	"api-echo/services"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-contrib/session"
	"github.com/robfig/cron/v3"
	"net/http"

	"log"
	"github.com/joho/godotenv"
)

func main() {

	db.Connect()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	e := echo.New()
	c := cron.New()

	fourMonths := 4 * 30 * 24 * time.Hour

	store := sessions.NewCookieStore([]byte("authkey"))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int(fourMonths.Seconds()),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	e.Use(session.Middleware(store))

	_, err = c.AddFunc("0 8,16,20 * * *", func() {
		log.Println("🔁 Calling news job...")
		services.FetchAndStoreNews()
	})
	if err != nil {
		log.Fatal("Error scheduling job:", err)
	}

	_, err = c.AddFunc("0 10,12,14,16,18,20 * * *", func() {
		log.Println("🔁 Calling quotation job...")
		services.FetchQuotation()
	})
	if err != nil {
		log.Fatal("Error scheduling quotation job:", err)
	}
	
	c.Start()

	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))

}
