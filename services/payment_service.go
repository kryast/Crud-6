package services

import (
	"time"

	"github.com/kryast/Crud-6.git/models"
	"github.com/kryast/Crud-6.git/repositories"
)

type PaymentService interface {
	Create(*models.Payment) error
	FindAll() ([]models.Payment, error)
	FindByID(id uint) (*models.Payment, error)
	Update(*models.Payment) error
}

type paymentService struct {
	repo           repositories.PaymentRepository
	orderItemsRepo repositories.OrderItemRepository
}

func NewPaymentService(pr repositories.PaymentRepository, oir repositories.OrderItemRepository) PaymentService {
	return &paymentService{repo: pr, orderItemsRepo: oir}
}

func (s *paymentService) Create(p *models.Payment) error {
	total, err := s.orderItemsRepo.SumSubtotalByOrderID(p.OrderID)
	if err != nil {
		return err
	}
	p.Amount = total
	if p.Status == "paid" {
		now := time.Now()
		p.PaidAt = &now
	}
	return s.repo.Create(p)
}

func (s *paymentService) FindAll() ([]models.Payment, error) {
	return s.repo.FindAll()
}

func (s *paymentService) FindByID(id uint) (*models.Payment, error) {
	return s.repo.FindByID(id)
}

func (s *paymentService) Update(p *models.Payment) error {
	total, err := s.orderItemsRepo.SumSubtotalByOrderID(p.OrderID)
	if err != nil {
		return err
	}
	p.Amount = total
	if p.Status == "paid" && p.PaidAt == nil {
		now := time.Now()
		p.PaidAt = &now
	}
	return s.repo.Update(p)
}
