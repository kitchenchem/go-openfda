package fda

type Recall struct {
	CfresID               string   `json:"cfres_id,omitempty"`
	EventDateInitiated    string   `json:"event_date_initiated,omitempty"`
	EventDateCreated      string   `json:"event_date_created,omitempty"`
	EventDatePosted       string   `json:"event_date_posted,omitempty"`
	EventDateTerminated   string   `json:"event_date_terminated,omitempty"`
	RecallStatus          string   `json:"recall_status,omitempty"`
	RecallingFirm         string   `json:"recalling_firm,omitempty"`
	FirmFEINumber         string   `json:"firm_fei_number,omitempty"`
	Address1              string   `json:"address_1,omitempty"`
	Address2              string   `json:"address_2,omitempty"`
	City                  string   `json:"city,omitempty"`
	State                 string   `json:"state,omitempty"`
	PostalCode            string   `json:"postal_code,omitempty"`
	Country               string   `json:"country,omitempty"`
	AdditionalInfoContact string   `json:"additional_info_contact,omitempty"`
	ReasonForRecall       string   `json:"reason_for_recall,omitempty"`
	KNumbers              []string `json:"k_numbers,omitempty"`
	OpenFDA               *OpenFDA `json:"openfda,omitempty"`
	OtherSubmissionDesc   string   `json:"other_submission_description,omitempty"`
	PMANumbers            []string `json:"pma_numbers,omitempty"`
	ProductDescription    string   `json:"product_description,omitempty"`
	CodeInfo              string   `json:"code_info,omitempty"`
	ProductCode           string   `json:"product_code,omitempty"`
	ProductResNumber      string   `json:"product_res_number,omitempty"`
	ProductQuantity       string   `json:"product_quantity,omitempty"`
	DistributionPattern   string   `json:"distribution_pattern,omitempty"`
	ResEventNumber        string   `json:"res_event_number,omitempty"`
	RootCauseDescription  string   `json:"root_cause_description,omitempty"`
	Action                string   `json:"action,omitempty"`
	Meta                  *Meta    `json:"meta,omitempty"`
}
