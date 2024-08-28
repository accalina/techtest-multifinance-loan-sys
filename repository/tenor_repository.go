package repository

import (
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
	return r.db.Create(tenor).Error
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
	return r.db.Model(&entity.Tenor{}).
		Where("id = ?", tenorID).
		Update("is_lunas", true).Error
}
