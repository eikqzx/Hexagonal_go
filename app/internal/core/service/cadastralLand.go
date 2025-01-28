package service

import (
	"Hexagonal_go/app/internal/core/model"
	"Hexagonal_go/app/internal/core/ports"
)

type CadastralLandService struct {
	cadastralLand ports.CadastralLandRepository
}

func NewCadastralLandService() *CadastralLandService {
	cadastralLandRepo := ports.NewCadastralLandRepository()
	return &CadastralLandService{cadastralLand: cadastralLandRepo}
}

func (s *CadastralLandService) FetchAllCadastralLand() ([]model.CadastralLand, error) {
	return s.cadastralLand.FetchAllCadastralLand()
}

func (s *CadastralLandService) CadastralLandByCadastralSeq(cadastralSeq int64) ([]model.CadastralLand, error) {
	return s.cadastralLand.CadastralLandByCadastralSeq(cadastralSeq)
}
