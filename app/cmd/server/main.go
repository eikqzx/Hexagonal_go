package main

import (
	"Hexagonal_go/app/internal/adapters/http"
	"log"
	"os"

	// "os/signal"
	// "syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // กำหนดค่าเริ่มต้นหากไม่มีการตั้งค่าใน .env
	}

	// สร้าง Fiber app
	app := fiber.New(fiber.Config{
		Prefork: true, // เปิดใช้งาน Prefork
	})

	// สร้าง HTTP Handler
	httpHandler := http.NewHandler()
	app.Get("/users", httpHandler.GetUsers)
	app.Get("/landoffices", httpHandler.GetAllLandOffice)
	app.Post("/01getSVACadastral", httpHandler.GetFetch01getSVACadastral)
	app.Post("/02getCadastralImageNull", httpHandler.GetFetch02SVACadastralImageNull)
	app.Post("/02updateCadastralImageNull", httpHandler.Update02SVACadastralImageNull)
	// เริ่มเซิร์ฟเวอร์
	log.Printf("Starting server on port %s", port)
	app.Listen(":" + port)

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// <-quit
	// log.Println("Shutting down server...")
}
