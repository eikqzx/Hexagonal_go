package persistence

import (
	"Hexagonal_go/app/internal/core/model"
	"fmt"
	"path/filepath"
)

type SurveyDocTypeImpl struct{}

func (r *SurveyDocTypeImpl) FetchSurveyDocType() ([]model.SurveyDocTypeGroup, error) {
	if DB == nil {
		InitDB()
	}

	queryPath := filepath.Join("sql", "MAS", "TB_MAS_SURVEYDOCTYPE", "select_group.sql")
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query cadastral data: %v", err)
	}
	defer rows.Close()

	var surveyDocTypes []model.SurveyDocTypeGroup
	for rows.Next() {
		var surveyDocType model.SurveyDocTypeGroup
		if err := rows.Scan(
			&surveyDocType.SURVEYDOCTYPE_GROUP, &surveyDocType.SURVEYDOCTYPE_SEQ, &surveyDocType.SURVEYDOCTYPE_SYS_,
			&surveyDocType.SURVEYDOCTYPE_ID, &surveyDocType.SURVEYDOCTYPE_NAME_TH, &surveyDocType.SURVEYDOCTYPE_NAME_EN,
			&surveyDocType.SURVEYDOCTYPE_ORDER, &surveyDocType.SURVEYDOCTYPE_NOTE, &surveyDocType.RECORD_STATUS,
			&surveyDocType.CREATE_USER, &surveyDocType.CREATE_DTM, &surveyDocType.LAST_UPD_USER,
			&surveyDocType.LAST_UPD_DTM,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		surveyDocTypes = append(surveyDocTypes, surveyDocType)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}
	return surveyDocTypes, nil
}
