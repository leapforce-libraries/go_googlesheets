package GoogleSheets

import (
	"net/http"

	bigquerytools "github.com/leapforce-libraries/go_bigquerytools"
	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

const (
	apiName         string = "GoogleSheets"
	apiURL          string = "https://sheets.googleapis.com/v4"
	authURL         string = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenURL        string = "https://oauth2.googleapis.com/token"
	tokenHTTPMethod string = http.MethodPost
	redirectURL     string = "http://localhost:8080/oauth/redirect"
)

// GoogleSheets stores GoogleSheets configuration
//
type GoogleSheets struct {
	oAuth2 *oauth2.OAuth2
}

// methods
//
func NewGoogleSheets(clientID string, clientSecret string, scope string, bigQuery *bigquerytools.BigQuery) *GoogleSheets {
	gd := GoogleSheets{}
	config := oauth2.OAuth2Config{
		ApiName:         apiName,
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		Scope:           scope,
		RedirectURL:     redirectURL,
		AuthURL:         authURL,
		TokenURL:        tokenURL,
		TokenHTTPMethod: tokenHTTPMethod,
	}
	gd.oAuth2 = oauth2.NewOAuth(config, bigQuery)
	return &gd
}

func (gd *GoogleSheets) ValidateToken() (*oauth2.Token, *errortools.Error) {
	return gd.oAuth2.ValidateToken()
}

func (gd *GoogleSheets) InitToken() *errortools.Error {
	return gd.oAuth2.InitToken()
}

func (gd *GoogleSheets) Get(url string, model interface{}) (*http.Response, *errortools.Error) {
	err := google.ErrorResponse{}
	_, res, e := gd.oAuth2.Get(url, model, &err)

	if e != nil {
		if err.Error.Message != "" {
			e.SetMessage(err.Error.Message)
		}
		return nil, e
	}

	return res, nil
}

func (gd *GoogleSheets) Patch(url string, model interface{}) (*http.Response, *errortools.Error) {
	err := google.ErrorResponse{}
	_, res, e := gd.oAuth2.Patch(url, nil, model, &err)

	if e != nil {
		if err.Error.Message != "" {
			e.SetMessage(err.Error.Message)
		}
		return nil, e
	}

	return res, nil
}
