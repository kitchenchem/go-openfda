package fda

import "net/http"

type EnforcementService struct {
	client *Client
}

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

type EnforcementOptions struct {
	QueryParameters
	Address1                 *string  `url:"address_1,omitempty" json:"address_1,omitempty"`
	Address2                 *string  `url:"address_2,omitempty" json:"address_2,omitempty"`
	CenterClassificationDate *string  `url:"center_classification_date,omitempty" json:"center_classification_date,omitempty"`
	City                     *string  `url:"city,omitempty" json:"city,omitempty"`
	Classification           *string  `url:"classification,omitempty" json:"classification,omitempty"`
	CodeInfo                 *string  `url:"code_info,omitempty" json:"code_info,omitempty"`
	Country                  *string  `url:"country,omitempty" json:"country,omitempty"`
	DistributionPattern      *string  `url:"distribution_pattern,omitempty" json:"distribution_pattern,omitempty"`
	EventID                  *string  `url:"event_id,omitempty" json:"event_id,omitempty"`
	InitialFirmNotification  *string  `url:"initial_firm_notification,omitempty" json:"initial_firm_notification,omitempty"`
	MoreCodeInfo             *string  `url:"more_code_info,omitempty" json:"more_code_info,omitempty"`
	OpenFDA                  *OpenFDA `url:"openfda,omitempty" json:"openfda,omitempty"`
	ProductCode              *string  `url:"product_code,omitempty" json:"product_code,omitempty"`
	ProductDescription       *string  `url:"product_description,omitempty" json:"product_description,omitempty"`
	ProductQuantity          *string  `url:"product_quantity,omitempty" json:"product_quantity,omitempty"`
	ProductType              *string  `url:"product_type,omitempty" json:"product_type,omitempty"`
	ReasonForRecall          *string  `url:"reason_for_recall,omitempty" json:"reason_for_recall,omitempty"`
	RecallInitiationDate     *string  `url:"recall_initiation_date,omitempty" json:"recall_initiation_date,omitempty"`
	RecallNumber             *string  `url:"recall_number,omitempty" json:"recall_number,omitempty"`
	RecallingFirm            *string  `url:"recalling_firm,omitempty" json:"recalling_firm,omitempty"`
	ReportDate               *string  `url:"report_date,omitempty" json:"report_date,omitempty"`
	State                    *string  `url:"state,omitempty" json:"state,omitempty"`
	Status                   *string  `url:"status,omitempty" json:"status,omitempty"`
	TerminationDate          *string  `url:"termination_date,omitempty" json:"termination_date,omitempty"`
	VoluntaryMandated        *string  `url:"voluntary_mandated,omitempty" json:"voluntary_mandated,omitempty"`
	Meta                     *Meta    `url:"meta,omitempty" json:"meta,omitempty"`
}

type EnforcementResponse struct {
	Results []*Enforcement `json:"results,omitempty"`
	Meta    *Meta          `json:"meta,omitempty"`
}

func (s *EnforcementService) GetUdi(opt *EnforcementOptions) (EnforcementResponse, *Response, error) {
	var result EnforcementResponse
	u := devicePath + enforcementRoute
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
