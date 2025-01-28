package main

import (
	"Hexagonal_go/app/internal/adapters/http"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} - ${latency}\n",
	}))
	// สร้าง HTTP Handler
	httpHandler := http.NewHandler()
	app.Get("/users", httpHandler.GetUsers)
	app.Get("/landoffices", httpHandler.GetAllLandOffice)
	app.Get("/landofficesSVA", httpHandler.GetSVALandOffice)
	app.Post("/getMasLandOfficeByProvince", httpHandler.GetLandOfficeByLandOfficeSeq)
	app.Post("/01getSVACadastral", httpHandler.GetFetch01getSVACadastral)
	app.Post("/02getSVACadastralImageNull", httpHandler.GetFetch02SVACadastralImageNull)
	app.Post("/02updateCadastralImageNull", httpHandler.Update02SVACadastralImageNull)
	app.Post("/02getCadastralImage", httpHandler.Get02CadastralImage)
	app.Post("/getSummarySheetCode", httpHandler.GetSummarySheetCode)
	app.Post("/getSummaryBoxBySheetCode", httpHandler.GetSummaryBoxBySheetCode)
	app.Post("/getCadastralKeyin", httpHandler.GetCadastralKeyin)
	app.Post("/getListAllKeyin", httpHandler.GetListAllKeyin)
	app.Get("/getSurveyDocTypeGroup", httpHandler.GetSurveyDocTypeGroup)
	app.Get("/typeOfSurvey", httpHandler.GetTypeOfSurvey)
	app.Get("/getProvince", httpHandler.GetAllProvince)
	app.Get("/getAmphur", httpHandler.GetAmphur)
	app.Get("/amphurByProvince/:id", httpHandler.GetAmphurByProvinceID)
	app.Get("/getTambol", httpHandler.GetTambol)
	app.Get("/tambolByAmphur/:id", httpHandler.GetTambolByAmphurID)
	app.Get("/cadastralLand", httpHandler.GetAllCadastralLand)
	app.Get("/cadastralLandBycadastralSeq/:id", httpHandler.GetCadastralLandByCadastralSeq)
	app.Post("/01getmap", httpHandler.GetMap01)
	// เริ่มเซิร์ฟเวอร์
	log.Printf("Starting server on port %s", port)
	app.Listen(":" + port)

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// <-quit
	// log.Println("Shutting down server...")
}
