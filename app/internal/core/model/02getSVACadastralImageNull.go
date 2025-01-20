package model

import "time"

// CadastralImage เป็นโมเดลสำหรับข้อมูล Cadastral Image
type CadastralImage struct {
	CadastralImageSeq       *int64     `json:"CADASTRAL_IMAGE_SEQ"`
	LandofficeSeq           *int64     `json:"LANDOFFICE_SEQ"`
	PrintplateTypeSeq       *int64     `json:"PRINTPLATE_TYPE_SEQ"`
	CadastralSeq            *int64     `json:"CADASTRAL_SEQ"`
	CadastralImageSeq_      *int64     `json:"CADASTRAL_IMAGE_SEQ_"`
	SurveyDocTypeSeq        *int64     `json:"SURVEYDOCTYPE_SEQ"`
	SurveyDocTypeSeq_       *int64     `json:"SURVEYDOCTYPE_SEQ_"`
	SurveyDocTypeAbbr       *string    `json:"SURVEYDOCTYPE_ABBR"`
	SurveyDocTypeAbbr_      *string    `json:"SURVEYDOCTYPE_ABBR_"`
	SurveyDocTypeNameTH     *string    `json:"SURVEYDOCTYPE_NAME_TH"`
	SurveyDocTypeGroup      *string    `json:"SURVEYDOCTYPE_GROUP"`
	SurveyDocTypeGroup_     *string    `json:"SURVEYDOCTYPE_GROUP_"`
	CadastralImageOrder     *int64     `json:"CADASTRAL_IMAGE_ORDER"`
	CadastralImagePno       *int64     `json:"CADASTRAL_IMAGE_PNO"`
	CadastralImagePno_      *int64     `json:"CADASTRAL_IMAGE_PNO_"`
	CadastralImageThumbnail *string    `json:"CADASTRAL_IMAGE_THUMBNAIL_URL"`
	CadastralImageUrl       *string    `json:"CADASTRAL_IMAGE_URL"`
	CadastralImagePath      *string    `json:"CADASTRAL_IMAGE_PATH"`
	CadastralImageFilename  *string    `json:"CADASTRAL_IMAGE_FILENAME"`
	CadastralImagePname     *string    `json:"CADASTRAL_IMAGE_PNAME"`
	ImageSize               *int64     `json:"IMAGE_SIZE"`
	ImageDpiX               *int64     `json:"IMAGE_DPI_X"`
	ImageDpiY               *int64     `json:"IMAGE_DPI_Y"`
	ImageResX               *int64     `json:"IMAGE_RES_X"`
	ImageResY               *int64     `json:"IMAGE_RES_Y"`
	DetectX_                *int64     `json:"DETECT_X_"`
	DetectY_                *int64     `json:"DETECT_Y_"`
	ProcessSeq_             *int64     `json:"PROCESS_SEQ_"`
	StatusSeq_              *string    `json:"STATUS_SEQ_"`
	CadastralImageNote      *string    `json:"CADASTRAL_IMAGE_NOTE"`
	RecordStatus            *string    `json:"RECORD_STATUS"`
	CreateUser              *string    `json:"CREATE_USER"`
	CreateDtm               *time.Time `json:"CREATE_DTM"`
	LastUpdUser             *string    `json:"LAST_UPD_USER"`
	LastUpdDtm              *time.Time `json:"LAST_UPD_DTM"`
	SyncOriginSeq           *int64     `json:"SYNC_ORIGIN_SEQ"`
	SyncDesSeq              *int64     `json:"SYNC_DES_SEQ"`
	MigrateCadasSeq         *int64     `json:"MIGRATE_CADAS_SEQ"`
	ImageProcessSeq_        *int64     `json:"IMAGE_PROCESS_SEQ_"`
	ImageNote               *string    `json:"IMAGE_NOTE"`
	IpAddress               *string    `json:"IP_ADDRESS"`
	Image1Dtm               *time.Time `json:"IMAGE_1_DTM"`
	Image2Dtm               *time.Time `json:"IMAGE_2_DTM"`
	Image3Dtm               *time.Time `json:"IMAGE_3_DTM"`
	ImageSize_              *int64     `json:"IMAGE_SIZE_"`
	ImageType               *string    `json:"IMAGE_TYPE"`
	ImageHorizontal         *int64     `json:"IMAGE_HORIZONTAL"`
	ImageVertical           *int64     `json:"IMAGE_VERTICAL"`
	Image24bit              *int64     `json:"IMAGE_24BIT"`
	ImageWidth              *int64     `json:"IMAGE_WIDTH"`
	ImageHeight             *int64     `json:"IMAGE_HEIGHT"`
	ImageCrop               *int64     `json:"IMAGE_CROP"`
	ImageLost               *int64     `json:"IMAGE_LOST"`
	ImageCorrupted          *int64     `json:"IMAGE_CORRUPTED"`
	ImageReverse            *string    `json:"IMAGE_REVERSE"`
	ImageSkew               *string    `json:"IMAGE_SKEW"`
	ImageEditCopy           *string    `json:"IMAGE_EDIT_COPY"`
	CiracoreImageNote       *string    `json:"CIRACORE_IMAGE_NOTE"`
	CiracoreImageStatus     *string    `json:"CIRACORE_IMAGE_STATUS"`
	CiracoreImageDtm        *time.Time `json:"CIRACORE_IMAGE_DTM"`
	CiracoreDataNote        *string    `json:"CIRACORE_DATA_NOTE"`
	CiracoreDataStatus      *int64     `json:"CIRACORE_DATA_STATUS"`
	CiracoreDataDtm         *time.Time `json:"CIRACORE_DATA_DTM"`
	OcrDataNote             *string    `json:"OCR_DATA_NOTE"`
	OcrDataStatus           *int64     `json:"OCR_DATA_STATUS"`
	OcrDataDtm              *time.Time `json:"OCR_DATA_DTM"`
}
