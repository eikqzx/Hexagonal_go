package service

import (
	"Hexagonal_go/app/internal/core/model"
	"Hexagonal_go/app/internal/core/ports"
	"fmt"
	"log"
)

type SummarySheetCodeService struct {
	summarySheetCodeRepo ports.GetSummarySheetCodeRepository
}

func NewSummarySheetCodeService() *SummarySheetCodeService {
	summarySheetCodeRepo := ports.NewGetSummarySheetCodeRepository()
	return &SummarySheetCodeService{summarySheetCodeRepo: summarySheetCodeRepo}
}

func (s *SummarySheetCodeService) FetchSummarySheetCode(cadastralSeq int64) ([]model.SummarySheetCode, error) {
	if s.summarySheetCodeRepo == nil {
		log.Println(s.summarySheetCodeRepo)
		return nil, fmt.Errorf("repository not initialized")
	}
	return s.summarySheetCodeRepo.FetchSummarySheetCode(cadastralSeq)
}
