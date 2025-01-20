package persistence

import (
	"Hexagonal_go/app/internal/core/model"
	"fmt"
	"path/filepath"
)

type GetSVACadastral01RepositoryImpl struct{}

func (r *GetSVACadastral01RepositoryImpl) Fetch01getSVACadastral(cadastralSeq int64) ([]model.Cadastral01, error) {
	if DB == nil {
		InitDB()
	}
	queryPath := filepath.Join("sql", "dataQuery", "01getSVACadastral.sql")
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}

	rows, err := DB.Query(query, cadastralSeq)
	if err != nil {
		return nil, fmt.Errorf("failed to query cadastral data: %v", err)
	}
	defer rows.Close()

	var cadastrals []model.Cadastral01
	for rows.Next() {
		var cadastral model.Cadastral01
		if err := rows.Scan(
			&cadastral.CADASTRAL_SEQ, &cadastral.CADASTRAL_NO, &cadastral.SHEETCODE,
			&cadastral.BOX_NO, &cadastral.NUMOFSURVEY_QTY, &cadastral.LANDOFFICE_SEQ,
			&cadastral.PRINTPLATE_TYPE_SEQ, &cadastral.PRIVATESURVEY_FLAG, &cadastral.TYPEOFSURVEY_SEQ,
			&cadastral.TYPEOFSURVEY_NAME_TH, &cadastral.TYPEOFSURVEY_ADD1_SEQ, &cadastral.TYPEOFSURVEY_ADD2_SEQ,
			&cadastral.TYPEOFSURVEY_ADD3_SEQ, &cadastral.CADASTRAL_SURVEY_NO, &cadastral.CADASTRAL_SURVEY_NO_,
			&cadastral.CADASTRAL_TAMBOL_SEQ, &cadastral.CADASTRAL_TAMBOL_NAME_TH, &cadastral.CADASTRAL_AMPHUR_SEQ,
			&cadastral.CADASTRAL_AMPHUR_NAME_TH, &cadastral.CADASTRAL_PROVINCE_SEQ, &cadastral.CADASTRAL_PROVINCE_NAME_TH,
			&cadastral.ZONE_LAND, &cadastral.SHEETTYPE_SEQ, &cadastral.SHEETTYPE_NAME_TH,
			&cadastral.CADASTRAL_UTMMAP, &cadastral.CADASTRAL_UTMMAP1, &cadastral.CADASTRAL_UTMMAP2,
			&cadastral.CADASTRAL_UTMMAP3, &cadastral.CADASTRAL_UTMMAP4, &cadastral.CADASTRAL_ORIGINMAP,
			&cadastral.CADASTRAL_ORIGINMAP1, &cadastral.CADASTRAL_ORIGINMAP2, &cadastral.CADASTRAL_ORIGINMAP3,
			&cadastral.CADASTRAL_LAND_NO, &cadastral.CADASTRAL_LAND_NO_, &cadastral.AIRPHOTOMAP_NAME,
			&cadastral.AIRPHOTOMAP, &cadastral.AIRPHOTOMAP1, &cadastral.AIRPHOTOMAP2,
			&cadastral.AIRPHOTOMAP3, &cadastral.SCALE_MAP_SEQ, &cadastral.SCALE_MAP,
			&cadastral.SCALE_RAWANG_SEQ, &cadastral.SCALE_RAWANG, &cadastral.OWNER_TITLE_SEQ,
			&cadastral.OWNER_TITLE_NAME, &cadastral.OWNER_NAME, &cadastral.OWNER_FNAME,
			&cadastral.OWNER_LNAME, &cadastral.BENCHMARK_SEQ, &cadastral.BENCHMARK_NAME_TH,
			&cadastral.BENCHMARK_QTY, &cadastral.BENCHMARK2_SEQ, &cadastral.BENCHMARK2_NAME_TH,
			&cadastral.BENCHMARK2_QTY, &cadastral.SURVEY_DTM, &cadastral.SURVEY_DTM_,
			&cadastral.SURVEYJOB_SEQ, &cadastral.TITLE_SEQ, &cadastral.TITLE_NAME,
			&cadastral.SURVEYOR_NAME, &cadastral.SURVEYOR_FNAME, &cadastral.SURVEYOR_LNAME,
			&cadastral.SURVEYOR_POSITION, &cadastral.SURVEYOR_POSITION_, &cadastral.SURVEYOR_LEVEL,
			&cadastral.OLD_RAI_NUM, &cadastral.OLD_NGAN_NUM, &cadastral.OLD_WA_NUM,
			&cadastral.OLD_SUBWA_NUM, &cadastral.CADASTRAL_LAND_QTY, &cadastral.CADASTRAL_IMAGE_QTY,
			&cadastral.CADASTRAL_OWNER_QTY, &cadastral.LOST_FLAG, &cadastral.CANCEL_FLAG,
			&cadastral.CANCEL_USER, &cadastral.CANCEL_CAUSE, &cadastral.CANCELJOB_FLAG,
			&cadastral.CADASTRAL_SEQ_, &cadastral.CADASTRAL_LOG_SEQ, &cadastral.PROCESS_SEQ_,
			&cadastral.STATUS_SEQ_, &cadastral.CADASTRAL_NOTE, &cadastral.RECORD_STATUS,
			&cadastral.CREATE_USER, &cadastral.CREATE_DTM, &cadastral.LAST_UPD_USER,
			&cadastral.LAST_UPD_DTM, &cadastral.LANDOFFICE_NAME_TH, &cadastral.LANDOFFICE_NAME_EN,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		cadastrals = append(cadastrals, cadastral)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return cadastrals, nil
}
