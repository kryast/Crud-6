package repositories

import (
	"gorm.io/gorm"
)

type OrderRepository interface {
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepo{db}
}
