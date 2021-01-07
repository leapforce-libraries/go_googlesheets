package GoogleSheets

import (
	google "github.com/leapforce-libraries/go_google"
)

const (
	apiName string = "GoogleSheets"
	apiURL  string = "https://sheets.googleapis.com/v4"
)

// GoogleSheets stores GoogleSheets configuration
//
type GoogleSheets struct {
	Client *google.GoogleClient
}

// methods
//
func NewGoogleSheets(clientID string, clientSecret string, scope string, bigQuery *google.BigQuery) *GoogleSheets {
	config := google.GoogleClientConfig{
		APIName:      apiName,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
	}

	googleClient := google.NewGoogleClient(config, bigQuery)

	return &GoogleSheets{googleClient}
}
