package handlers

import (
	"net/http"
	"strconv"

	"github.com/anoop-raw/Billing/services"

	"github.com/gofiber/fiber/v2"
)

type BillingHandler struct {
	service *services.BillingService
}

func NewBillingHandler(service *services.BillingService) *BillingHandler {
	return &BillingHandler{service: service}
}

func (h *BillingHandler) CreateLoan(c *fiber.Ctx) error {
	var request struct {
		Amount       float64 `json:"amount"`
		InterestRate float64 `json:"interest_rate"`
		Weeks        int     `json:"weeks"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	loan, err := h.service.CreateLoan(request.Amount, request.InterestRate, request.Weeks)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(loan)
}

func (h *BillingHandler) GetOutstanding(c *fiber.Ctx) error {
	loanID, _ := strconv.ParseUint(c.Params("loanID"), 10, 32)
	outstanding, err := h.service.GetOutstanding(uint(loanID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"outstanding": outstanding})
}

func (h *BillingHandler) GetLoan(c *fiber.Ctx) error {
	loanIDStr := c.Params("loanID")
	loanID, err := strconv.Atoi(loanIDStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid loan ID"})
	}

	loan, payments, err := h.service.GetLoan(uint(loanID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "could not fetch loan details"})
	}

	response := fiber.Map{
		"loan":     loan,
		"payments": payments,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func (h *BillingHandler) IsDelinquent(c *fiber.Ctx) error {
	loanID, _ := strconv.ParseUint(c.Params("loanID"), 10, 32)
	weekNumber, _ := strconv.ParseInt(c.Query("weekNumber"), 10, 32)
	isDelinquent, err := h.service.IsDelinquent(uint(loanID), weekNumber)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"isDelinquent": isDelinquent})
}

func (h *BillingHandler) MakePayment(c *fiber.Ctx) error {
	var request struct {
		Amount float64 `json:"amount"`
		Week   int     `json:"week"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	loanID, _ := strconv.ParseUint(c.Params("loanID"), 10, 32)
	if err := h.service.MakePayment(uint(loanID), request.Amount, request.Week); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "payment made successfully"})
}
