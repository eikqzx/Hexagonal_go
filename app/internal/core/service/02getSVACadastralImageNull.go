package service

import (
	"Hexagonal_go/app/internal/core/model"
	"Hexagonal_go/app/internal/core/ports"
	"database/sql"
	"fmt"
	"log"
)

type Get02getSVACadastralImageNullService struct {
	svaCadastralImageNullRepo ports.Get02SVACadastralImageNullRepo
}

func NewGet02getSVACadastralImageNullService() *Get02getSVACadastralImageNullService {
	// ใช้ NewLandOfficeRepository แทน NewUserRepository
	svaCadastralImageNullRepo := ports.NewGetSVACadastralImageNullRepository()
	return &Get02getSVACadastralImageNullService{svaCadastralImageNullRepo: svaCadastralImageNullRepo}
}

// GetAllLandOffice ดึงข้อมูลสำนักงานที่ดินทั้งหมด
func (s *Get02getSVACadastralImageNullService) Fetch02getSVACadastralImageNullService(landofficeSeq int64) ([]model.CadastralImage, error) {
	if s.svaCadastralImageNullRepo == nil {
		log.Println(s.svaCadastralImageNullRepo)
		return nil, fmt.Errorf("repository not initialized")
	}
	return s.svaCadastralImageNullRepo.Fetch02getSVACadastralImageNull(landofficeSeq)
}

func (s *Get02getSVACadastralImageNullService) Update02SVACadastralImageNullService(cadastralImage model.CadastralImage) (sql.Result, error) {
	if s.svaCadastralImageNullRepo == nil {
		log.Println(s.svaCadastralImageNullRepo)
		return nil, fmt.Errorf("repository not initialized")
	}
	return s.svaCadastralImageNullRepo.Update02SVACadastralImageNull(cadastralImage)
}

func (s *Get02getSVACadastralImageNullService) Get02CadastralImageService(cadastralSeq int64) ([]model.CadastralImage, error) {
	if s.svaCadastralImageNullRepo == nil {
		log.Println(s.svaCadastralImageNullRepo)
		return nil, fmt.Errorf("repository not initialized")
	}
	return s.svaCadastralImageNullRepo.Get02CadastralImage(cadastralSeq)
}
