package repositories

import (
	"github.com/kryast/Crud-6.git/models"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	Create(*models.OrderItem) error
}

type orderItemRepo struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepo{db}
}

func (r *orderItemRepo) Create(it *models.OrderItem) error {
	return r.db.Create(it).Error
}
