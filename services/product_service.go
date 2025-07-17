package services

import "github.com/kryast/Crud-6.git/repositories"

type ProductService interface {
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(r repositories.ProductRepository) ProductService {
	return &productService{r}
}
