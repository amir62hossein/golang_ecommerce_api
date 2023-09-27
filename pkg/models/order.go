package models

import (
	"time"

	"github.com/google/uuid"
)

type ORDER_STATUS string

const (
	PENDING ORDER_STATUS = "PENDING"
	DONE    ORDER_STATUS = "DONE"
	ABORT   ORDER_STATUS = "ABORT"
)

type Orders struct {
	DBModel
	CustomerID uuid.UUID    `json:"customer_id" gorm:"not null"`
	OrderDate  time.Time    `json:"order_date" gorm:"not null"`
	Status     ORDER_STATUS `json:"status" sql:"type:ENUM('PENDING', 'DONE', 'ABORT')"`
	ProductID  uuid.UUID    `json:"product_id" gorm:"type:uuid;not null"`
	Quantity   uint         `json:"quantity" gorm:"not null"`
}
