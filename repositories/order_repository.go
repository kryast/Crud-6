package repositories

import (
	"github.com/kryast/Crud-6.git/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(*models.Order) error
	FindAll() ([]models.Order, error)
	FindByID(id uint) (*models.Order, error)
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepo{db}
}

func (r *orderRepo) Create(o *models.Order) error {
	return r.db.Create(o).Error
}

func (r *orderRepo) FindAll() ([]models.Order, error) {
	var out []models.Order
	err := r.db.Preload("Customer").Find(&out).Error
	return out, err
}

func (r *orderRepo) FindByID(id uint) (*models.Order, error) {
	var out models.Order
	err := r.db.Preload("Customer").First(&out, id).Error
	if err != nil {
		return nil, err
	}
	return &out, nil
}
