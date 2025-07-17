package services

import "github.com/kryast/Crud-6.git/repositories"

type PaymentService interface {
}

type paymentService struct {
	repo           repositories.PaymentRepository
	orderItemsRepo repositories.OrderItemRepository
}

func NewPaymentService(pr repositories.PaymentRepository, oir repositories.OrderItemRepository) PaymentService {
	return &paymentService{repo: pr, orderItemsRepo: oir}
}
