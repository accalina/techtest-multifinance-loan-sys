package http

import (
	"mf-loan/entity"
	"mf-loan/usecase"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	useCase usecase.TransactionUseCase
}

func NewTransactionHandler(app *fiber.App, useCase usecase.TransactionUseCase) {
	handler := &TransactionHandler{useCase: useCase}

	app.Post("/transactions", handler.CreateTransaction)
	app.Get("/customers/:customer_id/transactions", handler.GetTransactionsByCustomerID)
}

// @Description	Create a new Tenor.
// @Summary		Create a new Tenor
// @Tags		Trasaction
// @Accept		json
// @Produce		json
// @Param		Customer	body		entity.TransactionDetail	true	"Tenor attribute"
// @Success		200		{object}	entity.TransactionDetail
// @Router		/transactions [post]
func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	transaction := new(entity.TransactionDetail)
	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
	}

	if err := h.useCase.CreateTransaction(transaction); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(transaction)
}

// @Summary			Get Trasaction Detail
// @Description		Get Trasaction by NIK.
// @Tags			Trasaction
// @Accept			json
// @Produce			json
// @Param        	customer_id	path     string  false  "NIK Customer"
// @Success			200		{array}		[]entity.TransactionDetail
// @Router			/customers/{customer_id}/transactions [get]
func (h *TransactionHandler) GetTransactionsByCustomerID(c *fiber.Ctx) error {
	customerID := c.Params("customer_id")
	transactions, err := h.useCase.GetTransactionsByCustomerID(customerID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "no transactions found for this customer"})
	}

	return c.JSON(transactions)
}
