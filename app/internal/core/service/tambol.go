package service

import (
	"Hexagonal_go/app/internal/core/model"
	"Hexagonal_go/app/internal/core/ports"
)

type TambolService struct {
	tambolRepo ports.TambolRepository
}

func NewTambolService() *TambolService {
	tambolRepo := ports.NewTambolRepository()
	return &TambolService{tambolRepo: tambolRepo}
}

func (s *TambolService) FetchAllTambol() ([]model.Tambol, error) {
	return s.tambolRepo.FetchAllTambol()
}

func (s *TambolService) FetchTambolByAmphurID(amphurID int64) ([]model.Tambol, error) {
	return s.tambolRepo.FetchTambolByAmphurID(amphurID)
}
