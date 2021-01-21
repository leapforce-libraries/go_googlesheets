package googlesheets

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

type SpreadSheet struct {
	SpreadsheetId  string                `json:"spreadsheetId"`
	Properties     SpreadSheetProperties `json:"properties"`
	Sheets         []Sheet               `json:"sheets"`
	SpreadsheetURL string                `json:"spreadsheetUrl"`
	DataSources    []DataSource          `json:"dataSources"`
	//DataSourceSchedules []SpreadSheetDataSourceSchedule `json:"dataSourceSchedules"`
}

type SpreadSheetProperties struct {
	Title string `json:"title"`
}

type Sheet struct {
	SheetID int    `json:"sheetId"`
	Title   string `json:"title"`
	Index   int    `json:"index"`
}

type DataSource struct {
	DataSourceID string         `json:"dataSourceId"`
	Spec         DataSourceSpec `json:"spec"`
	SheetID      int            `json:"sheetId"`
}

type DataSourceSpec struct {
	//Parameters []DataSourceParameter  `json:"parameters"`
	BigQuery BigQueryDataSourceSpec `json:"bigQuery"`
}

type BigQueryDataSourceSpec struct {
	ProjectID string            `json:"projectId"`
	QuerySpec BigQueryQuerySpec `json:"querySpec"`
	TableSpec BigQueryTableSpec `json:"tableSpec"`
}

type BigQueryQuerySpec struct {
	RawQuery string `json:"rawQuery"`
}

type BigQueryTableSpec struct {
	TableProjectID string `json:"tableProjectId"`
	TableID        string `json:"tableId"`
	DatasetID      string `json:"datasetId"`
}

func (service *Service) GetSpreadSheet(spreadSheetID string, includeGridData bool) (*SpreadSheet, *errortools.Error) {
	spreadSheet := SpreadSheet{}

	requestConfig := oauth2.RequestConfig{
		URL:           service.url(fmt.Sprintf("spreadsheets/%s?includeGridData=%v", spreadSheetID, includeGridData)),
		ResponseModel: &spreadSheet,
	}
	_, _, e := service.googleService.Get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &spreadSheet, nil
}

func (service *Service) CreateSpreadSheet(spreadSheetID string, includeGridData bool) (*SpreadSheet, *errortools.Error) {
	spreadSheet := SpreadSheet{}

	requestConfig := oauth2.RequestConfig{
		URL:           service.url("spreadsheets"),
		ResponseModel: &spreadSheet,
	}
	_, _, e := service.googleService.Post(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &spreadSheet, nil
}
