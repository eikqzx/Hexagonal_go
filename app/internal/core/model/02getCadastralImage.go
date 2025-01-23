package model

import "time"

// CadastralImage เป็นโมเดลสำหรับข้อมูล Cadastral Image
type CadastralImage02 struct {
	CADASTRAL_IMAGE_SEQ           *int64     `json:"CADASTRAL_IMAGE_SEQ"`
	LANDOFFICE_SEQ                *int64     `json:"LANDOFFICE_SEQ"`
	PRINTPLATE_TYPE_SEQ           *int64     `json:"PRINTPLATE_TYPE_SEQ"`
	CADASTRAL_SEQ                 *int64     `json:"CADASTRAL_SEQ"`
	CADASTRAL_IMAGE_SEQ_          *int64     `json:"CADASTRAL_IMAGE_SEQ_"`
	SURVEYDOCTYPE_SEQ             *int64     `json:"SURVEYDOCTYPE_SEQ"`
	SURVEYDOCTYPE_SEQ_            *int64     `json:"SURVEYDOCTYPE_SEQ_"`
	SURVEYDOCTYPE_ABBR            *string    `json:"SURVEYDOCTYPE_ABBR"`
	SURVEYDOCTYPE_ABBR_           *string    `json:"SURVEYDOCTYPE_ABBR_"`
	SURVEYDOCTYPE_NAME_TH         *string    `json:"SURVEYDOCTYPE_NAME_TH"`
	SURVEYDOCTYPE_GROUP           *string    `json:"SURVEYDOCTYPE_GROUP"`
	SURVEYDOCTYPE_GROUP_          *string    `json:"SURVEYDOCTYPE_GROUP_"`
	CADASTRAL_IMAGE_ORDER         *int64     `json:"CADASTRAL_IMAGE_ORDER"`
	CADASTRAL_IMAGE_PNO           *int64     `json:"CADASTRAL_IMAGE_PNO"`
	CADASTRAL_IMAGE_PNO_          *int64     `json:"CADASTRAL_IMAGE_PNO_"`
	CADASTRAL_IMAGE_THUMBNAIL_URL *string    `json:"CADASTRAL_IMAGE_THUMBNAIL_URL"`
	CADASTRAL_IMAGE_URL           *string    `json:"CADASTRAL_IMAGE_URL"`
	CADASTRAL_IMAGE_PATH          *string    `json:"CADASTRAL_IMAGE_PATH"`
	CADASTRAL_IMAGE_FILENAME      *string    `json:"CADASTRAL_IMAGE_FILENAME"`
	CADASTRAL_IMAGE_PNAME         *string    `json:"CADASTRAL_IMAGE_PNAME"`
	IMAGE_SIZE                    *int64     `json:"IMAGE_SIZE"`
	IMAGE_DPI_X                   *int64     `json:"IMAGE_DPI_X"`
	IMAGE_DPI_Y                   *int64     `json:"IMAGE_DPI_Y"`
	IMAGE_RES_X                   *int64     `json:"IMAGE_RES_X"`
	IMAGE_RES_Y                   *int64     `json:"IMAGE_RES_Y"`
	DETECT_X_                     *int64     `json:"DETECT_X_"`
	DETECT_Y_                     *int64     `json:"DETECT_Y_"`
	PROCESS_SEQ_                  *int64     `json:"PROCESS_SEQ_"`
	STATUS_SEQ_                   *int64     `json:"STATUS_SEQ_"`
	CADASTRAL_IMAGE_NOTE          *string    `json:"CADASTRAL_IMAGE_NOTE"`
	RECORD_STATUS                 *string    `json:"RECORD_STATUS"`
	CREATE_USER                   *string    `json:"CREATE_USER"`
	CREATE_DTM                    *time.Time `json:"CREATE_DTM"`
	LAST_UPD_USER                 *string    `json:"LAST_UPD_USER"`
	LAST_UPD_DTM                  *time.Time `json:"LAST_UPD_DTM"`
	SYNC_ORIGIN_SEQ               *int64     `json:"SYNC_ORIGIN_SEQ"`
	SYNC_DES_SEQ                  *int64     `json:"SYNC_DES_SEQ"`
	MIGRATE_CADAS_SEQ             *int64     `json:"MIGRATE_CADAS_SEQ"`
	IMAGE_PROCESS_SEQ_            *int64     `json:"IMAGE_PROCESS_SEQ_"`
	IMAGE_NOTE                    *string    `json:"IMAGE_NOTE"`
	IP_ADDRESS                    *string    `json:"IP_ADDRESS"`
	IMAGE_1_DTM                   *time.Time `json:"IMAGE_1_DTM"`
	IMAGE_2_DTM                   *time.Time `json:"IMAGE_2_DTM"`
	IMAGE_3_DTM                   *time.Time `json:"IMAGE_3_DTM"`
	IMAGE_SIZE_                   *int64     `json:"IMAGE_SIZE_"`
	IMAGE_TYPE                    *string    `json:"IMAGE_TYPE"`
	IMAGE_HORIZONTAL              *int64     `json:"IMAGE_HORIZONTAL"`
	IMAGE_VERTICAL                *int64     `json:"IMAGE_VERTICAL"`
	IMAGE_24BIT                   *int64     `json:"IMAGE_24BIT"`
	IMAGE_WIDTH                   *int64     `json:"IMAGE_WIDTH"`
	IMAGE_HEIGHT                  *int64     `json:"IMAGE_HEIGHT"`
	IMAGE_CROP                    *string    `json:"IMAGE_CROP"`
	IMAGE_LOST                    *int64     `json:"IMAGE_LOST"`
	IMAGE_CORRUPTED               *int64     `json:"IMAGE_CORRUPTED"`
	IMAGE_REVERSE                 *string    `json:"IMAGE_REVERSE"`
	IMAGE_SKEW                    *string    `json:"IMAGE_SKEW"`
	IMAGE_EDIT_COPY               *string    `json:"IMAGE_EDIT_COPY"`
	CIRACORE_IMAGE_NOTE           *string    `json:"CIRACORE_IMAGE_NOTE"`
	CIRACORE_IMAGE_STATUS         *string    `json:"CIRACORE_IMAGE_STATUS"`
	CIRACORE_IMAGE_DTM            *time.Time `json:"CIRACORE_IMAGE_DTM"`
	CIRACORE_DATA_NOTE            *string    `json:"CIRACORE_DATA_NOTE"`
	CIRACORE_DATA_STATUS          *int64     `json:"CIRACORE_DATA_STATUS"`
	CIRACORE_DATA_DTM             *time.Time `json:"CIRACORE_DATA_DTM"`
	OCR_DATA_NOTE                 *string    `json:"OCR_DATA_NOTE"`
	OCR_DATA_STATUS               *int64     `json:"OCR_DATA_STATUS"`
	OCR_DATA_DTM                  *time.Time `json:"OCR_DATA_DTM"`
}
