package service

import (
	"Hexagonal_go/app/internal/core/model"
	"Hexagonal_go/app/internal/core/ports"
)

type GetMap01Service struct {
	getMap01Repo ports.GetMap01Repository
}

func NewGetMap01Service() *GetMap01Service {
	getMap01Repo := ports.NewGetMap01Repository()
	return &GetMap01Service{getMap01Repo: getMap01Repo}
}

func (s *GetMap01Service) FetchMapLandGIS(ogrFID, landOfficeSeq *int64, utmMap1, utmMap2, utmMap3, utmMap4, utmScale, landNo *string) ([]model.MapLandGIS, error) {
	return s.getMap01Repo.FetchMapLandGIS(ogrFID, landOfficeSeq, utmMap1, utmMap2, utmMap3, utmMap4, utmScale, landNo)
}
