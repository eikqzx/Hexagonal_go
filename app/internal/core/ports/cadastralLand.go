package ports

import (
	"Hexagonal_go/app/internal/adapters/persistence"
	"Hexagonal_go/app/internal/core/model"
)

type CadastralLandRepository interface {
	FetchAllCadastralLand() ([]model.CadastralLand, error)
	CadastralLandByCadastralSeq(cadastralSeq int64) ([]model.CadastralLand, error)
}

func NewCadastralLandRepository() CadastralLandRepository {
	return &persistence.CadastralLandImpl{}
}
