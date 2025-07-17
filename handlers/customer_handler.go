package handlers

import "github.com/kryast/Crud-6.git/services"

type CustomerHandler struct {
	svc services.CustomerService
}

func NewCustomerHandler(s services.CustomerService) *CustomerHandler {
	return &CustomerHandler{s}
}
