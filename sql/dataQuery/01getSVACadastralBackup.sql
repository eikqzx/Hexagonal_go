SELECT * FROM (
SELECT 
 t.CADASTRAL_SEQ,
 t.CADASTRAL_NO,
 t.SHEETCODE,
 t.BOX_NO,
 t.NUMOFSURVEY_QTY,
 t.LANDOFFICE_SEQ,
 t.PRIVATESURVEY_FLAG,
 t.TYPEOFSURVEY_SEQ,
 t.TYPEOFSURVEY_NAME_TH,
 t.TYPEOFSURVEY_ADD1_SEQ,
 t.TYPEOFSURVEY_ADD2_SEQ,
 t.TYPEOFSURVEY_ADD3_SEQ,
 t.CADASTRAL_SURVEY_NO,
 t.CADASTRAL_SURVEY_NO_,
 t.CADASTRAL_TAMBOL_SEQ,
 t.CADASTRAL_TAMBOL_NAME_TH,
 t.CADASTRAL_AMPHUR_SEQ,
 t.CADASTRAL_AMPHUR_NAME_TH,
 t.CADASTRAL_PROVINCE_SEQ,
 t.CADASTRAL_PROVINCE_NAME_TH,
 t.ZONE_LAND,
 t.SHEETTYPE_SEQ,
 t.SHEETTYPE_NAME_TH,
 t.CADASTRAL_UTMMAP,
 t.CADASTRAL_UTMMAP1,
 t.CADASTRAL_UTMMAP2,
 t.CADASTRAL_UTMMAP3,
 t.CADASTRAL_UTMMAP4,
 t.CADASTRAL_ORIGINMAP,
 t.CADASTRAL_ORIGINMAP1,
 t.CADASTRAL_ORIGINMAP2,
 t.CADASTRAL_ORIGINMAP3,
 t.CADASTRAL_LAND_NO,
 t.CADASTRAL_LAND_NO_,
 t.AIRPHOTOMAP_NAME,
 t.AIRPHOTOMAP,
 t.AIRPHOTOMAP1,
 t.AIRPHOTOMAP2,
 t.AIRPHOTOMAP3,
 t.SCALE_MAP_SEQ,
 t.SCALE_MAP,
 t.SCALE_RAWANG_SEQ,
 t.SCALE_RAWANG,
 t.OWNER_TITLE_SEQ,
 t.OWNER_TITLE_NAME,
 t.OWNER_NAME,
 t.OWNER_FNAME,
 t.OWNER_LNAME,
 t.BENCHMARK_SEQ,
 t.BENCHMARK_NAME_TH,
 t.BENCHMARK_QTY,
 t.BENCHMARK2_SEQ,
 t.BENCHMARK2_NAME_TH,
 t.BENCHMARK2_QTY,
 TO_CHAR(t.SURVEY_DTM, 'YYYY-MM-DD HH24:MI:SS') AS SURVEY_DTM,
 TO_CHAR(t.SURVEY_DTM_, 'YYYY-MM-DD HH24:MI:SS') AS SURVEY_DTM_,
 t.SURVEYJOB_SEQ,
 t.TITLE_SEQ,
 t.TITLE_NAME,
 t.SURVEYOR_NAME,
 t.SURVEYOR_FNAME,
 t.SURVEYOR_LNAME,
 t.SURVEYOR_POSITION,
 t.SURVEYOR_POSITION_,
 t.SURVEYOR_LEVEL,
 t.OLD_RAI_NUM,
 t.OLD_NGAN_NUM,
 t.OLD_WA_NUM,
 t.OLD_SUBWA_NUM,
 t.CADASTRAL_LAND_QTY,
 t.CADASTRAL_IMAGE_QTY,
 t.CADASTRAL_OWNER_QTY,
 t.LOST_FLAG,
 t.CANCEL_FLAG,
 t.CANCEL_USER,
 t.CANCEL_CAUSE,
 t.CANCELJOB_FLAG,
 t.CADASTRAL_SEQ_,
 t.CADASTRAL_LOG_SEQ,
 t.PROCESS_SEQ_,
 t.STATUS_SEQ_,
 t.CADASTRAL_NOTE,
 t.RECORD_STATUS,
 t.CREATE_USER,
 t.CREATE_DTM,
 t.LAST_UPD_USER,
 t.LAST_UPD_DTM,
 t108.STATUS_108 ,
 TOFFICE.LANDOFFICE_NAME_TH
 FROM SVA_.TB_SVA_CADASTRAL t
 LEFT JOIN MAS_.TB_MAS_LANDOFFICE TOFFICE
 ON t.LANDOFFICE_SEQ = TOFFICE.LANDOFFICE_SEQ
 LEFT JOIN (
 	SELECT CADASTRAL_SEQ , STATUS_SEQ_ AS STATUS_108 FROM SVA_.TB_SVA_CADASTRAL_LOG WHERE PROCESS_SEQ_ = 108
 ) t108
 ON t.CADASTRAL_SEQ = t108.CADASTRAL_SEQ
 INNER JOIN EVD_.TB_EVD_CADASTRAL_IMAGE TINDEX 
 ON TINDEX.LANDOFFICE_SEQ = T.LANDOFFICE_SEQ 
 AND TINDEX.SHEETCODE = T.SHEETCODE 
 AND TINDEX.BOX_NO = T.BOX_NO 
 AND TINDEX.CADASTRAL_NO = T.CADASTRAL_NO 
 WHERE t.RECORD_STATUS = 'N'
 AND TINDEX.CADASTRAL_IMAGE = 1
 AND (:LANDOFFICE_SEQ IS NULL OR t.LANDOFFICE_SEQ = :LANDOFFICE_SEQ)
 AND (:SHEETCODE IS NULL OR t.SHEETCODE = :SHEETCODE)
 AND (:BOX_NO IS NULL OR t.BOX_NO = :BOX_NO)
 AND (:CADASTRAL_NO IS NULL OR t.CADASTRAL_NO >= :CADASTRAL_NO)
 AND (:CADASTRAL_NO_ IS NULL OR t.CADASTRAL_NO <= :CADASTRAL_NO_)
 AND (:CADASTRAL_SEQ IS NULL OR t.CADASTRAL_SEQ = :CADASTRAL_SEQ)
 AND (:NUMOFSURVEY_QTY IS NULL OR t.NUMOFSURVEY_QTY = :NUMOFSURVEY_QTY)
 ORDER BY t.LANDOFFICE_SEQ, t.SHEETCODE, t.BOX_NO, t.CADASTRAL_NO, t.NUMOFSURVEY_QTY
 ) t
 WHERE (:STATUS_108 IS NULL OR t.STATUS_108 = :STATUS_108)
 