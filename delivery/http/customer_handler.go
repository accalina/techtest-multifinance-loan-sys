package http

import (
	"mf-loan/entity"
	"mf-loan/usecase"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	useCase usecase.CustomerUseCase
}

func NewCustomerHandler(app *fiber.App, useCase usecase.CustomerUseCase) {
	handler := &CustomerHandler{useCase: useCase}

	app.Post("/customers", handler.CreateCustomer)
	app.Get("/customers/:id", handler.GetCustomerByID)
}

func (h *CustomerHandler) GetCustomerByID(c *fiber.Ctx) error {
	id := c.Params("id")
	customer, err := h.useCase.GetCustomerByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "customer not found"})
	}
	return c.JSON(customer)
}

func (h *CustomerHandler) CreateCustomer(c *fiber.Ctx) error {
	var customerPayload = new(entity.CustomerPayload)
	if err := c.BodyParser(&customerPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
	}

	// Parse the date string to time.Time
	parsedDate, err := time.Parse("2006-01-02", customerPayload.TanggalLahir)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid date format"})
	}

	customer := entity.DetailCustomer{
		NIK:          customerPayload.NIK,
		FullName:     customerPayload.FullName,
		LegalName:    customerPayload.LegalName,
		TempatLahir:  customerPayload.TempatLahir,
		TanggalLahir: parsedDate,
		Gaji:         customerPayload.Gaji,
		FotoKTP:      customerPayload.FotoKTP,
		FotoSelfie:   customerPayload.FotoSelfie,
	}

	if err := h.useCase.CreateCustomer(&customer); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(customer)
}
