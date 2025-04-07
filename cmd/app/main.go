package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joaodematejr/imersao22/go-gateway/internal/repository"
	"github.com/joaodematejr/imersao22/go-gateway/internal/service"
	"github.com/joaodematejr/imersao22/go-gateway/internal/web/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the application
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gateway"),
		getEnv("DB_SSL_MODE", "disable"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	port := getEnv("PORT", "8080")
	fmt.Printf("Starting server on port %s...\n", port)
	srv := server.NewServer(accountService, port)
	srv.ConfigureRoutes()

	if err := srv.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)

	}

}
