package main

import (
	"github.com/anoop-raw/Billing/handlers"
	"github.com/gofiber/fiber/v2"
)

// setting routes
func setupRoutes(app *fiber.App, handler *handlers.BillingHandler) {
	app.Get("/v1/loans/:loanID/outstanding", handler.GetOutstanding)
	app.Get("/v1/loans/:loanID/delinquent", handler.IsDelinquent)
	app.Post("/v1/loans/:loanID/payments", handler.MakePayment)
	app.Post("/v1/loans/create", handler.CreateLoan)
}
