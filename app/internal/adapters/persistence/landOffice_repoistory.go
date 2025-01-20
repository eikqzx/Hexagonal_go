package persistence

import (
	"Hexagonal_go/app/internal/core/model"
	"fmt"
	"path/filepath"
)

type LandOfficeRepositoryImpl struct{}

// Implement method ที่อยู่ใน interface LandOfficeRepository
func (r *LandOfficeRepositoryImpl) FetchAllLandOffice() ([]model.LandOffice, error) {
	if DB == nil {
		InitDB()
	}
	queryPath := filepath.Join("sql", "MAS", "TB_MAS_LANDOFFICE", "select.sql")
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

	var landOffices []model.LandOffice
	for rows.Next() {
		var landOffice model.LandOffice
		if err := rows.Scan(
			&landOffice.LANDOFFICE_SEQ, &landOffice.LANDOFFICE_ID, &landOffice.LANDOFFICE_DOL_ID,
			&landOffice.LANDOFFICE_DPIS_ID, &landOffice.LANDOFFICE_NAME_TH, &landOffice.LANDOFFICE_NAME_EN,
			&landOffice.LANDOFFICE_NAME_RECEIPT, &landOffice.LANDOFFICE_SNAME, &landOffice.LANDOFFICE_UNIT_NO,
			&landOffice.LANDOFFICE_ADDR, &landOffice.LANDOFFICE_MOO, &landOffice.LANDOFFICE_BLD,
			&landOffice.LANDOFFICE_ROAD, &landOffice.LANDOFFICE_SOI, &landOffice.LANDOFFICE_TORK,
			&landOffice.LANDOFFICE_POSTCODE, &landOffice.LANDOFFICE_TEL, &landOffice.LANDOFFICE_FAX,
			&landOffice.LANDOFFICE_URL, &landOffice.LANDOFFICE_EMAIL, &landOffice.GFMIS_CAPITAL_ID,
			&landOffice.GFMIS_CAPITAL_PAY_ID, &landOffice.AREA_CODE, &landOffice.ORA_GEOMETRY,
			&landOffice.LANDOFFICE_USED, &landOffice.LANDOFFICE_TYPE_SEQ, &landOffice.OFFICE_INTERNAL_FLAG,
			&landOffice.REF_LANDOFFICE_SEQ, &landOffice.REGION_SEQ, &landOffice.OFFICE_TYPE_SEQ,
			&landOffice.PROVINCE_SEQ, &landOffice.AMPHUR_SEQ, &landOffice.TAMBOL_SEQ,
			&landOffice.RECORD_STATUS, &landOffice.CREATE_USER, &landOffice.CREATE_DTM,
			&landOffice.LAST_UPD_USER, &landOffice.LAST_UPD_DTM, &landOffice.LANDOFFICE_LOOKUP,
			&landOffice.LANDOFFICE_REPORT, &landOffice.LANDOFFICE_OHT_NAME, &landOffice.LANDOFFICE_PICT,
			&landOffice.LANDOFFICE_BUILD_DTM, &landOffice.LANDOFFICE_CANCEL_DTM,
			&landOffice.LANDOFFICE_LEVELUP_DTM, &landOffice.LANDOFFICE_SNAME_EN, &landOffice.LANDOFFICE_MOU,
			&landOffice.LANDOFFICE_REGULATION, &landOffice.ON_DC, &landOffice.COMPANY_CODE,
			&landOffice.TAX_ID, &landOffice.LANDOFFICE_NOTE, &landOffice.IPV4_ADDRESS,
			&landOffice.TB_TOR_CADASTRAL, &landOffice.TB_TOR_CADASTRAL_LAND, &landOffice.TB_TOR_CADASTRAL_OWNER,
			&landOffice.TB_TOR_CADASTRAL_IMAGE, &landOffice.TB_TOR_CADASTRAL_BORROW,
			&landOffice.TB_TOR_CIRACORE_IMAGE,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		landOffices = append(landOffices, landOffice)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return landOffices, nil
}
