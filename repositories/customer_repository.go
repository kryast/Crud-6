package repositories

import (
	"github.com/kryast/Crud-6.git/models"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(*models.Customer) error
	FindAll() ([]models.Customer, error)
	FindByID(id uint) (*models.Customer, error)
	Update(*models.Customer) error
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

func (r *customerRepo) FindAll() ([]models.Customer, error) {
	var out []models.Customer
	return out, r.db.Find(&out).Error
}

func (r *customerRepo) FindByID(id uint) (*models.Customer, error) {
	var out models.Customer
	err := r.db.First(&out, id).Error
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (r *customerRepo) Update(c *models.Customer) error {
	return r.db.Save(c).Error
}
