package main

import (
	"Hexagonal_go/app/internal/adapters/http"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // กำหนดค่าเริ่มต้นหากไม่มีการตั้งค่าใน .env
	}

	// สร้าง Fiber app
	app := fiber.New()

	// สร้าง HTTP Handler
	httpHandler := http.NewHandler()
	app.Get("/users", httpHandler.GetUsers)

	// เริ่มเซิร์ฟเวอร์
	log.Printf("Starting server on port %s", port)
	app.Listen(":" + port)
}
