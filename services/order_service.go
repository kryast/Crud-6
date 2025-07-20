package services

import (
	"github.com/kryast/Crud-6.git/models"
	"github.com/kryast/Crud-6.git/repositories"
)

type OrderService interface {
	Create(*models.Order) error
	FindAll() ([]models.Order, error)
	FindByID(id uint) (*models.Order, error)
	Update(*models.Order) error
}

type orderService struct{ repo repositories.OrderRepository }

func NewOrderService(r repositories.OrderRepository) OrderService {
	return &orderService{r}
}

func (s *orderService) Create(o *models.Order) error            { return s.repo.Create(o) }
func (s *orderService) FindAll() ([]models.Order, error)        { return s.repo.FindAll() }
func (s *orderService) FindByID(id uint) (*models.Order, error) { return s.repo.FindByID(id) }
func (s *orderService) Update(o *models.Order) error            { return s.repo.Update(o) }
