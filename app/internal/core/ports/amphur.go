package ports

import (
	"Hexagonal_go/app/internal/adapters/persistence"
	"Hexagonal_go/app/internal/core/model"
)

type AmphurRepository interface {
	FetchAllAmphur() ([]model.Amphur, error) // ฟังก์ชันที่ต้องมีใน repository
	FetchAmphurByProvinceID(provinceID int64) ([]model.Amphur, error)
}

func NewAmphurRepository() AmphurRepository {
	// คืนค่า AmphurRepositoryImpl ที่ implement interface
	return &persistence.AmphurImpl{}
}
