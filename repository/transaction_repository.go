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
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *transactionRepository) GetTransactionsByCustomerID(customerID string) ([]entity.TransactionDetail, error) {
	var transactions []entity.TransactionDetail
	err := r.db.Where("customer_id = ?", customerID).Find(&transactions).Error
	return transactions, err
}
