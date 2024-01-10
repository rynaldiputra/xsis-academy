package entities

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Image       string  `json:"image"`
}
