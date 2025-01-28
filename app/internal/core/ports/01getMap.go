package ports

import (
	"Hexagonal_go/app/internal/adapters/persistence"
	"Hexagonal_go/app/internal/core/model"
)

type GetMap01Repository interface {
	FetchMapLandGIS(ogrFID, landOfficeSeq *int64, utmMap1, utmMap2, utmMap3, utmMap4, utmScale, landNo *string) ([]model.MapLandGIS, error) // ฟังก์ชันที่ต้องมีใน repository
}

func NewGetMap01Repository() GetMap01Repository {
	// คืนค่า LandOfficeRepositoryImpl ที่ implement interface
	return &persistence.GetMapImpl{}
}
