package model

import "time"

// TypeOfSurvey เป็นโมเดลสำหรับข้อมูล Type Of Survey
type TypeOfSurvey struct {
	TYPEOFSURVEY_SEQ       *int64     `json:"TYPEOFSURVEY_SEQ"`
	TYPEOFSURVEY_SYS_      *int64     `json:"TYPEOFSURVEY_SYS_"`
	TYPEOFSURVEY_ID        *string    `json:"TYPEOFSURVEY_ID"`
	TYPEOFSURVEY_ABBR      *string    `json:"TYPEOFSURVEY_ABBR"`
	TYPEOFSURVEY_NAME_TH   *string    `json:"TYPEOFSURVEY_NAME_TH"`
	TYPEOFSURVEY_SHORTNAME *string    `json:"TYPEOFSURVEY_SHORTNAME"`
	MAINTYPEOFSURVEY_SEQ   *int64     `json:"MAINTYPEOFSURVEY_SEQ"`
	GROUPTYPE_SEQ          *int64     `json:"GROUPTYPE_SEQ"`
	REG_GROUPTYPE_SEQ      *int64     `json:"REG_GROUPTYPE_SEQ"`
	PERFORMANCEREPORT_SEQ  *int64     `json:"PERFORMANCEREPORT_SEQ"`
	PRIVATE_FLAG           *string    `json:"PRIVATE_FLAG"`
	PROJECT_FLAG           *string    `json:"PROJECT_FLAG"`
	USEOLDDOC_FLAG         *string    `json:"USEOLDDOC_FLAG"`
	CAL_BY_PARCEL          *string    `json:"CAL_BY_PARCEL"`
	CAL_BY_DAY             *string    `json:"CAL_BY_DAY"`
	CAL_BY_DOC             *string    `json:"CAL_BY_DOC"`
	CAL_BY_FILE            *string    `json:"CAL_BY_FILE"`
	FEETYPE_SEQ            *int64     `json:"FEETYPE_SEQ"`
	ADD_FEETYPE_SEQ        *int64     `json:"ADD_FEETYPE_SEQ"`
	ADDITION_FLAG          *string    `json:"ADDITION_FLAG"`
	TYPEOFSURVEY_NOTE      *string    `json:"TYPEOFSURVEY_NOTE"`
	RECORD_STATUS          *string    `json:"RECORD_STATUS"`
	CREATE_USER            *string    `json:"CREATE_USER"`
	CREATE_DTM             *time.Time `json:"CREATE_DTM"`
	LAST_UPD_USER          *string    `json:"LAST_UPD_USER"`
	LAST_UPD_DTM           *time.Time `json:"LAST_UPD_DTM"`
}
