package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/godror/godror"
	"github.com/joho/godotenv"
)

var DB *sql.DB

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

	// ตั้งค่า connection pool
	DB.SetMaxOpenConns(25)                 // จำนวนการเชื่อมต่อสูงสุดที่สามารถเปิดได้พร้อมกัน
	DB.SetMaxIdleConns(25)                 // จำนวนการเชื่อมต่อสูงสุดที่สามารถอยู่ในสถานะ idle
	DB.SetConnMaxLifetime(5 * time.Minute) // ระยะเวลาสูงสุดที่การเชื่อมต่อสามารถใช้งานได้

	// ตรวจสอบการเชื่อมต่อ
	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}

func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			log.Println("Database connection closed successfully.")
		}
	}
}
