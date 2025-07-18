package repositories

import (
	"github.com/kryast/Crud-6.git/models"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(*models.Customer) error
}

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepo{db}
}

func (r *customerRepo) Create(c *models.Customer) error {
	return r.db.Create(c).Error
}
