package handlers

import "github.com/kryast/Crud-6.git/services"

type OrderItemHandler struct {
	svc services.OrderItemService
}

func NewOrderItemHandler(s services.OrderItemService) *OrderItemHandler {
	return &OrderItemHandler{s}
}
