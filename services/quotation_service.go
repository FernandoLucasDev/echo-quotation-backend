package services

import (
	"api-echo/db"
	"api-echo/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type QuotationResponse struct {
	Date  string            `json:"date"`
	Base  string            `json:"base"`
	Rates map[string]string `json:"rates"`
}

func FetchQuotation() {

	baseURL := os.Getenv("QUOTATION_API_URL")
	apiKey := os.Getenv("QUOTATION_API_KEY")

	if baseURL == "" || apiKey == "" {
		fmt.Println("Not able to get environment variables for quotation API")
		return
	}

	url := fmt.Sprintf("%s%s", baseURL, apiKey)

	resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error requesting quotations:", err)
        return
    }
    defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Erro ao ler o corpo:", err)
        return
    }

	var data QuotationResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}

	db.DB.Where("1 = 1").Delete(&models.Quotation{})

	for moeda, valor := range data.Rates {
		quotation := models.Quotation{
			Base:  data.Base,
			Name:  moeda,
			Value: valor,
			Date:  time.Now(),
		}

		if err := db.DB.Create(&quotation).Error; err != nil {
			fmt.Printf("Error saving currency quotation %s: %v\n", moeda, err)
		} else {
			fmt.Printf("Quotation saved %s successfully\n", moeda)
		}
	}

}