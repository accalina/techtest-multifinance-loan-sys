package http_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"mf-loan/delivery/http"
	"mf-loan/entity"
	"mf-loan/repository/tests/mocks"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	app := fiber.New()
	mockUseCase := new(mocks.CustomerUseCase)
	http.NewCustomerHandler(app, mockUseCase)

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

	mockUseCase.On("CreateCustomer", customer).Return(nil)

	reqBody := map[string]interface{}{
		"nik":           "1234567890123456",
		"full_name":     "John Doe",
		"legal_name":    "Johnathan Doe",
		"tempat_lahir":  "Jakarta",
		"tanggal_lahir": "1990-01-01",
		"gaji":          10000000,
		"foto_ktp":      "ktp_image.jpg",
		"foto_selfie":   "selfie_image.jpg",
	}

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/customers", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)

	assert.Equal(t, 201, res.StatusCode)
	mockUseCase.AssertExpectations(t)
}

func TestGetCustomerByID(t *testing.T) {
	app := fiber.New()
	mockUseCase := new(mocks.CustomerUseCase)
	http.NewCustomerHandler(app, mockUseCase)

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

	mockUseCase.On("GetCustomerByID", customer.NIK).Return(customer, nil)

	req := httptest.NewRequest("GET", "/customers/"+customer.NIK, nil)
	res, _ := app.Test(req)

	assert.Equal(t, 200, res.StatusCode)
	mockUseCase.AssertExpectations(t)
}

func TestGetCustomerByID_NotFound(t *testing.T) {
	app := fiber.New()
	mockUseCase := new(mocks.CustomerUseCase)
	http.NewCustomerHandler(app, mockUseCase)

	mockUseCase.On("GetCustomerByID", "9999999999999999").Return(nil, errors.New("customer not found"))

	req := httptest.NewRequest("GET", "/customers/9999999999999999", nil)
	res, _ := app.Test(req)

	assert.Equal(t, 404, res.StatusCode)
	mockUseCase.AssertExpectations(t)
}
