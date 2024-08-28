package repository

import (
	"mf-loan/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *entity.TransactionDetail) error
	GetTransactionsByCustomerID(customerID string) ([]entity.TransactionDetail, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateTransaction(transaction *entity.TransactionDetail) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) GetTransactionsByCustomerID(customerID string) ([]entity.TransactionDetail, error) {
	var transactions []entity.TransactionDetail
	err := r.db.Where("customer_id = ?", customerID).Find(&transactions).Error
	return transactions, err
}
