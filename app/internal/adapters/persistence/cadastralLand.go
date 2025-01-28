package persistence

import (
	"Hexagonal_go/app/internal/core/model"
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
