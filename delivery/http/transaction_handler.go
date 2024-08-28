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

func (h *TransactionHandler) GetTransactionsByCustomerID(c *fiber.Ctx) error {
	customerID := c.Params("customer_id")
	transactions, err := h.useCase.GetTransactionsByCustomerID(customerID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "no transactions found for this customer"})
	}

	return c.JSON(transactions)
}
