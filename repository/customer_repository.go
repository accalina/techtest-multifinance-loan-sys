package repository

import (
	"mf-loan/entity"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(customer *entity.DetailCustomer) error
	GetCustomerByID(NIK string) (*entity.DetailCustomer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) CreateCustomer(customer *entity.DetailCustomer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepository) GetCustomerByID(NIK string) (*entity.DetailCustomer, error) {
	var customer entity.DetailCustomer
	err := r.db.Where("nik = ?", NIK).First(&customer).Error
	return &customer, err
}
