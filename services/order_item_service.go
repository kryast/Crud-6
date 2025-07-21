package services

import (
	"github.com/kryast/Crud-6.git/models"
	"github.com/kryast/Crud-6.git/repositories"
)

type OrderItemService interface {
	Create(*models.OrderItem) error
	FindAll() ([]models.OrderItem, error)
	FindByID(id uint) (*models.OrderItem, error)
	Update(*models.OrderItem) error
}

type orderItemService struct {
	repo        repositories.OrderItemRepository
	productRepo repositories.ProductRepository
}

func NewOrderItemService(r repositories.OrderItemRepository, pr repositories.ProductRepository) OrderItemService {
	return &orderItemService{repo: r, productRepo: pr}
}

func (s *orderItemService) Create(it *models.OrderItem) error {
	p, err := s.productRepo.FindByID(it.ProductID)
	if err != nil {
		return err
	}
	it.Subtotal = float64(it.Quantity) * p.Price
	return s.repo.Create(it)
}

func (s *orderItemService) FindAll() ([]models.OrderItem, error) { return s.repo.FindAll() }

func (s *orderItemService) FindByID(id uint) (*models.OrderItem, error) { return s.repo.FindByID(id) }

func (s *orderItemService) Update(it *models.OrderItem) error {
	p, err := s.productRepo.FindByID(it.ProductID)
	if err != nil {
		return err
	}
	it.Subtotal = float64(it.Quantity) * p.Price
	return s.repo.Update(it)
}
