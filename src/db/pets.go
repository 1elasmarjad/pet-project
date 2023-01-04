package db

import (
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Species     string `json:"species" gorm:"not null;default:null"`
	Breed       string `json:"breed" gorm:"not null;default:null"`
	Name        string `json:"name" gorm:"not null;default:null"`
	MonthsLived uint   `json:"months_lived" gorm:"default:null"`
	IsOwned     bool   `json:"isOwned" gorm:"default:null"`
	Image       string `json:"imageURL" gorm:"default:null"`
}
