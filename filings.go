package openfec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Filing struct {
	AmendmentIndicator           string       `json:"amendment_indicator,omitempty"`
	BeginningImageNumber         string       `json:"beginning_image_number,omitempty"`
	CandidateID                  string       `json:"candidate_id,omitempty"`
	CandidateName                string       `json:"candidate_name,omitempty"`
	CashOnHandBeginningPeriod    int          `json:"cash_on_hand_beginning_period,omitempty"`
	CashOnHandEndPeriod          int          `json:"cash_on_hand_end_period,omitempty"`
	Committee                    string       `json:"committee,omitempty"`
	CommitteeID                  string       `json:"committee_id,omitempty"`
	CommitteeName                string       `json:"committee_name,omitempty"`
	CoverageEndDate              string       `json:"coverage_end_date,omitempty"`
	CoverageStartDate            string       `json:"coverage_start_date,omitempty"`
	Cycle                        int          `json:"cycle,omitempty"`
	DebtsOwedByCommittee         int          `json:"debts_owed_by_committee,omitempty"`
	DebtsOwedToCommittee         int          `json:"debts_owed_to_committee,omitempty"`
	DocumentDescription          string       `json:"document_description,omitempty"`
	DocumentType                 DocumentType `json:"document_type,omitempty"`
	DocumentTypeFull             string       `json:"document_type_full,omitempty"`
	ElectionYear                 int          `json:"election_year,omitempty"`
	EndingImageNumber            string       `json:"ending_image_number,omitempty"`
	FileNumber                   int          `json:"file_number,omitempty"`
	FormType                     string       `json:"form_type,omitempty"`
	HousePersonalFunds           int          `json:"house_personal_funds,omitempty"`
	NetDonations                 int          `json:"net_donations,omitempty"`
	OppositionPersonalFunds      int          `json:"opposition_personal_funds,omitempty"`
	Pages                        int          `json:"pages,omitempty"`
	PdfURL                       string       `json:"pdf_url,omitempty"`
	PreviousFileNumber           string       `json:"previous_file_number,omitempty"`
	PrimaryGeneralIndicator      string       `json:"primary_general_indicator,omitempty"`
	ReceiptDate                  string       `json:"receipt_date,omitempty"`
	ReportType                   string       `json:"report_type,omitempty"`
	ReportTypeFull               string       `json:"report_type_full,omitempty"`
	ReportYear                   int          `json:"report_year,omitempty"`
	RequestType                  string       `json:"request_type,omitempty"`
	SenatePersonalFunds          int          `json:"senate_personal_funds,omitempty"`
	SubID                        int          `json:"sub_id,omitempty"`
	TotalCommunicationCost       int          `json:"total_communication_cost,omitempty"`
	TotalDisbursements           int          `json:"total_disbursements,omitempty"`
	TotalIndependentExpenditures int          `json:"total_independent_expenditures,omitempty"`
	TotalIndividualContributions int          `json:"total_individual_contributions,omitempty"`
	TotalReceipts                int          `json:"total_receipts,omitempty"`
	TreasurerName                string       `json:"treasurer_name,omitempty"`
	UpdateDate                   string       `json:"update_date,omitempty"`
}

type DocumentType string

const (
	Document24HourContributionNotice                         = DocumentType("2")
	Document48HourContributionNotice                         = DocumentType("4")
	DocumentDebtSettlementStatement                          = DocumentType("A")
	DocumentAcknowledgmentofReceiptofDebtSettlementStatement = DocumentType("B")
	DocumentRFAIDebtSettlementFirstNotice                    = DocumentType("C")
	DocumentCommissionDebtSettlementReview                   = DocumentType("D")
	DocumentCommissionResponseTODebtSettlementRequest        = DocumentType("E")
	DocumentAdministrativeTermination                        = DocumentType("F")
	DocumentDebtSettlementPlanAmendment                      = DocumentType("G")
	DocumentDisavowalNotice                                  = DocumentType("H")
	DocumentDisavowalResponse                                = DocumentType("I")
	DocumentConduitReport                                    = DocumentType("J")
	DocumentTerminationApproval                              = DocumentType("K")
	DocumentRepeatNonFilerNotice                             = DocumentType("L")
	DocumentFilingFrequencyChangeNotice                      = DocumentType("M")
	DocumentPaperAmendmenttoElectronicReport                 = DocumentType("N")
	DocumentAcknowledgmentofFilingFrequencyChange            = DocumentType("O")
	DocumentRFAIDebtSettlementSecond                         = DocumentType("S")
	DocumentMiscellaneousReportTOFEC                         = DocumentType("T")
	DocumentRepeatViolationNotice_441A_Or_441B               = DocumentType("V")
	DocumentNoticeofPaperFiling                              = DocumentType("P")
	DocumentF3LFilingFrequencyChangeNotice                   = DocumentType("R")
	DocumentAcknowledgmentofF3LFilingFrequencyChange         = DocumentType("Q")
	DocumentUnregisteredCommitteeNotice                      = DocumentType("U")
)

type FilingQuery struct {
	CandidateID  string `json:"candidate_id,omitempty"`
	Sort         string `json:"sort,omitempty"`
	SortHideNull bool   `json:"sort_hide_null,omitempty"`
	// TODO(tmc):
	// - beginning_image_number
	// - report_year
	// - primary_general_indicator
	// - committee_id
	// - min_receipt_date
	// - form_type
	// - amendment_indicator
	// - report_type
	// - max_receipt_date
}

type FilingIter struct {
	client  *Client
	query   *FilingQuery
	page    *pagination
	err     error
	index   int
	current []*Filing
}

func newFilingIter(query *FilingQuery, client *Client) (*FilingIter, error) {
	i := &FilingIter{
		query:  query,
		client: client,
		page: &pagination{
			Page:    0,
			PerPage: 100,
		},
	}
	return i, i.err
}

func (i *FilingIter) Next() bool {
	if i.index+1 >= len(i.current) {
		i.page.Page++
		i.current, i.page, i.err = i.client.getFilings(i.query, i.page)
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

func (i *FilingIter) Err() error {
	return i.err
}

func (i *FilingIter) Value() *Filing {
	return i.current[i.index]
}

// GetCandidateFilings fetches filings for a given Candidate ID.
func (c *Client) GetCandidateFilings(candidateID string) (*FilingIter, error) {
	return newFilingIter(&FilingQuery{
		CandidateID: candidateID,
	}, c)
}

func (c *Client) getFilings(query *FilingQuery, pagination *pagination) ([]*Filing, *pagination, error) {
	uri := fmt.Sprintf("/filings")
	resp, err := c.do(uri, query, pagination)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	c.trace("GetFilings", uri, string(body))
	var result apiResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, nil, err
	}

	var filings []*Filing
	err = json.Unmarshal([]byte(result.Results), &filings)
	return filings, result.Pagination, err
}
