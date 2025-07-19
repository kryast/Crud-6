package services

import (
	"github.com/kryast/Crud-6.git/models"
	"github.com/kryast/Crud-6.git/repositories"
)

type OrderService interface {
	Create(*models.Order) error
}

type orderService struct{ repo repositories.OrderRepository }

func NewOrderService(r repositories.OrderRepository) OrderService {
	return &orderService{r}
}

func (s *orderService) Create(o *models.Order) error { return s.repo.Create(o) }
