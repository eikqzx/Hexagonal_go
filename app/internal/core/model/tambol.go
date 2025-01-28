package model

import "time"

type Tambol struct {
	TambolSeq    int        `json:"TAMBOL_SEQ"`
	TambolId     *string    `json:"TAMBOL_ID"`
	TambolAbbr   *string    `json:"TAMBOL_ABBR"`
	TambolNameTh *string    `json:"TAMBOL_NAME_TH"`
	TambolNameEn *string    `json:"TAMBOL_NAME_EN"`
	AmphurSeq    *int       `json:"AMPHUR_SEQ"`
	TambolNote   *string    `json:"TAMBOL_NOTE"`
	ZipCode      *string    `json:"ZIP_CODE"`
	TambolEvdId  *string    `json:"TAMBOL_EVD_ID"`
	TambolDolId  *string    `json:"TAMBOL_DOL_ID"`
	TambolDpisId *string    `json:"TAMBOL_DPIS_ID"`
	RecordStatus *string    `json:"RECORD_STATUS"`
	CreateUser   *string    `json:"CREATE_USER"`
	CreateDtm    *time.Time `json:"CREATE_DTM"`
	LastUpdUser  *string    `json:"LAST_UPD_USER"`
	LastUpdDtm   *time.Time `json:"LAST_UPD_DTM"`
}
