package ports

import (
	"Hexagonal_go/app/internal/adapters/persistence"
	"Hexagonal_go/app/internal/core/model"
	// "database/sql"
)

type TypeOfSurveyRepository interface {
	FetchAllTypeOfSurvey() ([]model.TypeOfSurvey, error) // ฟังก์ชันที่ต้องมีใน repository
}

func NewTypeOfSurveyRepository() TypeOfSurveyRepository {
	// คืนค่า TypeOfSurveyImpl ที่ implement interface
	return &persistence.TypeOfSurveyImpl{}
}
