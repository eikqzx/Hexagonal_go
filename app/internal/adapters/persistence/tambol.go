package persistence

import (
	"Hexagonal_go/app/internal/core/model"
	"fmt"
	"path/filepath"
)

type TambolImpl struct{}

func (t *TambolImpl) FetchAllTambol() ([]model.Tambol, error) {
	var tambolList []model.Tambol
	// สร้าง query ดึงข้อมูลจาก database
	if DB == nil {
		InitDB()
	}

	queryPath := filepath.Join("sql", "MAS", "TB_MAS_TAMBOL", "select_.sql")
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query Tambol: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tambol model.Tambol
		if err := rows.Scan(
			&tambol.TambolSeq, &tambol.TambolId, &tambol.TambolAbbr,
			&tambol.TambolNameTh, &tambol.TambolNameEn, &tambol.AmphurSeq,
			&tambol.TambolNote, &tambol.ZipCode, &tambol.TambolEvdId,
			&tambol.TambolDolId, &tambol.TambolDpisId, &tambol.RecordStatus,
			&tambol.CreateUser, &tambol.CreateDtm, &tambol.LastUpdUser,
			&tambol.LastUpdDtm,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		tambolList = append(tambolList, tambol)
	}
	return tambolList, nil
}

func (t *TambolImpl) FetchTambolByAmphurID(amphurID int64) ([]model.Tambol, error) {
	var tambolList []model.Tambol
	// สร้าง query ดึงข้อมูลจาก database
	if DB == nil {
		InitDB()
	}
	queryPath := filepath.Join("sql", "MAS", "TB_MAS_TAMBOL", "select_ByAmphur.sql")
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}

	rows, err := DB.Query(query, amphurID)
	if err != nil {
		return nil, fmt.Errorf("failed to query Tambol: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tambol model.Tambol
		if err := rows.Scan(
			&tambol.TambolSeq, &tambol.TambolId, &tambol.TambolAbbr,
			&tambol.TambolNameTh, &tambol.TambolNameEn, &tambol.AmphurSeq,
			&tambol.TambolNote, &tambol.ZipCode, &tambol.TambolEvdId,
			&tambol.TambolDolId, &tambol.TambolDpisId, &tambol.RecordStatus,
			&tambol.CreateUser, &tambol.CreateDtm, &tambol.LastUpdUser,
			&tambol.LastUpdDtm,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		tambolList = append(tambolList, tambol)
	}
	return tambolList, nil
}
