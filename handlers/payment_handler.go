package handlers

import "github.com/kryast/Crud-6.git/services"

type PaymentHandler struct {
	svc services.PaymentService
}

func NewPaymentHandler(s services.PaymentService) *PaymentHandler {
	return &PaymentHandler{s}
}
