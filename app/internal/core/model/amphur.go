package model

import "time"

type Amphur struct {
	AmphurSeq    *int       `json:"AMPHUR_SEQ"`
	AmphurID     *string    `json:"AMPHUR_ID"`
	AmphurAbbr   *string    `json:"AMPHUR_ABBR"`
	AmphurNameTH *string    `json:"AMPHUR_NAME_TH"`
	AmphurNameEN *string    `json:"AMPHUR_NAME_EN"`
	ProvinceSeq  *int       `json:"PROVINCE_SEQ"`
	AmphurEvdID  *string    `json:"AMPHUR_EVD_ID"`
	AmphurDolID  *string    `json:"AMPHUR_DOL_ID"`
	AmphurDpisID *string    `json:"AMPHUR_DPIS_ID"`
	RecordStatus *string    `json:"RECORD_STATUS"`
	CreateUser   *string    `json:"CREATE_USER"`
	CreateDTM    *time.Time `json:"CREATE_DTM"`
	LastUpdUser  *string    `json:"LAST_UPD_USER"`
	LastUpdDTM   *time.Time `json:"LAST_UPD_DTM"`
}
