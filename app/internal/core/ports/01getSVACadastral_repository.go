package ports

import (
	"Hexagonal_go/app/internal/adapters/persistence"
	"Hexagonal_go/app/internal/core/model"
)

type GetSVACadastral01Repository interface {
	Fetch01getSVACadastral(cadastralSeq int64) ([]model.Cadastral01, error) // ฟังก์ชันที่ต้องมีใน repository
}

func NewGetSVACadastral01Repository() GetSVACadastral01Repository {
	// คืนค่า LandOfficeRepositoryImpl ที่ implement interface
	return &persistence.GetSVACadastral01RepositoryImpl{}
}
