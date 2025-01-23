package persistence

import (
	"Hexagonal_go/app/internal/core/model"
	"fmt"
	"path/filepath"
)

type TypeOfSurveyImpl struct{}

func (t *TypeOfSurveyImpl) FetchAllTypeOfSurvey() ([]model.TypeOfSurvey, error) {
	if DB == nil {
		InitDB()
	}
	queryPath := filepath.Join("sql", "MAS", "TB_MAS_TYPEOFSURVEY", "select.sql")
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query cadastral data: %v", err)
	}
	defer rows.Close()
	var typeOfSurveys []model.TypeOfSurvey
	for rows.Next() {
		var typeOfSurvey model.TypeOfSurvey
		if err := rows.Scan(
			&typeOfSurvey.TYPEOFSURVEY_SEQ, &typeOfSurvey.TYPEOFSURVEY_SYS_, &typeOfSurvey.TYPEOFSURVEY_ID,
			&typeOfSurvey.TYPEOFSURVEY_ABBR, &typeOfSurvey.TYPEOFSURVEY_NAME_TH, &typeOfSurvey.TYPEOFSURVEY_SHORTNAME,
			&typeOfSurvey.MAINTYPEOFSURVEY_SEQ, &typeOfSurvey.GROUPTYPE_SEQ, &typeOfSurvey.REG_GROUPTYPE_SEQ,
			&typeOfSurvey.PERFORMANCEREPORT_SEQ, &typeOfSurvey.PRIVATE_FLAG, &typeOfSurvey.PROJECT_FLAG,
			&typeOfSurvey.USEOLDDOC_FLAG, &typeOfSurvey.CAL_BY_PARCEL, &typeOfSurvey.CAL_BY_DAY,
			&typeOfSurvey.CAL_BY_DOC, &typeOfSurvey.CAL_BY_FILE, &typeOfSurvey.FEETYPE_SEQ,
			&typeOfSurvey.ADD_FEETYPE_SEQ, &typeOfSurvey.ADDITION_FLAG, &typeOfSurvey.TYPEOFSURVEY_NOTE,
			&typeOfSurvey.RECORD_STATUS, &typeOfSurvey.CREATE_USER, &typeOfSurvey.CREATE_DTM,
			&typeOfSurvey.LAST_UPD_USER, &typeOfSurvey.LAST_UPD_DTM,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		typeOfSurveys = append(typeOfSurveys, typeOfSurvey)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return typeOfSurveys, nil
}
