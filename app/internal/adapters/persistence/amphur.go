package persistence

import (
	"Hexagonal_go/app/internal/core/model"
	"fmt"
	"path/filepath"
)

type AmphurImpl struct{}

func (a *AmphurImpl) FetchAllAmphur() ([]model.Amphur, error) {
	var amphurList []model.Amphur
	// สร้าง query ดึงข้อมูลจาก database
	if DB == nil {
		InitDB()
	}
	queryPath := filepath.Join("sql", "MAS", "TB_MAS_AMPHUR", "select_.sql")
	// log.Println(queryPath)
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query Amphur: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var amphur model.Amphur
		if err := rows.Scan(
			&amphur.AmphurSeq, &amphur.AmphurID, &amphur.AmphurAbbr,
			&amphur.AmphurNameTH, &amphur.AmphurNameEN, &amphur.ProvinceSeq,
			&amphur.AmphurEvdID, &amphur.AmphurDolID, &amphur.AmphurDpisID,
			&amphur.RecordStatus, &amphur.CreateUser, &amphur.CreateDTM,
			&amphur.LastUpdUser, &amphur.LastUpdDTM,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		amphurList = append(amphurList, amphur)
	}
	return amphurList, nil
}

func (a *AmphurImpl) FetchAmphurByProvinceID(provinceID int64) ([]model.Amphur, error) {
	var amphurList []model.Amphur
	// สร้าง query ดึงข้อมูลจาก database
	if DB == nil {
		InitDB()
	}
	queryPath := filepath.Join("sql", "MAS", "TB_MAS_AMPHUR", "select_ByProvince.sql")
	// log.Println(queryPath)
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}

	rows, err := DB.Query(query, provinceID)
	if err != nil {
		return nil, fmt.Errorf("failed to query Amphur: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var amphur model.Amphur
		if err := rows.Scan(
			&amphur.AmphurSeq, &amphur.AmphurID, &amphur.AmphurAbbr,
			&amphur.AmphurNameTH, &amphur.AmphurNameEN, &amphur.ProvinceSeq,
			&amphur.AmphurEvdID, &amphur.AmphurDolID, &amphur.AmphurDpisID,
			&amphur.RecordStatus, &amphur.CreateUser, &amphur.CreateDTM,
			&amphur.LastUpdUser, &amphur.LastUpdDTM,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		amphurList = append(amphurList, amphur)
	}
	return amphurList, nil
}
