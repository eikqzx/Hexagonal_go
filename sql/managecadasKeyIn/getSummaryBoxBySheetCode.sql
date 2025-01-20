SELECT 
 BOX_NO , 
 COUNT(CADASTRAL_NO) CADASTRAL_COUNT 
FROM SVA_.TB_SVA_CADASTRAL
WHERE 
 LANDOFFICE_SEQ = :LANDOFFICE_SEQ  
 AND SHEETCODE  = :SHEETCODE
GROUP BY BOX_NO  
ORDER BY BOX_NO