package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" binding:"required"`
	Category  string         `json:"category" binding:"required"`
	Price     float64        `json:"price" binding:"required,gt=0"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
