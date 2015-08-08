package openfec

import "encoding/json"

// BaseURL is the common URL prefix for all API requests.
var BaseURL = "https://api.open.fec.gov/v1"

type apiResponse struct {
	ApiVersion string          `json:"api_version,omitempty"`
	Pagination *Pagination     `json:"pagination,omitempty"`
	Results    json.RawMessage `json:"results"`
}
