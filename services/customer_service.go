package services

import "github.com/kryast/Crud-6.git/repositories"

type CustomerService interface {
}

type customerService struct {
	repo repositories.CustomerRepository
}

func NewCustomerService(r repositories.CustomerRepository) CustomerService {
	return &customerService{r}
}
