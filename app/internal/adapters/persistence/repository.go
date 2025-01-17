package persistence

import (
	"Hexagonal_go/app/internal/core/model"
)

// UserRepositoryImpl เป็นการนำไปใช้จริงของ UserRepository
type UserRepositoryImpl struct{}

// FetchAllUsers เป็นการ implement method ของ UserRepository interface
func (r *UserRepositoryImpl) FetchAllUsers() ([]model.User, error) {
	// ตัวอย่างการคืนค่าผลลัพธ์ที่จำลองฐานข้อมูล
	return []model.User{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Doe"},
	}, nil
}
