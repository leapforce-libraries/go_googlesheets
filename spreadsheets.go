package GoogleSheets

import (
	"fmt"
)

type SpreadSheet struct {
	SpreadsheetId string                `json:"spreadsheetId"`
	Properties    SpreadSheetProperties `json:"properties"`
}

type SpreadSheetProperties struct {
	Title string `json:"title"`
}

func (gs *GoogleSheets) GetSpreadSheet(spreadSheetID string, includeGridData bool) (*SpreadSheet, error) {
	url := fmt.Sprintf("%s/spreadsheets/%s?includeGridData=%v", apiURL, spreadSheetID, includeGridData)
	//fmt.Println(url)

	spreadSheet := SpreadSheet{}

	_, err := gs.Get(url, &spreadSheet)
	if err != nil {
		return nil, err
	}

	return &spreadSheet, nil
}
