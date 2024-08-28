package usecase

import (
	"mf-loan/entity"
	"mf-loan/repository"
)

type TransactionUseCase interface {
	CreateTransaction(transaction *entity.TransactionDetail) error
	GetTransactionsByCustomerID(customerID string) ([]entity.TransactionDetail, error)
}

type transactionUseCase struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionUseCase(repo repository.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{transactionRepo: repo}
}

func (u *transactionUseCase) CreateTransaction(transaction *entity.TransactionDetail) error {
	return u.transactionRepo.CreateTransaction(transaction)
}

func (u *transactionUseCase) GetTransactionsByCustomerID(customerID string) ([]entity.TransactionDetail, error) {
	return u.transactionRepo.GetTransactionsByCustomerID(customerID)
}
