package service

import (
	"Hexagonal_go/app/internal/core/model"
	"Hexagonal_go/app/internal/core/ports"
)

type AmphurService struct {
	amphurRepo ports.AmphurRepository
}

func NewAmphurService() *AmphurService {
	amphurRepo := ports.NewAmphurRepository()
	return &AmphurService{amphurRepo: amphurRepo}
}

func (s *AmphurService) FetchAllAmphur() ([]model.Amphur, error) {
	return s.amphurRepo.FetchAllAmphur()
}

func (s *AmphurService) FetchAmphurByProvinceID(provinceID int64) ([]model.Amphur, error) {
	return s.amphurRepo.FetchAmphurByProvinceID(provinceID)
}
