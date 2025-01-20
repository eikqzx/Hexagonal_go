package persistence

import (
	"Hexagonal_go/app/internal/core/model"
	"database/sql"
	"fmt"
	"path/filepath"
)

type GetSVACadastralImageNullImpl struct{}

func (r *GetSVACadastralImageNullImpl) Fetch02getSVACadastralImageNull(landofficeSeq int64) ([]model.CadastralImage, error) {
	if DB == nil {
		InitDB()
	}
	queryPath := filepath.Join("sql", "dataQuery", "02getSVACadastralImageNull.sql")
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}

	rows, err := DB.Query(query, landofficeSeq)
	if err != nil {
		return nil, fmt.Errorf("failed to query cadastral data: %v", err)
	}
	defer rows.Close()

	var cadastralImages []model.CadastralImage
	for rows.Next() {
		var cadastralImage model.CadastralImage
		if err := rows.Scan(
			&cadastralImage.CadastralImageSeq, &cadastralImage.LandofficeSeq, &cadastralImage.PrintplateTypeSeq,
			&cadastralImage.CadastralSeq, &cadastralImage.CadastralImageSeq_, &cadastralImage.SurveyDocTypeSeq,
			&cadastralImage.SurveyDocTypeSeq_, &cadastralImage.SurveyDocTypeAbbr, &cadastralImage.SurveyDocTypeAbbr_,
			&cadastralImage.SurveyDocTypeNameTH, &cadastralImage.SurveyDocTypeGroup, &cadastralImage.SurveyDocTypeGroup_,
			&cadastralImage.CadastralImageOrder, &cadastralImage.CadastralImagePno, &cadastralImage.CadastralImagePno_,
			&cadastralImage.CadastralImageThumbnail, &cadastralImage.CadastralImageUrl, &cadastralImage.CadastralImagePath,
			&cadastralImage.CadastralImageFilename, &cadastralImage.CadastralImagePname, &cadastralImage.ImageSize,
			&cadastralImage.ImageDpiX, &cadastralImage.ImageDpiY, &cadastralImage.ImageResX,
			&cadastralImage.ImageResY, &cadastralImage.DetectX_, &cadastralImage.DetectY_,
			&cadastralImage.ProcessSeq_, &cadastralImage.StatusSeq_, &cadastralImage.CadastralImageNote,
			&cadastralImage.RecordStatus, &cadastralImage.CreateUser, &cadastralImage.CreateDtm,
			&cadastralImage.LastUpdUser, &cadastralImage.LastUpdDtm, &cadastralImage.SyncOriginSeq,
			&cadastralImage.SyncDesSeq, &cadastralImage.MigrateCadasSeq, &cadastralImage.ImageProcessSeq_,
			&cadastralImage.ImageNote, &cadastralImage.IpAddress, &cadastralImage.Image1Dtm,
			&cadastralImage.Image2Dtm, &cadastralImage.Image3Dtm, &cadastralImage.ImageSize_,
			&cadastralImage.ImageType, &cadastralImage.ImageHorizontal, &cadastralImage.ImageVertical,
			&cadastralImage.Image24bit, &cadastralImage.ImageWidth, &cadastralImage.ImageHeight,
			&cadastralImage.ImageCrop, &cadastralImage.ImageLost, &cadastralImage.ImageCorrupted,
			&cadastralImage.ImageReverse, &cadastralImage.ImageSkew, &cadastralImage.ImageEditCopy,
			&cadastralImage.CiracoreImageNote, &cadastralImage.CiracoreImageStatus, &cadastralImage.CiracoreImageDtm,
			&cadastralImage.CiracoreDataNote, &cadastralImage.CiracoreDataStatus, &cadastralImage.CiracoreDataDtm,
			&cadastralImage.OcrDataNote, &cadastralImage.OcrDataStatus, &cadastralImage.OcrDataDtm,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		cadastralImages = append(cadastralImages, cadastralImage)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return cadastralImages, nil
}

func (r *GetSVACadastralImageNullImpl) Update02SVACadastralImageNull(cadastralImage model.CadastralImage) (sql.Result, error) {
	if DB == nil {
		InitDB()
	}

	queryPath := filepath.Join("sql", "dataQuery", "02updateSVACadastralImageNull.sql")
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}
	result, err := DB.Exec(query,
		cadastralImage.SurveyDocTypeSeq, cadastralImage.SurveyDocTypeAbbr, cadastralImage.SurveyDocTypeAbbr_,
		cadastralImage.SurveyDocTypeNameTH, cadastralImage.SurveyDocTypeGroup, cadastralImage.LastUpdUser,
		cadastralImage.CadastralImageSeq,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update cadastral image: %v", err)
	}

	return result, nil
}
