package usecase_test

import (
	"errors"
	"mf-loan/entity"
	"mf-loan/repository/tests/mocks"
	"mf-loan/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	mockRepo := new(mocks.CustomerRepository) // This is the mock we generated
	uc := usecase.NewCustomerUseCase(mockRepo)

	customer := &entity.DetailCustomer{
		NIK:          "1234567890123456",
		FullName:     "John Doe",
		LegalName:    "Johnathan Doe",
		TempatLahir:  "Jakarta",
		TanggalLahir: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Gaji:         10000000,
		FotoKTP:      "ktp_image.jpg",
		FotoSelfie:   "selfie_image.jpg",
	}

	mockRepo.On("CreateCustomer", customer).Return(nil)

	err := uc.CreateCustomer(customer)
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetCustomerByID(t *testing.T) {
	mockRepo := new(mocks.CustomerRepository)
	uc := usecase.NewCustomerUseCase(mockRepo)

	customer := &entity.DetailCustomer{
		NIK:          "1234567890123456",
		FullName:     "John Doe",
		LegalName:    "Johnathan Doe",
		TempatLahir:  "Jakarta",
		TanggalLahir: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Gaji:         10000000,
		FotoKTP:      "ktp_image.jpg",
		FotoSelfie:   "selfie_image.jpg",
	}

	mockRepo.On("GetCustomerByID", customer.NIK).Return(customer, nil)

	dbCustomer, err := uc.GetCustomerByID(customer.NIK)
	assert.Nil(t, err)
	assert.Equal(t, customer.FullName, dbCustomer.FullName)
	mockRepo.AssertExpectations(t)
}

func TestGetCustomerByID_NotFound(t *testing.T) {
	mockRepo := new(mocks.CustomerRepository)
	uc := usecase.NewCustomerUseCase(mockRepo)

	mockRepo.On("GetCustomerByID", "9999999999999999").Return(nil, errors.New("customer not found"))

	dbCustomer, err := uc.GetCustomerByID("9999999999999999")
	assert.Nil(t, dbCustomer)
	assert.NotNil(t, err)
	assert.Equal(t, "customer not found", err.Error())
	mockRepo.AssertExpectations(t)
}
