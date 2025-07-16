package repositories

import (
	"gorm.io/gorm"
)

type OrderItemRepository interface {
}

type orderItemRepo struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepo{db}
}
