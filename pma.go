package fda

import "net/http"

type PmaService struct {
	client *Client
}

type Pma struct {
	AdvisoryCommittee            string   `json:"advisory_committee,omitempty"`
	AdvisoryCommitteeDescription string   `json:"advisory_committee_description,omitempty"`
	AOStatement                  string   `json:"ao_statement,omitempty"`
	Applicant                    string   `json:"applicant,omitempty"`
	City                         string   `json:"city,omitempty"`
	DateReceived                 string   `json:"date_received,omitempty"`
	DecisionCode                 string   `json:"decision_code,omitempty"`
	DecisionDate                 string   `json:"decision_date,omitempty"`
	DocketNumber                 string   `json:"docket_number,omitempty"`
	ExpeditedReviewFlag          string   `json:"expedited_review_flag,omitempty"`
	FedRegNoticeDate             string   `json:"fed_reg_notice_date,omitempty"`
	GenericName                  string   `json:"generic_name,omitempty"`
	OpenFDA                      *OpenFDA `json:"openfda,omitempty"`
	PMANumber                    string   `json:"pma_number,omitempty"`
	ProductCode                  string   `json:"product_code,omitempty"`
	State                        string   `json:"state,omitempty"`
	Street1                      string   `json:"street_1,omitempty"`
	Street2                      string   `json:"street_2,omitempty"`
	SupplementNumber             string   `json:"supplement_number,omitempty"`
	SupplementReason             string   `json:"supplement_reason,omitempty"`
	SupplementType               string   `json:"supplement_type,omitempty"`
	TradeName                    string   `json:"trade_name,omitempty"`
	Zip                          string   `json:"zip,omitempty"`
	ZipExt                       string   `json:"zip_ext,omitempty"`
}

type PmaOptions struct {
	QueryParameters
	AdvisoryCommittee            *string  `url:"advisory_committee,omitempty" json:"advisory_committee,omitempty"`
	AdvisoryCommitteeDescription *string  `url:"advisory_committee_description,omitempty" json:"advisory_committee_description,omitempty"`
	AOStatement                  *string  `url:"ao_statement,omitempty" json:"ao_statement,omitempty"`
	Applicant                    *string  `url:"applicant,omitempty" json:"applicant,omitempty"`
	City                         *string  `url:"city,omitempty" json:"city,omitempty"`
	DateReceived                 *string  `url:"date_received,omitempty" json:"date_received,omitempty"`
	DecisionCode                 *string  `url:"decision_code,omitempty" json:"decision_code,omitempty"`
	DecisionDate                 *string  `url:"decision_date,omitempty" json:"decision_date,omitempty"`
	DocketNumber                 *string  `url:"docket_number,omitempty" json:"docket_number,omitempty"`
	ExpeditedReviewFlag          *string  `url:"expedited_review_flag,omitempty" json:"expedited_review_flag,omitempty"`
	FedRegNoticeDate             *string  `url:"fed_reg_notice_date,omitempty" json:"fed_reg_notice_date,omitempty"`
	GenericName                  *string  `url:"generic_name,omitempty" json:"generic_name,omitempty"`
	OpenFDA                      *OpenFDA `url:"openfda,omitempty" json:"openfda,omitempty"`
	PMANumber                    *string  `url:"pma_number,omitempty" json:"pma_number,omitempty"`
	ProductCode                  *string  `url:"product_code,omitempty" json:"product_code,omitempty"`
	State                        *string  `url:"state,omitempty" json:"state,omitempty"`
	Street1                      *string  `url:"street_1,omitempty" json:"street_1,omitempty"`
	Street2                      *string  `url:"street_2,omitempty" json:"street_2,omitempty"`
	SupplementNumber             *string  `url:"supplement_number,omitempty" json:"supplement_number,omitempty"`
	SupplementReason             *string  `url:"supplement_reason,omitempty" json:"supplement_reason,omitempty"`
	SupplementType               *string  `url:"supplement_type,omitempty" json:"supplement_type,omitempty"`
	TradeName                    *string  `url:"trade_name,omitempty" json:"trade_name,omitempty"`
	Zip                          *string  `url:"zip,omitempty" json:"zip,omitempty"`
	ZipExt                       *string  `url:"zip_ext,omitempty" json:"zip_ext,omitempty"`
}

type PmaResponse struct {
	Results []*Pma `json:"results,omitempty"`
	Meta    *Meta  `json:"meta,omitempty"`
}

func (s *PmaService) Get510k(opt *PmaOptions) (PmaResponse, *Response, error) {
	var result PmaResponse
	u := devicePath + pmaRoute

	req, err := s.client.NewRequest(http.MethodGet, u, opt)
	if err != nil {
		return result, nil, err
	}

	resp, err := s.client.Do(req, &result)
	if err != nil {
		return result, nil, err
	}

	return result, resp, nil

}
