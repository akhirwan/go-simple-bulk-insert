package main

import (
	"go-simple-bulk-insert/cmd"
)

func main() {
	cmd.Execute()
}

/* the part below is simple script */

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/schema"
// )

// // Struktur database
// type TestNumberCounter struct {
// 	Type        string `gorm:"primaryKey"`
// 	StartNumber int
// 	LastNumber  *int
// }

// type TestNumberTransaction struct {
// 	Number int `gorm:"primaryKey"`
// 	Action string
// }

// var db *gorm.DB

// func DBConns() {
// 	// Load .env file
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Read .env is failed")
// 	}

// 	// Buat connection string untuk MySQL
// 	dsn := fmt.Sprintf(
// 		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		os.Getenv("DB_USER"),
// 		os.Getenv("DB_PASSWORD"),
// 		os.Getenv("DB_HOST"),
// 		os.Getenv("DB_PORT"),
// 		os.Getenv("DB_NAME"),
// 	)

// 	// Koneksi ke database
// 	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("database connection is failed:", err)
// 	}

// 	log.Println("database connection is success")
// }

// // Struktur request
// type CreateRequest struct {
// 	Type   string `json:"type"`
// 	Total  int    `json:"total"`
// 	Action string `json:"action"`
// }

// // Handler untuk insert data
// func createHandler(c *fiber.Ctx) error {
// 	req := new(CreateRequest)
// 	if err := c.BodyParser(req); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "Invalid request payload"})
// 	}

// 	if req.Type == "" || req.Total <= 0 || req.Action == "" {
// 		return c.Status(400).JSON(fiber.Map{"error": "type, total & action are required and total must be a number"})
// 	}

// 	db.NamingStrategy = schema.NamingStrategy{
// 		SingularTable: true,
// 	}

// 	tx := db.Begin()

// 	var counter TestNumberCounter
// 	if err := tx.Where("type = ?", req.Type).First(&counter).Error; err != nil {
// 		tx.Rollback()
// 		return c.Status(400).JSON(fiber.Map{"error": "Invalid type"})
// 	}

// 	startNumber := counter.StartNumber
// 	lastNumber := counter.LastNumber
// 	currentNumber := startNumber
// 	if lastNumber != nil {
// 		currentNumber = *lastNumber + 1
// 	}

// 	for i := 0; i < req.Total; i++ {
// 		transaction := TestNumberTransaction{
// 			Number: currentNumber + i,
// 			Action: req.Action,
// 		}
// 		if err := tx.Create(&transaction).Error; err != nil {
// 			tx.Rollback()
// 			return c.Status(500).JSON(fiber.Map{"error": "Failed to insert transactions"})
// 		}
// 	}

// 	newLastNumber := currentNumber + req.Total - 1
// 	if err := tx.Model(&counter).Update("last_number", newLastNumber).Error; err != nil {
// 		tx.Rollback()
// 		return c.Status(500).JSON(fiber.Map{"error": "Failed to update counter"})
// 	}

// 	tx.Commit()

// 	return c.Status(200).JSON(fiber.Map{"message": "Data inserted"})
// }

// func main() {
// 	DBConns()
// 	// Fiber init
// 	app := fiber.New()

// 	app.Post("/create", createHandler)

// 	// Ambil port dari .env
// 	port := os.Getenv("APP_PORT")
// 	if port == "" {
// 		port = "3000" // Default port
// 	}

// 	log.Fatal(app.Listen(":" + port))
// }
