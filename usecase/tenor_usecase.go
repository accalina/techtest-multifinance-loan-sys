package usecase

import (
	"errors"
	"mf-loan/entity"
	"mf-loan/repository"
)

type TenorUseCase interface {
	CreateTenor(tenor *entity.Tenor) error
	GetTenorsByCustomerID(customerID string) ([]entity.Tenor, error)
	UpdateIsLunas(tenorID uint) error
}

type tenorUseCase struct {
	tenorRepo    repository.TenorRepository
	customerRepo repository.CustomerRepository
}

func NewTenorUseCase(tenorRepo repository.TenorRepository, customerRepo repository.CustomerRepository) TenorUseCase {
	return &tenorUseCase{tenorRepo: tenorRepo, customerRepo: customerRepo}
}

func (u *tenorUseCase) CreateTenor(tenor *entity.Tenor) error {
	// Check if user is on our db
	customer, err := u.customerRepo.GetCustomerByID(tenor.CustomerID)
	if err != nil || customer == nil {
		return errors.New("customer not found")
	}

	// Validate MonthNumber 1 = january, 12 = desember
	if tenor.MonthNumber < 1 || tenor.MonthNumber > 12 {
		return errors.New("month_number must be between 1 and 12")
	}

	// We will check if a tenor with the same customer_id and month_number exists and isLunas is false
	exists, err := u.tenorRepo.CheckExistingTenor(tenor.CustomerID, tenor.MonthNumber)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("a tenor with the same month_number and customer_id already exists and isLunas is false")
	}

	if err := u.tenorRepo.CreateTenor(tenor); err != nil {
		return errors.New("failed to create tenor: " + err.Error())
	}
	return nil
}

func (u *tenorUseCase) GetTenorsByCustomerID(customerID string) ([]entity.Tenor, error) {
	return u.tenorRepo.GetTenorsByCustomerID(customerID)
}

func (u *tenorUseCase) UpdateIsLunas(tenorID uint) error {
	return u.tenorRepo.UpdateIsLunas(tenorID)
}
