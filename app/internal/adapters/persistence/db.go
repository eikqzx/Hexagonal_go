package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/godror/godror" // Import godror driver
	"github.com/joho/godotenv"
)

var DB *sql.DB

// InitDB เชื่อมต่อกับ Oracle Database
func InitDB() {
	// โหลดไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// ดึงค่าการเชื่อมต่อจาก environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	if dbUser == "" || dbPassword == "" || dbConnectionString == "" {
		log.Fatal("Database environment variables are not properly set.")
	}

	// สร้าง Connection String
	dsn := fmt.Sprintf("%s/%s@%s", dbUser, dbPassword, dbConnectionString)

	// เชื่อมต่อฐานข้อมูล
	DB, err = sql.Open("godror", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to Oracle Database: %v", err)
	}

	// ทดสอบการเชื่อมต่อ
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping Oracle Database: %v", err)
	}
	log.Println("Connected to Oracle Database!")
}

// CloseDB ปิดการเชื่อมต่อกับฐานข้อมูล
func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed.")
	}
}
