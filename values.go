package googlesheets

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

type Values struct {
	Range          string     `json:"range"`
	MajorDimension string     `json:"majorDimension"`
	Values         [][]string `json:"values"`
}

type ValuesProperties struct {
	Title string `json:"title"`
}

func (service *Service) GetValues(spreadSheetID string, sheetName string, firstColumn string, lastColumn string, majorDimension string) (*Values, *errortools.Error) {
	batchRowSize := 100
	batchCount := 0

	values := Values{}

	for true {

		aRange := fmt.Sprintf("%s!%s%v:%s%v", sheetName, firstColumn, batchCount*batchRowSize+1, lastColumn, (batchCount+1)*batchRowSize)

		_values := Values{}

		requestConfig := oauth2.RequestConfig{
			URL:           service.url(fmt.Sprintf("spreadsheets/%s/values/%s?majorDimension=%s", spreadSheetID, aRange, majorDimension)),
			ResponseModel: &_values,
		}
		_, _, e := service.googleService.Get(&requestConfig)
		if e != nil {
			return nil, e
		}

		if len(_values.Values) > 0 {
			values.Values = append(values.Values, _values.Values...)
			batchCount++
		} else {
			break
		}
	}

	return &values, nil
}
