

INSERT /*+ APPEND */ 
INTO MAS_.TB_MAS_PROVINCE INS_TBL ( 
    PROVINCE_SEQ, 
    PROVINCE_ID, PROVINCE_ABBR, 
    PROVINCE_NAME_TH, PROVINCE_NAME_EN, 
    PROVINCE_TYPE_SEQ, PROVINCE_NOTE, 
    RECORD_STATUS, CREATE_USER, 
    CREATE_DTM
    --LAST_UPD_USER, LAST_UPD_DTM 
) VALUES ( 
    MAS_.SQ_MAS_PROVINCE_PROVINCE_SEQ.NEXTVAL, 
    :PROVINCE_ID, :PROVINCE_ABBR, 
    :PROVINCE_NAME_TH, :PROVINCE_NAME_EN, 
    :PROVINCE_TYPE_SEQ, :PROVINCE_NOTE, 
    :RECORD_STATUS, :CREATE_USER, 
    SYSDATE --:CREATE_DTM, 
    --:LAST_UPD_USER, :LAST_UPD_DTM 
) 


