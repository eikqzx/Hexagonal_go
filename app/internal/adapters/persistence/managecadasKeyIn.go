package persistence

import (
	"Hexagonal_go/app/internal/core/model"
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
		return nil, fmt.Errorf("failed to query cadastral data: %v", err)
	}
	defer rows.Close()

	var summarySheetCodes []model.SummarySheetCode
	for rows.Next() {
		var summarySheetCode model.SummarySheetCode
		if err := rows.Scan(
			&summarySheetCode.CadastralCount, &summarySheetCode.SheetCode, &summarySheetCode.UniqueBoxCount,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		summarySheetCodes = append(summarySheetCodes, summarySheetCode)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}
	return summarySheetCodes, nil
}
