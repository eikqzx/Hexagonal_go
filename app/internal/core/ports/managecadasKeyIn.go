package ports

import (
	"Hexagonal_go/app/internal/adapters/persistence"
	"Hexagonal_go/app/internal/core/model"
)

type GetSummarySheetCodeRepository interface {
	FetchSummarySheetCode(landofficeSeq int64) ([]model.SummarySheetCode, error) // ฟังก์ชันที่ต้องมีใน repository
	FetchSummaryBoxBySheetCode(landofficeSeq int64, sheetCode int64) ([]model.SummaryBoxBySheetCode, error)
	FetchGetCadastralKeyin(landofficeSeq int64, sheetCode int64, boxNo string) ([]model.CadastralKeyin, error)
	FetchListAllKeyin(landofficeSeq int64, page int64, pageRow int64, startSheetcode *int64, endSheetcode *int64, startBox *int64, endBox *int64) ([]model.ListAllKeyin, error)
}

func NewGetSummarySheetCodeRepository() GetSummarySheetCodeRepository {
	// คืนค่า SummarySheetCodeImpl ที่ implement interface
	return &persistence.SummarySheetCodeImpl{}
}
