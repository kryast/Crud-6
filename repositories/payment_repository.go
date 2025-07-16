package repositories

import (
	"gorm.io/gorm"
)

type PaymentRepository interface {
}

type paymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepo{db}
}
