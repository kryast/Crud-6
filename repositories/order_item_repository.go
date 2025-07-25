package repositories

import (
	"github.com/kryast/Crud-6.git/models"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	Create(*models.OrderItem) error
	FindAll() ([]models.OrderItem, error)
	FindByID(id uint) (*models.OrderItem, error)
	Update(*models.OrderItem) error
	Delete(id uint) error
	SumSubtotalByOrderID(orderID uint) (float64, error)
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

func (r *orderItemRepo) Update(it *models.OrderItem) error {
	return r.db.Save(it).Error
}

func (r *orderItemRepo) Delete(id uint) error {
	return r.db.Delete(&models.OrderItem{}, id).Error
}

func (r *orderItemRepo) SumSubtotalByOrderID(orderID uint) (float64, error) {
	var total float64
	err := r.db.Model(&models.OrderItem{}).Where("order_id = ?", orderID).Select("COALESCE(SUM(subtotal),0)").Scan(&total).Error
	return total, err
}
