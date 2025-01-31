package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Inisialisasi Fiber
	app := fiber.New()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Read .env is failed")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/DB_NAME?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	log.Println(dsn)

	// Jalankan server
	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
