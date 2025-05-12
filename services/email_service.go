package services

import (
	"api-echo/controllers"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	gomail "gopkg.in/mail.v2"
)

func SendEmail(to string) error {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	rand.Seed(time.Now().UnixNano())
	activation := fmt.Sprintf("%06d", rand.Intn(1000000))

	ok := controllers.CreateActivationCode(to, activation)

	var body string

	subject := "Email via Gmail"

	body = "Use the following code to authenticate: " + activation
	if !ok {
		body = "Code generation failed. Please try again later."
	}

	from := os.Getenv("EMAIL_FROM")
	password:= os.Getenv("EMAIL_APP_PASSWORD")
	smtp := os.Getenv("EMAIL_SMTP")

	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	dialer := gomail.NewDialer(smtp, 587, from, password)

	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println("Erro ao enviar o e-mail:", err)
		return err
	}

	fmt.Println("E-mail enviado com sucesso!")

	return nil

}
