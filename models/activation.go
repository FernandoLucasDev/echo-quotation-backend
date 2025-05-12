package models

type Activation struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email"`
	Activation string `json:"activation"`
}