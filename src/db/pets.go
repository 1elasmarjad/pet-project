package db

import (
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model

	ID int `json:"id" gorm:"primaryKey;autoIncrement"`

	Species     string `json:"species" gorm:"not null"`
	Breed       string `json:"breed" gorm:"not null"`
	Name        string `json:"name" gorm:"not null"`
	MonthsLived int    `json:"months_lived" gorm:"default:null"`
	IsOwned     bool   `json:"isOwned" gorm:"default:null"`
	Image       string `json:"imageURL" gorm:"default:null"`
}
