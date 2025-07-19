package services

import (
	"github.com/kryast/Crud-6.git/models"
	"github.com/kryast/Crud-6.git/repositories"
)

type ProductService interface {
	Create(*models.Product) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(r repositories.ProductRepository) ProductService {
	return &productService{r}
}

func (s *productService) Create(p *models.Product) error { return s.repo.Create(p) }
