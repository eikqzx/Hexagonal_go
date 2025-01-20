package service

import (
	"Hexagonal_go/app/internal/core/model"
	"Hexagonal_go/app/internal/core/ports"
	"fmt"
	"log"
)

type GetSVACadastral01Service struct {
	svaCadastralRepo ports.GetSVACadastral01Repository
}

func NewGetSVACadastral01Service() *GetSVACadastral01Service {
	// ใช้ NewLandOfficeRepository แทน NewUserRepository
	svaCadastralRepo := ports.NewGetSVACadastral01Repository()
	return &GetSVACadastral01Service{svaCadastralRepo: svaCadastralRepo}
}

// GetAllLandOffice ดึงข้อมูลสำนักงานที่ดินทั้งหมด
func (s *GetSVACadastral01Service) Fetch01getSVACadastral(cadastralSeq int64) ([]model.Cadastral01, error) {
	if s.svaCadastralRepo == nil {
		log.Println(s.svaCadastralRepo)
		return nil, fmt.Errorf("repository not initialized")
	}
	return s.svaCadastralRepo.Fetch01getSVACadastral(cadastralSeq)
}
