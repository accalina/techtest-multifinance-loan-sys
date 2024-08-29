package repository_test

import (
	"mf-loan/entity"
	"mf-loan/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	// Use an in-memory SQLite database for testing
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&entity.DetailCustomer{})
	return db
}

func TestCreateCustomer(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewCustomerRepository(db)

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

	err := repo.CreateCustomer(customer)
	assert.Nil(t, err)

	var dbCustomer entity.DetailCustomer
	db.First(&dbCustomer, "nik = ?", customer.NIK)

	assert.Equal(t, customer.FullName, dbCustomer.FullName)
	assert.Equal(t, customer.Gaji, dbCustomer.Gaji)
}

func TestGetCustomerByID(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewCustomerRepository(db)

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

	repo.CreateCustomer(customer)

	dbCustomer, err := repo.GetCustomerByID(customer.NIK)
	assert.Nil(t, err)
	assert.Equal(t, customer.FullName, dbCustomer.FullName)
	assert.Equal(t, customer.LegalName, dbCustomer.LegalName)
}
