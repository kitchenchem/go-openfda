package fda

import "net/http"

type FDA510kService struct {
	client *Client
}

type FDA510k struct {
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

type FDA510kOptions struct {
	QueryParameters
	Address1                *string  `json:"address_1,omitempty"`
	Address2                *string  `json:"address_2,omitempty"`
	AdvisoryCommittee       *string  `url:"advisory_committee,omitempty" json:"advisory_committee,omitempty"`
	AdvisoryCommitteeDesc   *string  `json:"advisory_committee_description,omitempty"`
	Applicant               *string  `json:"applicant,omitempty"`
	City                    *string  `json:"city,omitempty"`
	ClearanceType           *string  `json:"clearance_type,omitempty"`
	Contact                 *string  `json:"contact,omitempty"`
	CountryCode             *string  `json:"country_code,omitempty"`
	DateReceived            *string  `json:"date_received,omitempty"`
	DecisionCode            *string  `json:"decision_code,omitempty"`
	DecisionDate            *string  `json:"decision_date,omitempty"`
	DecisionDescription     *string  `json:"decision_description,omitempty"`
	DeviceName              *string  `json:"device_name,omitempty"`
	ExpeditedReviewFlag     *string  `json:"expedited_review_flag,omitempty"`
	KNumber                 *string  `json:"k_number,omitempty"`
	OpenFDA                 *OpenFDA `json:"openfda,omitempty"`
	PostalCode              *string  `json:"postal_code,omitempty"`
	ProductCode             *string  `json:"product_code,omitempty"`
	ReviewAdvisoryCommittee *string  `json:"review_advisory_committee,omitempty"`
	State                   *string  `json:"state,omitempty"`
	StatementOrSummary      *string  `json:"statement_or_summary,omitempty"`
	ThirdPartyFlag          *string  `json:"third_party_flag,omitempty"`
	ZipCode                 *string  `json:"zip_code,omitempty"`
	Meta                    *Meta    `json:"meta,omitempty"`
}

func (s *FDA510kService) Get510k(opt *FDA510kOptions) ([]*FDA510k, *Response, error) {
	u := DevicePath + F510k

	req, err := s.client.NewRequest(http.MethodGet, u, opt)
	if err != nil {
		return nil, nil, err
	}

	var f []*FDA510k
	resp, err := s.client.Do(req, &f)
	if err != nil {
		return nil, nil, err
	}
	return f, resp, nil
}
