package service

import (
	"Hexagonal_go/app/internal/core/model"
	"Hexagonal_go/app/internal/core/ports"
)

type TypeOfSurveyService struct {
	typeOfSurveyRepo ports.TypeOfSurveyRepository
}

func NewTypeOfSurveyService() *TypeOfSurveyService {
	typeOfSurveyRepo := ports.NewTypeOfSurveyRepository()
	return &TypeOfSurveyService{typeOfSurveyRepo: typeOfSurveyRepo}
}

func (s *TypeOfSurveyService) FetchTypeOfSurvey() ([]model.TypeOfSurvey, error) {
	return s.typeOfSurveyRepo.FetchAllTypeOfSurvey()
}
