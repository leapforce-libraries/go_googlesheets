package GoogleSheets

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
)

type Values struct {
	Range          string     `json:"range"`
	MajorDimension string     `json:"majorDimension"`
	Values         [][]string `json:"values"`
}

type ValuesProperties struct {
	Title string `json:"title"`
}

func (gs *GoogleSheets) GetValues(spreadSheetID string, sheetName string, firstColumn string, lastColumn string, majorDimension string) (*Values, *errortools.Error) {
	batchRowSize := 100
	batchCount := 0

	values := Values{}

	for true {

		aRange := fmt.Sprintf("%s!%s%v:%s%v", sheetName, firstColumn, batchCount*batchRowSize+1, lastColumn, (batchCount+1)*batchRowSize)

		url := fmt.Sprintf("%s/spreadsheets/%s/values/%s?majorDimension=%s", apiURL, spreadSheetID, aRange, majorDimension)
		//fmt.Println(url)

		values_ := Values{}

		_, _, e := gs.Client.Get(url, &values_)
		if e != nil {
			return nil, e
		}

		if len(values_.Values) > 0 {
			values.Values = append(values.Values, values_.Values...)
			batchCount++
		} else {
			break
		}
	}

	return &values, nil
}
