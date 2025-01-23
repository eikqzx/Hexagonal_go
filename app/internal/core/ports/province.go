package ports

import (
	"Hexagonal_go/app/internal/adapters/persistence"
	"Hexagonal_go/app/internal/core/model"
)

type ProvinceRepository interface {
	FetchAllProvinces() ([]model.Province, error)
}

func NewProvinceRepository() ProvinceRepository {
	return &persistence.ProvinceRepositoryImpl{}
}
