package persistence

import (
	"Hexagonal_go/app/internal/core/model"
	"database/sql"
	"fmt"
	"path/filepath"
)

type SummarySheetCodeImpl struct{}

func (r *SummarySheetCodeImpl) FetchSummarySheetCode(landofficeSeq int64) ([]model.SummarySheetCode, error) {
	if DB == nil {
		InitDB()
	}

	queryPath := filepath.Join("sql", "managecadasKeyIn", "getSummarySheetCode.sql")
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}

	rows, err := DB.Query(query, landofficeSeq)
	if err != nil {
		return nil, fmt.Errorf("failed to query getSummarySheetCode data: %v", err)
	}
	defer rows.Close()

	var summarySheetCodes []model.SummarySheetCode
	for rows.Next() {
		var summarySheetCode model.SummarySheetCode
		if err := rows.Scan(
			&summarySheetCode.SHEETCODE, &summarySheetCode.UNIQUE_BOX_COUNT, &summarySheetCode.CADASTRAL_COUNT,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		// log.Println(summarySheetCode)
		summarySheetCodes = append(summarySheetCodes, summarySheetCode)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}
	return summarySheetCodes, nil
}

func (r *SummarySheetCodeImpl) FetchSummaryBoxBySheetCode(landofficeSeq int64, sheetcode int64) ([]model.SummaryBoxBySheetCode, error) {
	if DB == nil {
		InitDB()
	}

	queryPath := filepath.Join("sql", "managecadasKeyIn", "getSummaryBoxBySheetCode.sql")
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}

	rows, err := DB.Query(query, landofficeSeq, sheetcode)
	if err != nil {
		return nil, fmt.Errorf("failed to query FetchSummaryBoxBySheetCode data: %v", err)
	}
	defer rows.Close()
	var summaryBoxs []model.SummaryBoxBySheetCode
	for rows.Next() {
		var summaryBox model.SummaryBoxBySheetCode
		if err := rows.Scan(
			&summaryBox.BOX_NO, &summaryBox.CADASTRAL_COUNT,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		summaryBoxs = append(summaryBoxs, summaryBox)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return summaryBoxs, nil
}

func (r *SummarySheetCodeImpl) FetchGetCadastralKeyin(landofficeSeq int64, sheetcode int64, boxNo string) ([]model.CadastralKeyin, error) {
	if DB == nil {
		InitDB()
	}
	queryPath := filepath.Join("sql", "managecadasKeyIn", "getCadastralKeyin.sql")
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}

	rows, err := DB.Query(query, landofficeSeq, sheetcode, boxNo)
	if err != nil {
		return nil, fmt.Errorf("failed to query getCadastralKeyin data: %v", err)
	}
	defer rows.Close()
	var cadastralKeyin []model.CadastralKeyin
	for rows.Next() {
		var cadastralKey model.CadastralKeyin
		if err := rows.Scan(
			&cadastralKey.CADASTRAL_NO, &cadastralKey.NUMOFSURVEY_QTY, &cadastralKey.CADASTRAL_SEQ,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		cadastralKeyin = append(cadastralKeyin, cadastralKey)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return cadastralKeyin, nil
}

func (r *SummarySheetCodeImpl) FetchListAllKeyin(landofficeSeq int64, page int64, pageRow int64, startSheetcode *int64, endSheetcode *int64, startBox *int64, endBox *int64) ([]model.ListAllKeyin, error) {
	if DB == nil {
		InitDB()
	}
	queryPath := filepath.Join("sql", "managecadasKeyIn", "getListAllKeyin.sql")
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}
	// log.Println(landofficeSeq)
	rows, err := DB.Query(query,
		sql.Named("LANDOFFICE_SEQ", landofficeSeq),
		sql.Named("START_SHEETCODE", startSheetcode),
		sql.Named("END_SHEETCODE", endSheetcode),
		sql.Named("START_BOX", startBox),
		sql.Named("END_BOX", endBox),
		sql.Named("PAGE", page),
		sql.Named("PAGE_ROW", pageRow),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query getListAllKeyin data: %v", err)
	}
	defer rows.Close()
	var listAllKeyin []model.ListAllKeyin
	for rows.Next() {
		var listAll model.ListAllKeyin

		if err := rows.Scan(
			&listAll.BOX_NO, &listAll.CADASTRAL_NO, &listAll.CADASTRAL_SEQ,
			&listAll.NUMOFSURVEY_QTY, &listAll.ROWNUMBER, &listAll.SHEETCODE,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		listAllKeyin = append(listAllKeyin, listAll)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return listAllKeyin, nil
}
