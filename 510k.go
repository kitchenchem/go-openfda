package fda

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
