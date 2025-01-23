package model

import "time"

// SurveyDocType เป็นโมเดลสำหรับข้อมูล Survey Document Type
type SurveyDocTypeGroup struct {
	SURVEYDOCTYPE_GROUP   string    `json:"SURVEYDOCTYPE_GROUP"`
	SURVEYDOCTYPE_SEQ     int64     `json:"SURVEYDOCTYPE_SEQ"`
	SURVEYDOCTYPE_SYS_    int64     `json:"SURVEYDOCTYPE_SYS_"`
	SURVEYDOCTYPE_ID      *string   `json:"SURVEYDOCTYPE_ID"`
	SURVEYDOCTYPE_NAME_TH string    `json:"SURVEYDOCTYPE_NAME_TH"`
	SURVEYDOCTYPE_NAME_EN *string   `json:"SURVEYDOCTYPE_NAME_EN"`
	SURVEYDOCTYPE_ORDER   int64     `json:"SURVEYDOCTYPE_ORDER"`
	SURVEYDOCTYPE_NOTE    *string   `json:"SURVEYDOCTYPE_NOTE"`
	RECORD_STATUS         string    `json:"RECORD_STATUS"`
	CREATE_USER           string    `json:"CREATE_USER"`
	CREATE_DTM            time.Time `json:"CREATE_DTM"`
	LAST_UPD_USER         string    `json:"LAST_UPD_USER"`
	LAST_UPD_DTM          time.Time `json:"LAST_UPD_DTM"`
}
