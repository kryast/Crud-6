package services

import (
	"github.com/kryast/Crud-6.git/models"
	"github.com/kryast/Crud-6.git/repositories"
)

type CustomerService interface {
	Create(*models.Customer) error
}

type customerService struct {
	repo repositories.CustomerRepository
}

func NewCustomerService(r repositories.CustomerRepository) CustomerService {
	return &customerService{r}
}

func (s *customerService) Create(c *models.Customer) error {
	return s.repo.Create(c)
}
