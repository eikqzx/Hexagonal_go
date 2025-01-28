package model

import "time"

// CadastralLand เป็นโมเดลสำหรับข้อมูล Cadastral Land
type CadastralLand struct {
	CadastralLandSeq         *int       `json:"CADASTRAL_LAND_SEQ"`
	CadastralSeq             *int       `json:"CADASTRAL_SEQ"`
	CadastralSeq_            *int       `json:"CADASTRAL_SEQ_"`
	LandOrder                *int       `json:"LAND_ORDER"`
	ZoneLand                 *string    `json:"ZONE_LAND"`
	LandofficeSeq            *int       `json:"LANDOFFICE_SEQ"`
	SheettypeSeq             *int       `json:"SHEETTYPE_SEQ"`
	SheettypeNameTh          *string    `json:"SHEETTYPE_NAME_TH"`
	CadastralLandUtmap       *string    `json:"CADASTRAL_LAND_UTMMAP"`
	CadastralLandUtmap1      *string    `json:"CADASTRAL_LAND_UTMMAP1"`
	CadastralLandUtmap2      *string    `json:"CADASTRAL_LAND_UTMMAP2"`
	CadastralLandUtmap3      *string    `json:"CADASTRAL_LAND_UTMMAP3"`
	CadastralLandUtmap4      *string    `json:"CADASTRAL_LAND_UTMMAP4"`
	CadastralLandUtmap1_     *string    `json:"CADASTRAL_LAND_UTMMAP1_"`
	CadastralLandUtmap2_     *string    `json:"CADASTRAL_LAND_UTMMAP2_"`
	CadastralLandUtmap3_     *string    `json:"CADASTRAL_LAND_UTMMAP3_"`
	CadastralLandUtmap4_     *string    `json:"CADASTRAL_LAND_UTMMAP4_"`
	CadastralLandOriginmap   *string    `json:"CADASTRAL_LAND_ORIGINMAP"`
	CadastralLandOriginmap1  *string    `json:"CADASTRAL_LAND_ORIGINMAP1"`
	CadastralLandOriginmap2  *string    `json:"CADASTRAL_LAND_ORIGINMAP2"`
	CadastralLandOriginmap3  *string    `json:"CADASTRAL_LAND_ORIGINMAP3"`
	CadastralLandOriginmap1_ *string    `json:"CADASTRAL_LAND_ORIGINMAP1_"`
	CadastralLandOriginmap2_ *string    `json:"CADASTRAL_LAND_ORIGINMAP2_"`
	CadastralLandOriginmap3_ *string    `json:"CADASTRAL_LAND_ORIGINMAP3_"`
	CadastralLandNo          *string    `json:"CADASTRAL_LAND_NO"`
	CadastralLandNo_         *string    `json:"CADASTRAL_LAND_NO_"`
	AirphotomapName          *string    `json:"AIRPHOTOMAP_NAME"`
	Airphotomap              *string    `json:"AIRPHOTOMAP"`
	Airphotomap1             *string    `json:"AIRPHOTOMAP1"`
	Airphotomap2             *string    `json:"AIRPHOTOMAP2"`
	Airphotomap3             *string    `json:"AIRPHOTOMAP3"`
	Airphotomap1_            *string    `json:"AIRPHOTOMAP1_"`
	Airphotomap2_            *string    `json:"AIRPHOTOMAP2_"`
	Airphotomap3_            *string    `json:"AIRPHOTOMAP3_"`
	ScaleSeq                 *int       `json:"SCALE_SEQ"`
	ScaleMap                 *string    `json:"SCALE_MAP"`
	ScaleMap_                *string    `json:"SCALE_MAP_"`
	CadastralLandRaiNum      *int       `json:"CADASTRAL_LAND_RAI_NUM"`
	CadastralLandNganNum     *int       `json:"CADASTRAL_LAND_NGAN_NUM"`
	CadastralLandWaNum       *int       `json:"CADASTRAL_LAND_WA_NUM"`
	CadastralLandSubwaNum    *int       `json:"CADASTRAL_LAND_SUBWA_NUM"`
	CadastralProvinceSeq     *int       `json:"CADASTRAL_PROVINCE_SEQ"`
	CadastralAmphurSeq       *int       `json:"CADASTRAL_AMPHUR_SEQ"`
	CadastralTambolSeq       *int       `json:"CADASTRAL_TAMBOL_SEQ"`
	CadastralSurveyNo        *string    `json:"CADASTRAL_SURVEY_NO"`
	CadastralSurveyNo_       *string    `json:"CADASTRAL_SURVEY_NO_"`
	StaticFlag               *string    `json:"STATIC_FLAG"`
	ParcelSeq                *int       `json:"PARCEL_SEQ"`
	CadastralLandLogSeq      *int       `json:"CADASTRAL_LAND_LOG_SEQ"`
	ProcessSeq_              *int       `json:"PROCESS_SEQ_"`
	StatusSeq_               *int       `json:"STATUS_SEQ_"`
	CadastralLandNote        *string    `json:"CADASTRAL_LAND_NOTE"`
	RecordStatus             *string    `json:"RECORD_STATUS"`
	CreateUser               *string    `json:"CREATE_USER"`
	CreateDtm                *time.Time `json:"CREATE_DTM"`
	LastUpdUser              *string    `json:"LAST_UPD_USER"`
	LastUpdDtm               *time.Time `json:"LAST_UPD_DTM"`
	LandofficeNameTh         *string    `json:"LANDOFFICE_NAME_TH"`
	ProvinceNameTh           *string    `json:"CADASTRAL_PROVINCE_NAME_TH"`
	AmphurNameTh             *string    `json:"CADASTRAL_AMPHUR_NAME_TH"`
	TambolNameTh             *string    `json:"CADASTRAL_TAMBOL_NAME_TH"`
}
