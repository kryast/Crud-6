package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CustomerID uint           `json:"customer_id" binding:"required"`
	Customer   Customer       `json:"customer" gorm:"foreignKey:CustomerID"`
	Status     string         `json:"status" binding:"required,oneof=pending paid"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
	// OrderItems []OrderItem `json:"order_items"`
	// Payment    Payment    `json:"payment"`
}
