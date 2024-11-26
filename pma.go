package fda

type PMA struct {
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
