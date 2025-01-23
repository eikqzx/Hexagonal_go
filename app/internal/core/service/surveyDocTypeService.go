package service

import (
	"Hexagonal_go/app/internal/core/model"
	"Hexagonal_go/app/internal/core/ports"
)

type SurveyDocTypeService struct {
	surveyDocTypeRepo ports.SurveyDocTypeRepository
}

func NewSurveyDocTypeService() *SurveyDocTypeService {
	surveyDocTypeRepo := ports.NewSurveyDocTypeRepository()
	return &SurveyDocTypeService{surveyDocTypeRepo: surveyDocTypeRepo}
}

func (s *SurveyDocTypeService) FetchSurveyDocType() ([]model.SurveyDocTypeGroup, error) {
	return s.surveyDocTypeRepo.FetchSurveyDocType()
}
