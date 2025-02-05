

MERGE INTO MAS_.TB_MAS_SCALE Target_ 
USING ( 
    SELECT :arg SCALE_SEQ, :SCALE_SYS_ SCALE_SYS_, 
        :SCALE_ID SCALE_ID, :SCALE_ABBR SCALE_ABBR, 
        :SCALE_NAME_TH SCALE_NAME_TH, :SCALE_NAME_EN SCALE_NAME_EN, 
        :SCALE_PAGE_QTY SCALE_PAGE_QTY, 
        :SCALE_NOTE SCALE_NOTE, 
        :RECORD_STATUS RECORD_STATUS, 
        :CREATE_USER CREATE_USER, SYSDATE CREATE_DTM, 
        :CREATE_USER LAST_UPD_USER, SYSDATE LAST_UPD_DTM 
    FROM DUAL 
    WHERE 1 = 1 
) Source_ 
    ON ( Source_.SCALE_SEQ = Target_.SCALE_SEQ ) 

WHEN NOT MATCHED THEN 
    INSERT ( 
        SCALE_SEQ, SCALE_SYS_, 
        SCALE_ID, SCALE_ABBR, 
        SCALE_NAME_TH, SCALE_NAME_EN, 
        SCALE_PAGE_QTY, 
        SCALE_NOTE, 
        RECORD_STATUS, 
        CREATE_USER, CREATE_DTM--, LAST_UPD_USER, LAST_UPD_DTM 
    ) VALUES ( 
        MAS_.SQ_MAS_SCALE_SCALE_SEQ.NEXTVAL, Source_.SCALE_SYS_, 
        Source_.SCALE_ID, Source_.SCALE_ABBR, 
        Source_.SCALE_NAME_TH, Source_.SCALE_NAME_EN, 
        Source_.SCALE_PAGE_QTY, 
        Source_.SCALE_NOTE, 
        Source_.RECORD_STATUS, 
        Source_.CREATE_USER, Source_.CREATE_DTM--, Source_.LAST_UPD_USER, Source_.LAST_UPD_DTM 
    ) 

WHEN MATCHED THEN 
    UPDATE SET 
        -- Target_.SCALE_SEQ = Source_.SCALE_SEQ, 
        Target_.SCALE_SYS_ = Source_.SCALE_SYS_, 
        Target_.SCALE_ID = Source_.SCALE_ID, Target_.SCALE_ABBR = Source_.SCALE_ABBR, 
        Target_.SCALE_NAME_TH = Source_.SCALE_NAME_TH, Target_.SCALE_NAME_EN = Source_.SCALE_NAME_EN, 
        Target_.SCALE_PAGE_QTY = Source_.SCALE_PAGE_QTY, 
        Target_.SCALE_NOTE = Source_.SCALE_NOTE, 
        Target_.RECORD_STATUS = Source_.RECORD_STATUS, 
        -- Target_.CREATE_USER = Source_.CREATE_USER, Target_.CREATE_DTM = Source_.CREATE_DTM, 
        Target_.LAST_UPD_USER = Source_.LAST_UPD_USER,  Target_.LAST_UPD_DTM = Source_.LAST_UPD_DTM 


