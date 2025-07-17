package handlers

import "github.com/kryast/Crud-6.git/services"

type ProductHandler struct {
	svc services.ProductService
}

func NewProductHandler(s services.ProductService) *ProductHandler {
	return &ProductHandler{s}
}
