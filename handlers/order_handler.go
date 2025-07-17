package handlers

import "github.com/kryast/Crud-6.git/services"

type OrderHandler struct {
	svc services.OrderService
}

func NewOrderHandler(s services.OrderService) *OrderHandler {
	return &OrderHandler{s}
}
