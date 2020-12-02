package GoogleSheets

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
)

type SpreadSheet struct {
	SpreadsheetId string                `json:"spreadsheetId"`
	Properties    SpreadSheetProperties `json:"properties"`
}

type SpreadSheetProperties struct {
	Title string `json:"title"`
}

func (gs *GoogleSheets) GetSpreadSheet(spreadSheetID string, includeGridData bool) (*SpreadSheet, *errortools.Error) {
	url := fmt.Sprintf("%s/spreadsheets/%s?includeGridData=%v", apiURL, spreadSheetID, includeGridData)
	//fmt.Println(url)

	spreadSheet := SpreadSheet{}

	_, e := gs.Get(url, &spreadSheet)
	if e != nil {
		return nil, e
	}

	return &spreadSheet, nil
}
