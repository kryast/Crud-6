package repositories

import (
	"gorm.io/gorm"
)

type CustomerRepository interface {
}

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepo{db}
}
