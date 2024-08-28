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
	tenorRepo repository.TenorRepository
}

func NewTenorUseCase(repo repository.TenorRepository) TenorUseCase {
	return &tenorUseCase{tenorRepo: repo}
}

func (u *tenorUseCase) CreateTenor(tenor *entity.Tenor) error {
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

	return u.tenorRepo.CreateTenor(tenor)
}

func (u *tenorUseCase) GetTenorsByCustomerID(customerID string) ([]entity.Tenor, error) {
	return u.tenorRepo.GetTenorsByCustomerID(customerID)
}

func (u *tenorUseCase) UpdateIsLunas(tenorID uint) error {
	return u.tenorRepo.UpdateIsLunas(tenorID)
}
