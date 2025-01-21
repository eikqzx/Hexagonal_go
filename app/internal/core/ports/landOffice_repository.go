package ports

import (
	"Hexagonal_go/app/internal/adapters/persistence"
	"Hexagonal_go/app/internal/core/model"
)

// LandOfficeRepository Interface สำหรับดึงข้อมูลผู้ใช้
type LandOfficeRepository interface {
	FetchAllLandOffice() ([]model.LandOffice, error) // ฟังก์ชันที่ต้องมีใน repository
	FetchSVALandOffice() ([]model.LandOfficeSVA, error)
	FetchLandOfficeByProvinceSeq(provinceSeq int64) ([]model.LandOffice, error)
}

func NewLandOfficeRepository() LandOfficeRepository {
	// คืนค่า LandOfficeRepositoryImpl ที่ implement interface
	return &persistence.LandOfficeRepositoryImpl{}
}
