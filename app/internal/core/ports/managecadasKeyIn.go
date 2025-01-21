package ports

import (
	"Hexagonal_go/app/internal/adapters/persistence"
	"Hexagonal_go/app/internal/core/model"
)

type GetSummarySheetCodeRepository interface {
	FetchSummarySheetCode(cadastralSeq int64) ([]model.SummarySheetCode, error) // ฟังก์ชันที่ต้องมีใน repository
}

func NewGetSummarySheetCodeRepository() GetSummarySheetCodeRepository {
	// คืนค่า SummarySheetCodeImpl ที่ implement interface
	return &persistence.SummarySheetCodeImpl{}
}
