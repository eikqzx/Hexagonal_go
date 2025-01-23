package persistence

import (
	"Hexagonal_go/app/internal/core/model"
	"fmt"
	"path/filepath"
)

type ProvinceRepositoryImpl struct{}

func (r *ProvinceRepositoryImpl) FetchAllProvinces() ([]model.Province, error) {
	if DB == nil {
		InitDB()
	}
	queryPath := filepath.Join("sql", "MAS", "TB_MAS_PROVINCE", "select_.sql")
	// log.Println(queryPath)
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query land offices: %v", err)
	}
	defer rows.Close()
	var provinces []model.Province
	for rows.Next() {
		var province model.Province
		if err := rows.Scan(
			&province.PROVINCE_SEQ, &province.PROVINCE_ID, &province.PROVINCE_ABBR,
			&province.PROVINCE_NAME_TH, &province.PROVINCE_NAME_EN, &province.REGION_SEQ,
			&province.COUNTRY_SEQ, &province.PROVINCE_DOL_ID, &province.PROVINCE_EVD_ID,
			&province.PROVINCE_DPIS_ID, &province.RECORD_STATUS, &province.CREATE_USER,
			&province.CREATE_DTM, &province.LAST_UPD_USER, &province.LAST_UPD_DTM,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		provinces = append(provinces, province)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return provinces, nil
}
