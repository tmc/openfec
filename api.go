package openfec

import "encoding/json"

// BaseURL is the common URL prefix for all API requests.
var BaseURL = "https://api.open.fec.gov/v1"

type apiResponse struct {
	ApiVersion string `json:"api_version,omitempty"`
	Pagination struct {
		Count   float64 `json:"count,omitempty"`
		Page    float64 `json:"page,omitempty"`
		Pages   float64 `json:"pages,omitempty"`
		PerPage float64 `json:"per_page,omitempty"`
	} `json:"pagination,omitempty"`
	Results json.RawMessage `json:"results"`
}
