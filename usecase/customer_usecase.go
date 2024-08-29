package usecase

import (
	"errors"
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
	if err := u.customerRepo.CreateCustomer(customer); err != nil {
		return errors.New("failed to create customer: " + err.Error())
	}
	return nil
}

func (u *customerUseCase) GetCustomerByID(NIK string) (*entity.DetailCustomer, error) {
	return u.customerRepo.GetCustomerByID(NIK)
}
