package googlesheets

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
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

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			URL:           service.url(fmt.Sprintf("spreadsheets/%s/values/%s?majorDimension=%s", spreadSheetID, aRange, majorDimension)),
			ResponseModel: &_values,
		}
		_, _, e := service.googleService.HTTPRequest(&requestConfig)
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
