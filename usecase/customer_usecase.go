package usecase

import (
	"mf-loan/entity"
	"mf-loan/repository"
)

type CustomerUseCase interface {
	CreateCustomer(customer *entity.DetailCustomer) error
	GetCustomerByID(NIK string) (*entity.DetailCustomer, error)
}

type customerUseCase struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{customerRepo: repo}
}

func (u *customerUseCase) CreateCustomer(customer *entity.DetailCustomer) error {
	return u.customerRepo.CreateCustomer(customer)
}

func (u *customerUseCase) GetCustomerByID(NIK string) (*entity.DetailCustomer, error) {
	return u.customerRepo.GetCustomerByID(NIK)
}
