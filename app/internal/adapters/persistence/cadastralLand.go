package persistence

import (
	"Hexagonal_go/app/internal/core/model"
	"database/sql"
	"fmt"
	"path/filepath"
)

type CadastralLandImpl struct{}

func (c *CadastralLandImpl) FetchAllCadastralLand() ([]model.CadastralLand, error) {
	if DB == nil {
		InitDB()
	}

	queryPath := filepath.Join("sql", "SVA_", "TB_SVA_CADASTRAL_LAND", "select_.sql")
	// log.Println(queryPath)
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query CADASTRAL_LAND: %v", err)
	}
	defer rows.Close()

	var cadastralLandList []model.CadastralLand

	for rows.Next() {
		var cadastralLand model.CadastralLand
		if err = rows.Scan(
			&cadastralLand.CadastralLandSeq, &cadastralLand.CadastralSeq, &cadastralLand.CadastralSeq_,
			&cadastralLand.LandOrder, &cadastralLand.ZoneLand, &cadastralLand.LandofficeSeq,
			&cadastralLand.SheettypeSeq, &cadastralLand.SheettypeNameTh, &cadastralLand.CadastralLandUtmap,
			&cadastralLand.CadastralLandUtmap1, &cadastralLand.CadastralLandUtmap2, &cadastralLand.CadastralLandUtmap3,
			&cadastralLand.CadastralLandUtmap4, &cadastralLand.CadastralLandUtmap1_, &cadastralLand.CadastralLandUtmap2_,
			&cadastralLand.CadastralLandUtmap3_, &cadastralLand.CadastralLandUtmap4_, &cadastralLand.CadastralLandOriginmap,
			&cadastralLand.CadastralLandOriginmap1, &cadastralLand.CadastralLandOriginmap2, &cadastralLand.CadastralLandOriginmap3,
			&cadastralLand.CadastralLandOriginmap1_, &cadastralLand.CadastralLandOriginmap2_, &cadastralLand.CadastralLandOriginmap3_,
			&cadastralLand.CadastralLandNo, &cadastralLand.CadastralLandNo_, &cadastralLand.AirphotomapName,
			&cadastralLand.Airphotomap, &cadastralLand.Airphotomap1, &cadastralLand.Airphotomap2,
			&cadastralLand.Airphotomap3, &cadastralLand.Airphotomap1_, &cadastralLand.Airphotomap2_,
			&cadastralLand.Airphotomap3_, &cadastralLand.ScaleSeq, &cadastralLand.ScaleMap,
			&cadastralLand.ScaleMap_, &cadastralLand.CadastralLandRaiNum, &cadastralLand.CadastralLandNganNum,
			&cadastralLand.CadastralLandWaNum, &cadastralLand.CadastralLandSubwaNum, &cadastralLand.CadastralProvinceSeq,
			&cadastralLand.CadastralAmphurSeq, &cadastralLand.CadastralTambolSeq, &cadastralLand.CadastralSurveyNo,
			&cadastralLand.CadastralSurveyNo_, &cadastralLand.StaticFlag, &cadastralLand.ParcelSeq,
			&cadastralLand.CadastralLandLogSeq, &cadastralLand.ProcessSeq_, &cadastralLand.StatusSeq_,
			&cadastralLand.CadastralLandNote, &cadastralLand.RecordStatus, &cadastralLand.CreateUser,
			&cadastralLand.CreateDtm, &cadastralLand.LastUpdUser, &cadastralLand.LastUpdDtm,
			&cadastralLand.ProvinceNameTh, &cadastralLand.AmphurNameTh, &cadastralLand.TambolNameTh,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		cadastralLandList = append(cadastralLandList, cadastralLand)
	}
	return cadastralLandList, nil
}

func (c *CadastralLandImpl) CadastralLandByCadastralSeq(cadastralSeq int64) ([]model.CadastralLand, error) {
	if DB == nil {
		InitDB()
	}

	queryPath := filepath.Join("sql", "SVA_", "TB_SVA_CADASTRAL_LAND", "select_ByCadastralSeq.sql")
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}

	rows, err := DB.Query(query, cadastralSeq)
	if err != nil {
		return nil, fmt.Errorf("failed to query CADASTRAL_LAND: %v", err)
	}
	defer rows.Close()

	var cadastralLandList []model.CadastralLand

	for rows.Next() {
		var cadastralLand model.CadastralLand
		if err = rows.Scan(
			&cadastralLand.CadastralLandSeq, &cadastralLand.CadastralSeq, &cadastralLand.CadastralSeq_,
			&cadastralLand.LandOrder, &cadastralLand.ZoneLand, &cadastralLand.LandofficeSeq,
			&cadastralLand.SheettypeSeq, &cadastralLand.SheettypeNameTh, &cadastralLand.CadastralLandUtmap,
			&cadastralLand.CadastralLandUtmap1, &cadastralLand.CadastralLandUtmap2, &cadastralLand.CadastralLandUtmap3,
			&cadastralLand.CadastralLandUtmap4, &cadastralLand.CadastralLandUtmap1_, &cadastralLand.CadastralLandUtmap2_,
			&cadastralLand.CadastralLandUtmap3_, &cadastralLand.CadastralLandUtmap4_, &cadastralLand.CadastralLandOriginmap,
			&cadastralLand.CadastralLandOriginmap1, &cadastralLand.CadastralLandOriginmap2, &cadastralLand.CadastralLandOriginmap3,
			&cadastralLand.CadastralLandOriginmap1_, &cadastralLand.CadastralLandOriginmap2_, &cadastralLand.CadastralLandOriginmap3_,
			&cadastralLand.CadastralLandNo, &cadastralLand.CadastralLandNo_, &cadastralLand.AirphotomapName,
			&cadastralLand.Airphotomap, &cadastralLand.Airphotomap1, &cadastralLand.Airphotomap2,
			&cadastralLand.Airphotomap3, &cadastralLand.Airphotomap1_, &cadastralLand.Airphotomap2_,
			&cadastralLand.Airphotomap3_, &cadastralLand.ScaleSeq, &cadastralLand.ScaleMap,
			&cadastralLand.ScaleMap_, &cadastralLand.CadastralLandRaiNum, &cadastralLand.CadastralLandNganNum,
			&cadastralLand.CadastralLandWaNum, &cadastralLand.CadastralLandSubwaNum, &cadastralLand.CadastralProvinceSeq,
			&cadastralLand.CadastralAmphurSeq, &cadastralLand.CadastralTambolSeq, &cadastralLand.CadastralSurveyNo,
			&cadastralLand.CadastralSurveyNo_, &cadastralLand.StaticFlag, &cadastralLand.ParcelSeq,
			&cadastralLand.CadastralLandLogSeq, &cadastralLand.ProcessSeq_, &cadastralLand.StatusSeq_,
			&cadastralLand.CadastralLandNote, &cadastralLand.RecordStatus, &cadastralLand.CreateUser,
			&cadastralLand.CreateDtm, &cadastralLand.LastUpdUser, &cadastralLand.LastUpdDtm,
			&cadastralLand.ProvinceNameTh, &cadastralLand.AmphurNameTh, &cadastralLand.TambolNameTh,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		cadastralLandList = append(cadastralLandList, cadastralLand)
	}
	return cadastralLandList, nil
}

func (c *CadastralLandImpl) UpdateCadastralLand(cadastralLand model.CadastralLand, arg int64) (sql.Result, error) {
	if DB == nil {
		InitDB()
	}

	queryPath := filepath.Join("sql", "SVA_", "TB_SVA_CADASTRAL_LAND", "update_.sql")
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}

	result, err := DB.Exec(query,
		query,
		sql.Named("CADASTRAL_SEQ", cadastralLand.CadastralSeq),
		sql.Named("CADASTRAL_SEQ_", cadastralLand.CadastralSeq_),
		sql.Named("LAND_ORDER", cadastralLand.LandOrder),
		sql.Named("ZONE_LAND", cadastralLand.ZoneLand),
		sql.Named("LANDOFFICE_SEQ", cadastralLand.LandofficeSeq),
		sql.Named("SHEETTYPE_SEQ", cadastralLand.SheettypeSeq),
		sql.Named("SHEETTYPE_NAME_TH", cadastralLand.SheettypeNameTh),
		sql.Named("CADASTRAL_LAND_UTMMAP", cadastralLand.CadastralLandUtmap),
		sql.Named("CADASTRAL_LAND_UTMMAP1", cadastralLand.CadastralLandUtmap1),
		sql.Named("CADASTRAL_LAND_UTMMAP2", cadastralLand.CadastralLandUtmap2),
		sql.Named("CADASTRAL_LAND_UTMMAP3", cadastralLand.CadastralLandUtmap3),
		sql.Named("CADASTRAL_LAND_UTMMAP4", cadastralLand.CadastralLandUtmap4),
		sql.Named("CADASTRAL_LAND_UTMMAP1_", cadastralLand.CadastralLandUtmap1_),
		sql.Named("CADASTRAL_LAND_UTMMAP2_", cadastralLand.CadastralLandUtmap2_),
		sql.Named("CADASTRAL_LAND_UTMMAP3_", cadastralLand.CadastralLandUtmap3_),
		sql.Named("CADASTRAL_LAND_UTMMAP4_", cadastralLand.CadastralLandUtmap4_),
		sql.Named("CADASTRAL_LAND_ORIGINMAP", cadastralLand.CadastralLandOriginmap),
		sql.Named("CADASTRAL_LAND_ORIGINMAP1", cadastralLand.CadastralLandOriginmap1),
		sql.Named("CADASTRAL_LAND_ORIGINMAP2", cadastralLand.CadastralLandOriginmap2),
		sql.Named("CADASTRAL_LAND_ORIGINMAP3", cadastralLand.CadastralLandOriginmap3),
		sql.Named("CADASTRAL_LAND_ORIGINMAP1_", cadastralLand.CadastralLandOriginmap1_),
		sql.Named("CADASTRAL_LAND_ORIGINMAP2_", cadastralLand.CadastralLandOriginmap2_),
		sql.Named("CADASTRAL_LAND_ORIGINMAP3_", cadastralLand.CadastralLandOriginmap3_),
		sql.Named("CADASTRAL_LAND_NO", cadastralLand.CadastralLandNo),
		sql.Named("CADASTRAL_LAND_NO_", cadastralLand.CadastralLandNo_),
		sql.Named("AIRPHOTOMAP_NAME", cadastralLand.AirphotomapName),
		sql.Named("AIRPHOTOMAP", cadastralLand.Airphotomap),
		sql.Named("AIRPHOTOMAP1", cadastralLand.Airphotomap1),
		sql.Named("AIRPHOTOMAP2", cadastralLand.Airphotomap2),
		sql.Named("AIRPHOTOMAP3", cadastralLand.Airphotomap3),
		sql.Named("AIRPHOTOMAP1_", cadastralLand.Airphotomap1_),
		sql.Named("AIRPHOTOMAP2_", cadastralLand.Airphotomap2_),
		sql.Named("AIRPHOTOMAP3_", cadastralLand.Airphotomap3_),
		sql.Named("SCALE_SEQ", cadastralLand.ScaleSeq),
		sql.Named("SCALE_MAP", cadastralLand.ScaleMap),
		sql.Named("SCALE_MAP_", cadastralLand.ScaleMap_),
		sql.Named("CADASTRAL_LAND_RAI_NUM", cadastralLand.CadastralLandRaiNum),
		sql.Named("CADASTRAL_LAND_NGAN_NUM", cadastralLand.CadastralLandNganNum),
		sql.Named("CADASTRAL_LAND_WA_NUM", cadastralLand.CadastralLandWaNum),
		sql.Named("CADASTRAL_LAND_SUBWA_NUM", cadastralLand.CadastralLandSubwaNum),
		sql.Named("CADASTRAL_PROVINCE_SEQ", cadastralLand.CadastralProvinceSeq),
		sql.Named("CADASTRAL_PROVINCE_NAME_TH", cadastralLand.ProvinceNameTh),
		sql.Named("CADASTRAL_AMPHUR_SEQ", cadastralLand.CadastralAmphurSeq),
		sql.Named("CADASTRAL_AMPHUR_NAME_TH", cadastralLand.AmphurNameTh),
		sql.Named("CADASTRAL_TAMBOL_SEQ", cadastralLand.CadastralTambolSeq),
		sql.Named("CADASTRAL_TAMBOL_NAME_TH", cadastralLand.TambolNameTh),
		sql.Named("CADASTRAL_SURVEY_NO", cadastralLand.CadastralSurveyNo),
		sql.Named("CADASTRAL_SURVEY_NO_", cadastralLand.CadastralSurveyNo_),
		sql.Named("STATIC_FLAG", cadastralLand.StaticFlag),
		sql.Named("PARCEL_SEQ", cadastralLand.ParcelSeq),
		sql.Named("CADASTRAL_LAND_LOG_SEQ", cadastralLand.CadastralLandLogSeq),
		sql.Named("PROCESS_SEQ_", cadastralLand.ProcessSeq_),
		sql.Named("STATUS_SEQ_", cadastralLand.StatusSeq_),
		sql.Named("CADASTRAL_LAND_NOTE", cadastralLand.CadastralLandNote),
		sql.Named("RECORD_STATUS", cadastralLand.RecordStatus),
		sql.Named("LAST_UPD_USER", cadastralLand.LastUpdUser),
		sql.Named("CADASTRAL_LAND_SEQ", cadastralLand.CadastralLandSeq),
		sql.Named("arg", arg),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update cadastral land: %v", err)
	}

	return result, nil
}

func (c *CadastralLandImpl) InsertCadastralLand(cadastralLand *model.CadastralLand) (sql.Result, error) {
	if DB == nil {
		InitDB()
	}

	queryPath := filepath.Join("sql", "SVA_", "TB_SVA_CADASTRAL_LAND", "insert_.sql")
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}

	result, err := DB.Exec(query,
		sql.Named("CADASTRAL_LAND_NGAN_NUM", cadastralLand.CadastralLandNganNum),
		sql.Named("CADASTRAL_LAND_RAI_NUM", cadastralLand.CadastralLandRaiNum),
		sql.Named("CADASTRAL_LAND_WA_NUM", cadastralLand.CadastralLandWaNum),
		sql.Named("CADASTRAL_LAND_SUBWA_NUM", cadastralLand.CadastralLandSubwaNum),
		sql.Named("CADASTRAL_LAND_UTMMAP1", cadastralLand.CadastralLandUtmap1),
		sql.Named("CADASTRAL_LAND_UTMMAP2", cadastralLand.CadastralLandUtmap2),
		sql.Named("CADASTRAL_LAND_UTMMAP3", cadastralLand.CadastralLandUtmap3),
		sql.Named("CADASTRAL_LAND_UTMMAP4", cadastralLand.CadastralLandUtmap4),
		sql.Named("CADASTRAL_PROVINCE_SEQ", cadastralLand.CadastralProvinceSeq),
		sql.Named("CADASTRAL_AMPHUR_SEQ", cadastralLand.CadastralAmphurSeq),
		sql.Named("CADASTRAL_TAMBOL_SEQ", cadastralLand.CadastralTambolSeq),
		sql.Named("SCALE_MAP", cadastralLand.ScaleMap),
		sql.Named("SCALE_SEQ", cadastralLand.ScaleSeq),
		sql.Named("CADASTRAL_SEQ", cadastralLand.CadastralSeq),
		sql.Named("LANDOFFICE_SEQ", cadastralLand.LandofficeSeq),
		sql.Named("CREATE_USER", cadastralLand.CreateUser),
		sql.Named("RECORD_STATUS", cadastralLand.RecordStatus),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert cadastral land: %v", err)
	}

	return result, nil
}
