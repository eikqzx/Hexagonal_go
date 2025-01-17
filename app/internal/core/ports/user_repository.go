package ports

import (
	"Hexagonal_go/app/internal/adapters/persistence"
	"Hexagonal_go/app/internal/core/model"
)

// UserRepository Interface สำหรับดึงข้อมูลผู้ใช้
type UserRepository interface {
	FetchAllUsers() ([]model.User, error)
}

// NewUserRepository สร้าง UserRepository ใหม่
func NewUserRepository() UserRepository {
	// คืนค่า UserRepositoryImpl ที่ implement interface
	return &persistence.UserRepositoryImpl{}
}
