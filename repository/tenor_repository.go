package repository

import (
	"errors"
	"mf-loan/entity"

	"gorm.io/gorm"
)

type TenorRepository interface {
	CreateTenor(tenor *entity.Tenor) error
	GetTenorsByCustomerID(customerID string) ([]entity.Tenor, error)
	CheckExistingTenor(customerID string, monthNumber int) (bool, error)
	UpdateIsLunas(tenorID uint) error
}

type tenorRepository struct {
	db *gorm.DB
}

func NewTenorRepository(db *gorm.DB) TenorRepository {
	return &tenorRepository{db: db}
}

func (r *tenorRepository) CreateTenor(tenor *entity.Tenor) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Validation: Check if there's an existing tenor with the same customer_id and month_number and isLunas is false
	var existingTenor entity.Tenor
	if err := tx.Where("id_customer = ? AND month_number = ? AND is_lunas = ?", tenor.CustomerID, tenor.MonthNumber, false).First(&existingTenor).Error; err == nil {
		tx.Rollback()
		return errors.New("a non-paid tenor with the same month number already exists for this customer")
	}

	if err := tx.Create(&tenor).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *tenorRepository) GetTenorsByCustomerID(customerID string) ([]entity.Tenor, error) {
	var tenors []entity.Tenor
	err := r.db.Where("customer_id = ?", customerID).Find(&tenors).Error
	return tenors, err
}

func (r *tenorRepository) CheckExistingTenor(customerID string, monthNumber int) (bool, error) {
	var count int64
	err := r.db.Model(&entity.Tenor{}).
		Where("customer_id = ? AND month_number = ? AND is_lunas = ?", customerID, monthNumber, false).
		Count(&count).Error
	return count > 0, err
}

func (r *tenorRepository) UpdateIsLunas(tenorID uint) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&entity.Tenor{}).Where("id = ?", tenorID).Update("is_lunas", true).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
