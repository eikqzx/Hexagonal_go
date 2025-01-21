package model

import "time"

// LandOffice เป็นโมเดลสำหรับข้อมูล LandOffice
type LandOfficeSVA struct {
	LANDOFFICE_SEQ      int64     `json:"LANDOFFICE_SEQ"`
	LANDOFFICE_SYS_     *string   `json:"LANDOFFICE_SYS_"`
	LANDOFFICE_NO       *string   `json:"LANDOFFICE_NO"`
	LANDOFFICE_ID       string    `json:"LANDOFFICE_ID"`
	LANDOFFICE_ABBR     *string   `json:"LANDOFFICE_ABBR"`
	LANDOFFICE_NAME_TH  string    `json:"LANDOFFICE_NAME_TH"`
	LANDOFFICE_NAME_EN  string    `json:"LANDOFFICE_NAME_EN"`
	PROVINCE_SEQ        int64     `json:"PROVINCE_SEQ"`
	PC_SERVER           *string   `json:"PC_SERVER"`
	CADASTRAL_SEQ       *int64    `json:"CADASTRAL_SEQ"`
	CADASTRAL_IMAGE_SEQ *int64    `json:"CADASTRAL_IMAGE_SEQ"`
	IP_ADDRESS          *string   `json:"IP_ADDRESS"`
	LANDOFFICE_NOTE     *string   `json:"LANDOFFICE_NOTE"`
	RECORD_STATUS       string    `json:"RECORD_STATUS"`
	CREATE_USER         string    `json:"CREATE_USER"`
	CREATE_DTM          time.Time `json:"CREATE_DTM"`
	LAST_UPD_USER       string    `json:"LAST_UPD_USER"`
	LAST_UPD_DTM        time.Time `json:"LAST_UPD_DTM"`
}
