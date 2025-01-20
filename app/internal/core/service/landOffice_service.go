package service

import (
	"Hexagonal_go/app/internal/core/model"
	"Hexagonal_go/app/internal/core/ports"
)

// LandOfficeService บริการที่ใช้ในการจัดการกับสำนักงานที่ดิน
type LandOfficeService struct {
	landOfficeRepo ports.LandOfficeRepository
}

// NewLandOfficeService สร้าง LandOfficeService ใหม่
func NewLandOfficeService() *LandOfficeService {
	// ใช้ NewLandOfficeRepository แทน NewUserRepository
	landOfficeRepo := ports.NewLandOfficeRepository()
	return &LandOfficeService{landOfficeRepo: landOfficeRepo}
}

// GetAllLandOffice ดึงข้อมูลสำนักงานที่ดินทั้งหมด
func (s *LandOfficeService) GetAllLandOffice() ([]model.LandOffice, error) {
	return s.landOfficeRepo.FetchAllLandOffice()
}
