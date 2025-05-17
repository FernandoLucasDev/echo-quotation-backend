package models

import "time"

type News struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Midia     string `json:"midia"`
	Link      string `json:"link"`
	Image     string `json:"image"`
	CreatedAt time.Time
}
