package repositories

import (
	"github.com/kryast/Crud-6.git/models"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(*models.Payment) error
	FindAll() ([]models.Payment, error)
	FindByID(id uint) (*models.Payment, error)
	Update(*models.Payment) error
}

type paymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepo{db}
}

func (r *paymentRepo) Create(p *models.Payment) error {
	return r.db.Create(p).Error
}

func (r *paymentRepo) FindAll() ([]models.Payment, error) {
	var out []models.Payment
	err := r.db.Preload("Order").Find(&out).Error
	return out, err
}

func (r *paymentRepo) FindByID(id uint) (*models.Payment, error) {
	var out models.Payment
	err := r.db.Preload("Order").First(&out, id).Error
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (r *paymentRepo) Update(p *models.Payment) error {
	return r.db.Save(p).Error
}
