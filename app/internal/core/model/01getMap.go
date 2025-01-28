package model

// MapLandGIS เป็นโมเดลสำหรับข้อมูล Map Land GIS
type MapLandGIS struct {
	OGR_FID        *int64  `json:"OGR_FID"`
	LANDOFFICE_SEQ *int64  `json:"LANDOFFICE_SEQ"`
	UTMMAP1        *string `json:"UTMMAP1"`
	UTMMAP2        *string `json:"UTMMAP2"`
	UTMMAP3        *string `json:"UTMMAP3"`
	UTMMAP4        *string `json:"UTMMAP4"`
	UTMSCALE       *string `json:"UTMSCALE"`
	LAND_NO        *string `json:"LAND_NO"`
	WKT            *string `json:"WKT"`
}
