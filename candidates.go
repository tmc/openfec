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

// GetCandidates fetches a list of basic candidate information.
//
// Fetch basic information about candidates, and use parameters to filter results to the candidates you're looking for.
// Each result reflects a unique FEC candidate ID. That ID is particular to the candidate for a particular office sought. If a candidate runs for the same office multiple times, the ID stays the same. If the same person runs for another office — for example, a House candidate runs for a Senate office — that candidate will get a unique ID for each office.
func (c *Client) GetCandidates() ([]*Candidate, *Pagination, error) {
	uri := fmt.Sprintf("/candidates")
	resp, err := c.do(uri, nil)
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
	fmt.Println(string(result.Results))

	var candidates []*Candidate
	err = json.Unmarshal([]byte(result.Results), &candidates)
	return candidates, result.Pagination, err
}
