package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	OrderID       uint           `json:"order_id" binding:"required"`
	Order         Order          `json:"order" gorm:"foreignKey:OrderID"`
	PaymentMethod string         `json:"payment_method" binding:"required,oneof=cash credit_card ewallet"`
	Amount        float64        `json:"amount"`
	Status        string         `json:"status" binding:"required,oneof=pending paid failed"`
	PaidAt        *time.Time     `json:"paid_at"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}
