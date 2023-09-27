package models

import "github.com/google/uuid"

type Product struct {
	DBModel
	Name        string    `json:"name" gorm:"size:255;not null"`
	Description string    `json:"description" gorm:"not null"`
	Price       string    `json:"price" gorm:"not null"`
	CategoryID  uuid.UUID `json:"category_id" gorm:"type:uuid;not null"`
}
