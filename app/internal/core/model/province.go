package model

import "time"

// Province เป็นโมเดลสำหรับข้อมูล Province
type Province struct {
	PROVINCE_SEQ     *int64     `json:"PROVINCE_SEQ"`
	PROVINCE_ID      *string    `json:"PROVINCE_ID"`
	PROVINCE_ABBR    *string    `json:"PROVINCE_ABBR"`
	PROVINCE_NAME_TH *string    `json:"PROVINCE_NAME_TH"`
	PROVINCE_NAME_EN *string    `json:"PROVINCE_NAME_EN"`
	REGION_SEQ       *int64     `json:"REGION_SEQ"`
	COUNTRY_SEQ      *int64     `json:"COUNTRY_SEQ"`
	PROVINCE_DOL_ID  *string    `json:"PROVINCE_DOL_ID"`
	PROVINCE_EVD_ID  *string    `json:"PROVINCE_EVD_ID"`
	PROVINCE_DPIS_ID *string    `json:"PROVINCE_DPIS_ID"`
	RECORD_STATUS    *string    `json:"RECORD_STATUS"`
	CREATE_USER      *string    `json:"CREATE_USER"`
	CREATE_DTM       *time.Time `json:"CREATE_DTM"`
	LAST_UPD_USER    *string    `json:"LAST_UPD_USER"`
	LAST_UPD_DTM     *time.Time `json:"LAST_UPD_DTM"`
}
