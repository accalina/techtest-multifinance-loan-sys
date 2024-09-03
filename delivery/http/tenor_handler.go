package http

import (
	"mf-loan/entity"
	"mf-loan/usecase"

	"github.com/gofiber/fiber/v2"
)

type TenorHandler struct {
	useCase usecase.TenorUseCase
}

func NewTenorHandler(app *fiber.App, useCase usecase.TenorUseCase) {
	handler := &TenorHandler{useCase: useCase}

	app.Post("/tenors", handler.CreateTenor)
	app.Get("/customers/:customer_id/tenors", handler.GetTenorsByCustomerID)
	app.Patch("/tenors/:id/lunas", handler.UpdateIsLunas)
}

// @Description	Create a new Tenor.
// @Summary		Create a new Tenor
// @Tags		Tenor
// @Accept		json
// @Produce		json
// @Param		Customer	body		entity.Tenor	true	"Tenor attribute"
// @Success		200		{object}	entity.Tenor
// @Router		/tenors [post]
func (h *TenorHandler) CreateTenor(c *fiber.Ctx) error {
	tenor := new(entity.Tenor)
	if err := c.BodyParser(tenor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
	}

	if err := h.useCase.CreateTenor(tenor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(tenor)
}

// @Summary			Get Customer Detail
// @Description		Get Customer by NIK.
// @Tags			Customer
// @Accept			json
// @Produce			json
// @Param        	customer_id	path     string  false  "NIK Customer"
// @Success			200		{array}		entity.DetailCustomer
// @Router			/customers/{customer_id}/tenors [get]
func (h *TenorHandler) GetTenorsByCustomerID(c *fiber.Ctx) error {
	customerID := c.Params("customer_id")
	tenors, err := h.useCase.GetTenorsByCustomerID(customerID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "no tenors found for this customer"})
	}

	return c.JSON(tenors)
}

// @Description	Update Tenor to Lunas.
// @Summary		Update Tenor to Lunas
// @Tags		Tenor
// @Accept		json
// @Produce		json
// @Param       id	path     string  false  "NIK Customer"
// @Router		/tenors/{id}/lunas [patch]
func (h *TenorHandler) UpdateIsLunas(c *fiber.Ctx) error {
	tenorID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid tenor id"})
	}

	if err := h.useCase.UpdateIsLunas(uint(tenorID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}
