package models

import "time"

type Quotation struct {
	Base     string `json:"base"`
	Name	 string `json:"name"`
	Value   string `json:"value"`
	Date   time.Time `json:"date"`
}