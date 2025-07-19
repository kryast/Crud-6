package services

import (
	"github.com/kryast/Crud-6.git/models"
	"github.com/kryast/Crud-6.git/repositories"
)

type ProductService interface {
	Create(*models.Product) error
	FindAll() ([]models.Product, error)
	FindByID(id uint) (*models.Product, error)
	Update(*models.Product) error
	Delete(id uint) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(r repositories.ProductRepository) ProductService {
	return &productService{r}
}

func (s *productService) Create(p *models.Product) error            { return s.repo.Create(p) }
func (s *productService) FindAll() ([]models.Product, error)        { return s.repo.FindAll() }
func (s *productService) FindByID(id uint) (*models.Product, error) { return s.repo.FindByID(id) }
func (s *productService) Update(p *models.Product) error            { return s.repo.Update(p) }
func (s *productService) Delete(id uint) error                      { return s.repo.Delete(id) }
