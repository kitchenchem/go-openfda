package fda

type Enforcement struct {
	Address1                 string   `json:"address_1,omitempty"`
	Address2                 string   `json:"address_2,omitempty"`
	CenterClassificationDate string   `json:"center_classification_date,omitempty"`
	City                     string   `json:"city,omitempty"`
	Classification           string   `json:"classification,omitempty"`
	CodeInfo                 string   `json:"code_info,omitempty"`
	Country                  string   `json:"country,omitempty"`
	DistributionPattern      string   `json:"distribution_pattern,omitempty"`
	EventID                  string   `json:"event_id,omitempty"`
	InitialFirmNotification  string   `json:"initial_firm_notification,omitempty"`
	MoreCodeInfo             string   `json:"more_code_info,omitempty"`
	OpenFDA                  *OpenFDA `json:"openfda,omitempty"`
	ProductCode              string   `json:"product_code,omitempty"`
	ProductDescription       string   `json:"product_description,omitempty"`
	ProductQuantity          string   `json:"product_quantity,omitempty"`
	ProductType              string   `json:"product_type,omitempty"`
	ReasonForRecall          string   `json:"reason_for_recall,omitempty"`
	RecallInitiationDate     string   `json:"recall_initiation_date,omitempty"`
	RecallNumber             string   `json:"recall_number,omitempty"`
	RecallingFirm            string   `json:"recalling_firm,omitempty"`
	ReportDate               string   `json:"report_date,omitempty"`
	State                    string   `json:"state,omitempty"`
	Status                   string   `json:"status,omitempty"`
	TerminationDate          string   `json:"termination_date,omitempty"`
	VoluntaryMandated        string   `json:"voluntary_mandated,omitempty"`
	Meta                     *Meta    `json:"meta,omitempty"`
}
