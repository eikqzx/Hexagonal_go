package ports

import (
	"Hexagonal_go/app/internal/adapters/persistence"
	"Hexagonal_go/app/internal/core/model"
	"database/sql"
)

type CadastralLandRepository interface {
	FetchAllCadastralLand() ([]model.CadastralLand, error)
	CadastralLandByCadastralSeq(cadastralSeq int64) ([]model.CadastralLand, error)
	UpdateCadastralLand(cadastralLand model.CadastralLand, id int64) (sql.Result, error)
	InsertCadastralLand(cadastralLand *model.CadastralLand) (sql.Result, error)
}

func NewCadastralLandRepository() CadastralLandRepository {
	return &persistence.CadastralLandImpl{}
}
