package services

import (
	"github.com/kryast/Crud-6.git/models"
	"github.com/kryast/Crud-6.git/repositories"
)

type CustomerService interface {
	Create(*models.Customer) error
	FindAll() ([]models.Customer, error)
	FindByID(id uint) (*models.Customer, error)
	Update(*models.Customer) error
}

type customerService struct {
	repo repositories.CustomerRepository
}

func NewCustomerService(r repositories.CustomerRepository) CustomerService {
	return &customerService{r}
}

func (s *customerService) Create(c *models.Customer) error            { return s.repo.Create(c) }
func (s *customerService) FindAll() ([]models.Customer, error)        { return s.repo.FindAll() }
func (s *customerService) FindByID(id uint) (*models.Customer, error) { return s.repo.FindByID(id) }
func (s *customerService) Update(c *models.Customer) error            { return s.repo.Update(c) }
