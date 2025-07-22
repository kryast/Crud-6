package repositories

import (
	"github.com/kryast/Crud-6.git/models"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(*models.Payment) error
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
