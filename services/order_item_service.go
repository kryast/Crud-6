package services

import "github.com/kryast/Crud-6.git/repositories"

type OrderItemService interface {
}

type orderItemService struct {
	repo        repositories.OrderItemRepository
	productRepo repositories.ProductRepository
}

func NewOrderItemService(r repositories.OrderItemRepository, pr repositories.ProductRepository) OrderItemService {
	return &orderItemService{repo: r, productRepo: pr}
}
