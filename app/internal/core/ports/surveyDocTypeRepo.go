package ports

import (
	"Hexagonal_go/app/internal/adapters/persistence"
	"Hexagonal_go/app/internal/core/model"
)

type SurveyDocTypeRepository interface {
	FetchSurveyDocType() ([]model.SurveyDocTypeGroup, error)
}

func NewSurveyDocTypeRepository() SurveyDocTypeRepository {
	return &persistence.SurveyDocTypeImpl{}
}
