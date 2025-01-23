package model

type SummarySheetCode struct {
	SHEETCODE        int64 `json:"SHEETCODE"`
	UNIQUE_BOX_COUNT int64 `json:"UNIQUE_BOX_COUNT"`
	CADASTRAL_COUNT  int64 `json:"CADASTRAL_COUNT"`
}

type SummaryBoxBySheetCode struct {
	BOX_NO          string `json:"BOX_NO"`
	CADASTRAL_COUNT int64  `json:"CADASTRAL_COUNT"`
}

type CadastralKeyin struct {
	CADASTRAL_NO    int64 `json:"CADASTRAL_NO"`
	NUMOFSURVEY_QTY int64 `json:"NUMOFSURVEY_QTY"`
	CADASTRAL_SEQ   int64 `json:"CADASTRAL_SEQ"`
}

type ListAllKeyin struct {
	BOX_NO          string `json:"CADASTRAL_SEQ"`
	CADASTRAL_NO    int64  `json:"SHEETCODE"`
	CADASTRAL_SEQ   int64  `json:"BOX_NO"`
	NUMOFSURVEY_QTY int64  `json:"CADASTRAL_NO"`
	ROWNUMBER       int64  `json:"NUMOFSURVEY_QTY"`
	SHEETCODE       int64  `json:"ROWNUMBER"`
}
