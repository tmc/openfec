package openfec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Candidate struct {
	ActiveThrough          int    `json:"active_through,omitempty"`
	CandidateID            string `json:"candidate_id,omitempty"`
	CandidateStatus        string `json:"candidate_status,omitempty"`
	CandidateStatusFull    string `json:"candidate_status_full,omitempty"`
	Cycles                 []int  `json:"cycles,omitempty"`
	District               string `json:"district,omitempty"`
	ElectionYears          []int  `json:"election_years,omitempty"`
	IncumbentChallenge     string `json:"incumbent_challenge,omitempty"`
	IncumbentChallengeFull string `json:"incumbent_challenge_full,omitempty"`
	Name                   string `json:"name,omitempty"`
	Office                 string `json:"office,omitempty"`
	OfficeFull             string `json:"office_full,omitempty"`
	Party                  string `json:"party,omitempty"`
	PartyFull              string `json:"party_full,omitempty"`
	State                  string `json:"state,omitempty"`
}

type IncumbentChallenge string

const (
	Incumbent  = IncumbentChallenge("I")
	Challenger = IncumbentChallenge("C")
	Open       = IncumbentChallenge("O")
)

type CandidateStatus string

const (
	PresentCandidate = CandidateStatus("C")
	FutureCandidate  = CandidateStatus("F")
	NotYetACandidate = CandidateStatus("N")
	PriorCandidate   = CandidateStatus("P")
)

type Office string

const (
	House     = Office("H")
	Senate    = Office("S")
	President = Office("P")
)

type CandidateQuery struct {
	Sort               string               `json:"sort,omitempty"`
	SortHideNull       bool                 `json:"sort_hide_null,omitempty"`
	Party              string               `json:"party,omitempty"`
	IncumbentChallenge []IncumbentChallenge `json:"incumbent_challenge,omitempty"`
	Cycle              []int                `json:"cycle,omitempty"`
	Year               int                  `json:"year,omitempty"`
	CandidateStatus    []CandidateStatus    `json:"candidate_status,omitempty"`
	State              string               `json:"state,omitempty"`
	District           int                  `json:"district,omitempty"`
	Office             []Office             `json:"office,omitempty"`
	Name               string               `json:"name,omitempty"`
	CandidateID        string               `json:"candidate_id,omitempty"`
	Query              string               `json:"q,omitempty"`
}

type CandidateIter struct {
	client  *Client
	query   *CandidateQuery
	page    *pagination
	err     error
	index   int
	current []*Candidate
}

func newCandidateIter(query *CandidateQuery, client *Client) (*CandidateIter, error) {
	i := &CandidateIter{
		query:  query,
		client: client,
		page: &pagination{
			Page:    1,
			PerPage: 100,
		},
	}
	i.current, i.page, i.err = client.getCandidates(i.query, i.page)
	return i, i.err
}

func (i *CandidateIter) Next() bool {
	//spew.Dump("Next:", i.index, i.index+1, len(i.current))
	if i.index+1 >= len(i.current) {
		i.page.Page++
		i.current, i.page, i.err = i.client.getCandidates(i.query, i.page)
		if i.err != nil {
			return false
		}
		if len(i.current) == 0 {
			return false
		}
		i.index = 0
	} else {
		i.index++
	}
	return true
}

func (i *CandidateIter) Err() error {
	return i.err
}

func (i *CandidateIter) Value() *Candidate {
	return i.current[i.index]
}

// GetCandidates fetches a list of basic candidate information.
//
// query is an optional parameter that supplies additional parameters
//
// Fetch basic information about candidates, and use parameters to filter results to the candidates you're looking for.
// Each result reflects a unique FEC candidate ID. That ID is particular to the candidate for a particular office sought. If a candidate runs for the same office multiple times, the ID stays the same. If the same person runs for another office — for example, a House candidate runs for a Senate office — that candidate will get a unique ID for each office.
func (c *Client) GetCandidates(query *CandidateQuery) (*CandidateIter, error) {
	return newCandidateIter(query, c)
}

func (c *Client) getCandidates(query *CandidateQuery, pagination *pagination) ([]*Candidate, *pagination, error) {
	uri := fmt.Sprintf("/candidates")
	resp, err := c.do(uri, query, pagination)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	c.trace("GetCandidates", uri, string(body))
	var result apiResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, nil, err
	}

	var candidates []*Candidate
	err = json.Unmarshal([]byte(result.Results), &candidates)
	return candidates, result.Pagination, err
}
