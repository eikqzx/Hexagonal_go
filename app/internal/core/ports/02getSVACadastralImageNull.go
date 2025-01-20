package ports

import (
	"Hexagonal_go/app/internal/adapters/persistence"
	"Hexagonal_go/app/internal/core/model"
	"database/sql"
)

type Get02SVACadastralImageNullRepo interface {
	Fetch02getSVACadastralImageNull(landofficeSeq int64) ([]model.CadastralImage, error) // ฟังก์ชันที่ต้องมีใน repository
	Update02SVACadastralImageNull(cadastralImage model.CadastralImage) (sql.Result, error)
}

func NewGetSVACadastralImageNullRepository() Get02SVACadastralImageNullRepo {
	return &persistence.GetSVACadastralImageNullImpl{}
}
