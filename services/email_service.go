package services

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	gomail "gopkg.in/mail.v2"
)

func SendEmail(to string, body string) bool {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	subject := "Email via Gmail"

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
		return false
	}

	fmt.Println("E-mail enviado com sucesso!")

	return true

}
