package model

import "time"

// LandOffice เป็นโมเดลสำหรับข้อมูล LandOffice
type LandOffice struct {
	LANDOFFICE_SEQ          int64      `json:"LANDOFFICE_SEQ"`
	LANDOFFICE_ID           *string    `json:"LANDOFFICE_ID"`
	LANDOFFICE_DOL_ID       *string    `json:"LANDOFFICE_DOL_ID"`
	LANDOFFICE_DPIS_ID      *string    `json:"LANDOFFICE_DPIS_ID"`
	LANDOFFICE_NAME_TH      *string    `json:"LANDOFFICE_NAME_TH"`
	LANDOFFICE_NAME_EN      *string    `json:"LANDOFFICE_NAME_EN"`
	LANDOFFICE_NAME_RECEIPT *string    `json:"LANDOFFICE_NAME_RECEIPT"`
	LANDOFFICE_SNAME        *string    `json:"LANDOFFICE_SNAME"`
	LANDOFFICE_UNIT_NO      *string    `json:"LANDOFFICE_UNIT_NO"`
	LANDOFFICE_ADDR         *string    `json:"LANDOFFICE_ADDR"`
	LANDOFFICE_MOO          *string    `json:"LANDOFFICE_MOO"`
	LANDOFFICE_BLD          *string    `json:"LANDOFFICE_BLD"`
	LANDOFFICE_ROAD         *string    `json:"LANDOFFICE_ROAD"`
	LANDOFFICE_SOI          *string    `json:"LANDOFFICE_SOI"`
	LANDOFFICE_TORK         *string    `json:"LANDOFFICE_TORK"`
	LANDOFFICE_POSTCODE     *string    `json:"LANDOFFICE_POSTCODE"`
	LANDOFFICE_TEL          *string    `json:"LANDOFFICE_TEL"`
	LANDOFFICE_FAX          *string    `json:"LANDOFFICE_FAX"`
	LANDOFFICE_URL          *string    `json:"LANDOFFICE_URL"`
	LANDOFFICE_EMAIL        *string    `json:"LANDOFFICE_EMAIL"`
	GFMIS_CAPITAL_ID        *string    `json:"GFMIS_CAPITAL_ID"`
	GFMIS_CAPITAL_PAY_ID    *string    `json:"GFMIS_CAPITAL_PAY_ID"`
	AREA_CODE               *string    `json:"AREA_CODE"`
	ORA_GEOMETRY            *string    `json:"ORA_GEOMETRY"`
	LANDOFFICE_USED         *int64     `json:"LANDOFFICE_USED"`
	LANDOFFICE_TYPE_SEQ     *int64     `json:"LANDOFFICE_TYPE_SEQ"`
	OFFICE_INTERNAL_FLAG    *string    `json:"OFFICE_INTERNAL_FLAG"`
	REF_LANDOFFICE_SEQ      int64      `json:"REF_LANDOFFICE_SEQ"`
	REGION_SEQ              *int64     `json:"REGION_SEQ"`
	OFFICE_TYPE_SEQ         *int64     `json:"OFFICE_TYPE_SEQ"`
	PROVINCE_SEQ            *int64     `json:"PROVINCE_SEQ"`
	AMPHUR_SEQ              *int64     `json:"AMPHUR_SEQ"`
	TAMBOL_SEQ              *int64     `json:"TAMBOL_SEQ"`
	RECORD_STATUS           *string    `json:"RECORD_STATUS"`
	CREATE_USER             *string    `json:"CREATE_USER"`
	CREATE_DTM              *time.Time `json:"CREATE_DTM"`
	LAST_UPD_USER           *string    `json:"LAST_UPD_USER"`
	LAST_UPD_DTM            *time.Time `json:"LAST_UPD_DTM"`
	LANDOFFICE_LOOKUP       *string    `json:"LANDOFFICE_LOOKUP"`
	LANDOFFICE_REPORT       *string    `json:"LANDOFFICE_REPORT"`
	LANDOFFICE_OHT_NAME     *string    `json:"LANDOFFICE_OHT_NAME"`
	LANDOFFICE_PICT         *string    `json:"LANDOFFICE_PICT"`
	LANDOFFICE_BUILD_DTM    *time.Time `json:"LANDOFFICE_BUILD_DTM"`
	LANDOFFICE_CANCEL_DTM   *time.Time `json:"LANDOFFICE_CANCEL_DTM"`
	LANDOFFICE_LEVELUP_DTM  *time.Time `json:"LANDOFFICE_LEVELUP_DTM"`
	LANDOFFICE_SNAME_EN     *string    `json:"LANDOFFICE_SNAME_EN"`
	LANDOFFICE_MOU          *string    `json:"LANDOFFICE_MOU"`
	LANDOFFICE_REGULATION   *int64     `json:"LANDOFFICE_REGULATION"`
	ON_DC                   *int64     `json:"ON_DC"`
	COMPANY_CODE            *string    `json:"COMPANY_CODE"`
	TAX_ID                  *string    `json:"TAX_ID"`
	LANDOFFICE_NOTE         *string    `json:"LANDOFFICE_NOTE"`
	IPV4_ADDRESS            *string    `json:"IPV4_ADDRESS"`
	TB_TOR_CADASTRAL        *string    `json:"TB_TOR_CADASTRAL"`
	TB_TOR_CADASTRAL_LAND   *string    `json:"TB_TOR_CADASTRAL_LAND"`
	TB_TOR_CADASTRAL_OWNER  *string    `json:"TB_TOR_CADASTRAL_OWNER"`
	TB_TOR_CADASTRAL_IMAGE  *string    `json:"TB_TOR_CADASTRAL_IMAGE"`
	TB_TOR_CADASTRAL_BORROW *string    `json:"TB_TOR_CADASTRAL_BORROW"`
	TB_TOR_CIRACORE_IMAGE   *string    `json:"TB_TOR_CIRACORE_IMAGE"`
	PROVINCE_NAME_TH        *string    `json:"PROVINCE_NAME_TH"`
}
