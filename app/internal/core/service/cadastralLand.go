package service

import (
	"Hexagonal_go/app/internal/core/model"
	"Hexagonal_go/app/internal/core/ports"
	"database/sql"
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

func (s *CadastralLandService) UpdateCadastralLand(cadastralLand model.CadastralLand, id int64) (sql.Result, error) {
	return s.cadastralLand.UpdateCadastralLand(cadastralLand, id)
}

func (s *CadastralLandService) InsertCadastralLand(cadastralLand *model.CadastralLand) (sql.Result, error) {
	return s.cadastralLand.InsertCadastralLand(cadastralLand)
}
