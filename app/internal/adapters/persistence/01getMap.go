package persistence

import (
	"Hexagonal_go/app/internal/core/model"
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
)

type GetMapImpl struct{}

func (r *GetMapImpl) FetchMapLandGIS(ogrFID, landOfficeSeq *int64, utmMap1, utmMap2, utmMap3, utmMap4, utmScale, landNo *string) ([]model.MapLandGIS, error) {
	if DB == nil {
		InitDB()
	}

	queryPath := filepath.Join("sql", "data", "map", "01getmap.sql")
	query, err := ReadSQLFromFile(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL query file: %v", err)
	}
	rows, err := DB.Query(query, sql.Named("OGR_FID", ogrFID), sql.Named("LANDOFFICE_SEQ", landOfficeSeq), sql.Named("UTMMAP1", utmMap1), sql.Named("UTMMAP2", utmMap2), sql.Named("UTMMAP3", utmMap3), sql.Named("UTMMAP4", utmMap4), sql.Named("UTMSCALE", utmScale), sql.Named("LAND_NO", landNo))
	if err != nil {
		return nil, fmt.Errorf("failed to query map land GIS: %v", err)
	}
	defer rows.Close()

	var mapLandGISList []model.MapLandGIS
	for rows.Next() {
		var mapLandGIS model.MapLandGIS
		if err := rows.Scan(
			&mapLandGIS.OGR_FID, &mapLandGIS.LANDOFFICE_SEQ, &mapLandGIS.UTMMAP1, &mapLandGIS.UTMMAP2,
			&mapLandGIS.UTMMAP3, &mapLandGIS.UTMMAP4, &mapLandGIS.UTMSCALE, &mapLandGIS.LAND_NO,
			&mapLandGIS.WKT,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		mapLandGISList = append(mapLandGISList, mapLandGIS)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	log.Printf("Fetched %d rows", len(mapLandGISList))
	return mapLandGISList, nil
}
