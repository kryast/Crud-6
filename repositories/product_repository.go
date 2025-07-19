package repositories

import (
	"github.com/kryast/Crud-6.git/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(*models.Product) error
	FindAll() ([]models.Product, error)
	FindByID(id uint) (*models.Product, error)
	Update(*models.Product) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db}
}

func (r *productRepo) Create(p *models.Product) error {
	return r.db.Create(p).Error
}

func (r *productRepo) FindAll() ([]models.Product, error) {
	var out []models.Product
	return out, r.db.Find(&out).Error
}

func (r *productRepo) FindByID(id uint) (*models.Product, error) {
	var out models.Product
	err := r.db.First(&out, id).Error
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (r *productRepo) Update(p *models.Product) error {
	return r.db.Save(p).Error
}
