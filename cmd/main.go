package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/anoop-raw/Billing/database"
	"github.com/anoop-raw/Billing/handlers"
	"github.com/anoop-raw/Billing/repo"
	service "github.com/anoop-raw/Billing/services"
	"github.com/gofiber/fiber/v2"
)

func waitForDatabase() error {
	maxAttempts := 5
	for i := 0; i < maxAttempts; i++ {
		time.Sleep(5 * time.Second)
		conn, err := net.DialTimeout("tcp", "db:5432", 1*time.Second)
		if err == nil {
			conn.Close()
			return nil
		}
		fmt.Printf("Attempt %d: Waiting for database to be ready\n", i+1)
	}
	return fmt.Errorf("database is not ready after %d attempts", maxAttempts)
}

func main() {
	if err := waitForDatabase(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	dbInstance := database.ConnectDb()

	// Initialize repository with the database instance
	repo := repo.NewSQLRepository(dbInstance.Db)

	// Initialize service with the repository
	accountService := service.NewBillingService(repo)

	// Initialize handler with the service
	accountHandler := handlers.NewBillingHandler(accountService)

	app := fiber.New()

	setupRoutes(app, accountHandler)

	app.Listen(":6000")
}
