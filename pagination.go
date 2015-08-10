package openfec

type pagination struct {
	Count   int `json:"count,omitempty"`
	Page    int `json:"page,omitempty"`
	Pages   int `json:"pages,omitempty"`
	PerPage int `json:"per_page,omitempty"`
}
