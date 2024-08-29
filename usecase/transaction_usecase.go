package usecase

import (
	"errors"
	"mf-loan/entity"
	"mf-loan/repository"
)

type TransactionUseCase interface {
	CreateTransaction(transaction *entity.TransactionDetail) error
	GetTransactionsByCustomerID(customerID string) ([]entity.TransactionDetail, error)
}

type transactionUseCase struct {
	transactionRepo repository.TransactionRepository
	customerRepo    repository.CustomerRepository
}

func NewTransactionUseCase(trRepo repository.TransactionRepository, customerRepo repository.CustomerRepository) TransactionUseCase {
	return &transactionUseCase{transactionRepo: trRepo, customerRepo: customerRepo}
}

func (u *transactionUseCase) CreateTransaction(transaction *entity.TransactionDetail) error {
	// Check if user is on our db
	customer, err := u.customerRepo.GetCustomerByID(transaction.CustomerID)
	if err != nil || customer == nil {
		return errors.New("customer not found")
	}
	return u.transactionRepo.CreateTransaction(transaction)
}

func (u *transactionUseCase) GetTransactionsByCustomerID(customerID string) ([]entity.TransactionDetail, error) {
	return u.transactionRepo.GetTransactionsByCustomerID(customerID)
}
