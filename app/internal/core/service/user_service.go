package service

import (
	"Hexagonal_go/app/internal/core/model"
	"Hexagonal_go/app/internal/core/ports"
)

// UserService บริการที่ใช้ในการจัดการกับผู้ใช้
type UserService struct {
	userRepo ports.UserRepository
}

// NewUserService สร้าง UserService ใหม่
func NewUserService() *UserService {
	userRepo := ports.NewUserRepository()
	return &UserService{userRepo: userRepo}
}

// GetAllUsers ดึงข้อมูลผู้ใช้ทั้งหมด
func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.userRepo.FetchAllUsers()
}
