package repositories

import (
	"github.com/kryast/Crud-6.git/models"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	Create(*models.OrderItem) error
	FindAll() ([]models.OrderItem, error)
	FindByID(id uint) (*models.OrderItem, error)
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

func (r *orderItemRepo) FindAll() ([]models.OrderItem, error) {
	var out []models.OrderItem
	err := r.db.Preload("Order").Preload("Product").Find(&out).Error
	return out, err
}

func (r *orderItemRepo) FindByID(id uint) (*models.OrderItem, error) {
	var out models.OrderItem
	err := r.db.Preload("Order").Preload("Product").First(&out, id).Error
	if err != nil {
		return nil, err
	}
	return &out, nil
}
