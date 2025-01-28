package ports

import (
	"Hexagonal_go/app/internal/adapters/persistence"
	"Hexagonal_go/app/internal/core/model"
)

type TambolRepository interface {
	FetchAllTambol() ([]model.Tambol, error) // ฟังก์ชันที่ต้องมีใน repository
	FetchTambolByAmphurID(amphurID int64) ([]model.Tambol, error)
}

func NewTambolRepository() TambolRepository {
	// คืนค่า TambolRepositoryImpl ที่ implement interface
	return &persistence.TambolImpl{}
}
