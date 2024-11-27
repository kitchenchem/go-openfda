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

type RecallOptions struct {
	QueryParameters
	CfresID               *string   `url:"cfres_id,omitempty" json:"cfres_id,omitempty"`
	EventDateInitiated    *string   `url:"event_date_initiated,omitempty" json:"event_date_initiated,omitempty"`
	EventDateCreated      *string   `url:"event_date_created,omitempty" json:"event_date_created,omitempty"`
	EventDatePosted       *string   `url:"event_date_posted,omitempty" json:"event_date_posted,omitempty"`
	EventDateTerminated   *string   `url:"event_date_terminated,omitempty" json:"event_date_terminated,omitempty"`
	RecallStatus          *string   `url:"recall_status,omitempty" json:"recall_status,omitempty"`
	RecallingFirm         *string   `url:"recalling_firm,omitempty" json:"recalling_firm,omitempty"`
	FirmFEINumber         *string   `url:"firm_fei_number,omitempty" json:"firm_fei_number,omitempty"`
	Address1              *string   `url:"address_1,omitempty" json:"address_1,omitempty"`
	Address2              *string   `url:"address_2,omitempty" json:"address_2,omitempty"`
	City                  *string   `url:"city,omitempty" json:"city,omitempty"`
	State                 *string   `url:"state,omitempty" json:"state,omitempty"`
	PostalCode            *string   `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Country               *string   `url:"country,omitempty" json:"country,omitempty"`
	AdditionalInfoContact *string   `url:"additional_info_contact,omitempty" json:"additional_info_contact,omitempty"`
	ReasonForRecall       *string   `url:"reason_for_recall,omitempty" json:"reason_for_recall,omitempty"`
	KNumbers              *[]string `url:"k_numbers,omitempty" json:"k_numbers,omitempty"`
	OpenFDA               *OpenFDA  `url:"openfda,omitempty" json:"openfda,omitempty"`
	OtherSubmissionDesc   *string   `url:"other_submission_description,omitempty" json:"other_submission_description,omitempty"`
	PMANumbers            *[]string `url:"pma_numbers,omitempty" json:"pma_numbers,omitempty"`
	ProductDescription    *string   `url:"product_description,omitempty" json:"product_description,omitempty"`
	CodeInfo              *string   `url:"code_info,omitempty" json:"code_info,omitempty"`
	ProductCode           *string   `url:"product_code,omitempty" json:"product_code,omitempty"`
	ProductResNumber      *string   `url:"product_res_number,omitempty" json:"product_res_number,omitempty"`
	ProductQuantity       *string   `url:"product_quantity,omitempty" json:"product_quantity,omitempty"`
	DistributionPattern   *string   `url:"distribution_pattern,omitempty" json:"distribution_pattern,omitempty"`
	ResEventNumber        *string   `url:"res_event_number,omitempty" json:"res_event_number,omitempty"`
	RootCauseDescription  *string   `url:"root_cause_description,omitempty" json:"root_cause_description,omitempty"`
	Action                *string   `url:"action,omitempty" json:"action,omitempty"`
	Meta                  *Meta     `url:"meta,omitempty" json:"meta,omitempty"`
}
