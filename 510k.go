package fda

import (
	"net/http"
)

type Fda510kService struct {
	client *Client
}

type Fda510k struct {
	Address1                string   `json:"address_1,omitempty"`
	Address2                string   `json:"address_2,omitempty"`
	AdvisoryCommittee       string   `json:"advisory_committee,omitempty"`
	AdvisoryCommitteeDesc   string   `json:"advisory_committee_description,omitempty"`
	Applicant               string   `json:"applicant,omitempty"`
	City                    string   `json:"city,omitempty"`
	ClearanceType           string   `json:"clearance_type,omitempty"`
	Contact                 string   `json:"contact,omitempty"`
	CountryCode             string   `json:"country_code,omitempty"`
	DateReceived            string   `json:"date_received,omitempty"`
	DecisionCode            string   `json:"decision_code,omitempty"`
	DecisionDate            string   `json:"decision_date,omitempty"`
	DecisionDescription     string   `json:"decision_description,omitempty"`
	DeviceName              string   `json:"device_name,omitempty"`
	ExpeditedReviewFlag     string   `json:"expedited_review_flag,omitempty"`
	KNumber                 string   `json:"k_number,omitempty"`
	OpenFDA                 *OpenFDA `json:"openfda,omitempty"`
	PostalCode              string   `json:"postal_code,omitempty"`
	ProductCode             string   `json:"product_code,omitempty"`
	ReviewAdvisoryCommittee string   `json:"review_advisory_committee,omitempty"`
	State                   string   `json:"state,omitempty"`
	StatementOrSummary      string   `json:"statement_or_summary,omitempty"`
	ThirdPartyFlag          string   `json:"third_party_flag,omitempty"`
	ZipCode                 string   `json:"zip_code,omitempty"`
	Meta                    *Meta    `json:"meta,omitempty"`
}

type Fda510kOptions struct {
	QueryParameters
	Address1                *string  `url:"address_1,omitempty" json:"address_1,omitempty"`
	Address2                *string  `url:"address_2,omitempty" json:"address_2,omitempty"`
	AdvisoryCommittee       *string  `url:"advisory_committee,omitempty" json:"advisory_committee,omitempty"`
	AdvisoryCommitteeDesc   *string  `url:"advisory_committee_description,omitempty" json:"advisory_committee_description,omitempty"`
	Applicant               *string  `url:"applicant,omitempty" json:"applicant,omitempty"`
	City                    *string  `url:"city,omitempty" json:"city,omitempty"`
	ClearanceType           *string  `url:"clearance_type,omitempty" json:"clearance_type,omitempty"`
	Contact                 *string  `url:"contact,omitempty" json:"contact,omitempty"`
	CountryCode             *string  `url:"country_code,omitempty" json:"country_code,omitempty"`
	DateReceived            *string  `url:"date_received,omitempty" json:"date_received,omitempty"`
	DecisionCode            *string  `url:"decision_code,omitempty" json:"decision_code,omitempty"`
	DecisionDate            *string  `url:"decision_date,omitempty" json:"decision_date,omitempty"`
	DecisionDescription     *string  `url:"decision_description,omitempty" json:"decision_description,omitempty"`
	DeviceName              *string  `url:"device_name,omitempty" json:"device_name,omitempty"`
	ExpeditedReviewFlag     *string  `url:"expedited_review_flag,omitempty" json:"expedited_review_flag,omitempty"`
	KNumber                 *string  `url:"k_number,omitempty" json:"k_number,omitempty"`
	OpenFDA                 *OpenFDA `url:"openfda,omitempty" json:"openfda,omitempty"`
	PostalCode              *string  `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	ProductCode             *string  `url:"product_code,omitempty" json:"product_code,omitempty"`
	ReviewAdvisoryCommittee *string  `url:"review_advisory_committee,omitempty" json:"review_advisory_committee,omitempty"`
	State                   *string  `url:"state,omitempty" json:"state,omitempty"`
	StatementOrSummary      *string  `url:"statement_or_summary,omitempty" json:"statement_or_summary,omitempty"`
	ThirdPartyFlag          *string  `url:"third_party_flag,omitempty" json:"third_party_flag,omitempty"`
	ZipCode                 *string  `url:"zip_code,omitempty" json:"zip_code,omitempty"`
	Meta                    *Meta    `url:"meta,omitempty" json:"meta,omitempty"`
}

type Fda510kResponse struct {
	Results []*Fda510k `json:"results,omitempty"`
	Meta    *Meta      `json:"meta,omitempty"`
}

func (s *Fda510kService) Get510k(opt *Fda510kOptions, options ...RequestOptionFunc) (Fda510kResponse, *Response, error) {
	var result Fda510kResponse
	u := devicePath + fda510kRoute
	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return result, nil, err
	}

	resp, err := s.client.Do(req, &result)
	if err != nil {
		return result, nil, err
	}

	return result, resp, nil

}
