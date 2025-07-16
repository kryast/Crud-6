package repositories

import (
	"gorm.io/gorm"
)

type ProductRepository interface {
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db}
}
