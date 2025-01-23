package service

import (
	"Hexagonal_go/app/internal/core/model"
	"Hexagonal_go/app/internal/core/ports"
)

type ProvinceService struct {
	provinceRepo ports.ProvinceRepository
}

func NewProvinceService() *ProvinceService {
	provinceRepo := ports.NewProvinceRepository()
	return &ProvinceService{provinceRepo: provinceRepo}
}

func (s *ProvinceService) FetchAllProvinces() ([]model.Province, error) {
	return s.provinceRepo.FetchAllProvinces()
}
