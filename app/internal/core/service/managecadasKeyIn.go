package service

import (
	"Hexagonal_go/app/internal/core/model"
	"Hexagonal_go/app/internal/core/ports"
	"fmt"
)

type SummarySheetCodeService struct {
	summarySheetCodeRepo ports.GetSummarySheetCodeRepository
}

func NewSummarySheetCodeService() *SummarySheetCodeService {
	summarySheetCodeRepo := ports.NewGetSummarySheetCodeRepository()
	return &SummarySheetCodeService{summarySheetCodeRepo: summarySheetCodeRepo}
}

func (s *SummarySheetCodeService) FetchSummarySheetCode(landofficeSeq int64) ([]model.SummarySheetCode, error) {
	if s.summarySheetCodeRepo == nil {
		// log.Println(s.summarySheetCodeRepo)
		return nil, fmt.Errorf("repository not initialized")
	}
	return s.summarySheetCodeRepo.FetchSummarySheetCode(landofficeSeq)
}

func (s *SummarySheetCodeService) FetchSummaryBoxBySheetCode(landofficeSeq int64, sheetCode int64) ([]model.SummaryBoxBySheetCode, error) {
	if s.summarySheetCodeRepo == nil {
		// log.Println(s.summarySheetCodeRepo)
		return nil, fmt.Errorf("repository not initialized")
	}
	return s.summarySheetCodeRepo.FetchSummaryBoxBySheetCode(landofficeSeq, sheetCode)
}

func (s *SummarySheetCodeService) FetchGetCadastralKeyin(landofficeSeq int64, sheetCode int64, boxNo string) ([]model.CadastralKeyin, error) {
	if s.summarySheetCodeRepo == nil {
		// log.Println(s.summarySheetCodeRepo)
		return nil, fmt.Errorf("repository not initialized")
	}
	return s.summarySheetCodeRepo.FetchGetCadastralKeyin(landofficeSeq, sheetCode, boxNo)
}

func (s *SummarySheetCodeService) FetchGetListAllKeyin(landofficeSeq int64, page int64, pageRow int64, startSheetcode *int64, endSheetcode *int64, startBox *int64, endBox *int64) ([]model.ListAllKeyin, error) {
	if s.summarySheetCodeRepo == nil {
		// log.Println(s.summarySheetCodeRepo)
		return nil, fmt.Errorf("repository not initialized")
	}
	return s.summarySheetCodeRepo.FetchListAllKeyin(landofficeSeq, page, pageRow, startSheetcode, endSheetcode, startBox, endBox)
}
