package models

import "github.com/google/uuid"

type Comments struct {
	DBModel
	ProductID   uuid.UUID `json:"product_id" gorm:"type:uuid;not null"`
	CustomerID  uuid.UUID `json:"customer_id" gorm:"type:uuid;not null"`
	CommentText string    `json:"comment_text" gorm:"not null"`
}
