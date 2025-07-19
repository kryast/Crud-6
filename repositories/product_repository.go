package repositories

import (
	"github.com/kryast/Crud-6.git/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(*models.Product) error
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
