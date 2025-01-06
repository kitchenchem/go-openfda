package fda

import "net/http"

type ClassificationService struct {
	client *Client
}

type Classification struct {
	Definition                  string   `json:"definition,omitempty"`
	DeviceClass                 string   `json:"device_class,omitempty"`
	DeviceName                  string   `json:"device_name,omitempty"`
	GMPExemptFlag               string   `json:"gmp_exempt_flag,omitempty"`
	ImplantFlag                 string   `json:"implant_flag,omitempty"`
	LifeSustainSupportFlag      string   `json:"life_sustain_support_flag,omitempty"`
	MedicalSpecialty            string   `json:"medical_specialty,omitempty"`
	MedicalSpecialtyDescription string   `json:"medical_specialty_description,omitempty"`
	OpenFDA                     *OpenFDA `json:"openfda,omitempty"`
	ProductCode                 string   `json:"product_code,omitempty"`
	RegulationNumber            string   `json:"regulation_number,omitempty"`
	ReviewCode                  string   `json:"review_code,omitempty"`
	ReviewPanel                 string   `json:"review_panel,omitempty"`
	SubmissionTypeID            string   `json:"submission_type_id,omitempty"`
	SummaryMalfunctionReporting string   `json:"summary_malfunction_reporting,omitempty"`
	ThirdPartyFlag              string   `json:"third_party_flag,omitempty"`
	UnclassifiedReason          string   `json:"unclassified_reason,omitempty"`
}

type ClassificationOptions struct {
	QueryParameters
	Definition                  *string  `url:"definition,omitempty" json:"definition,omitempty"`
	DeviceClass                 *string  `url:"device_class,omitempty" json:"device_class,omitempty"`
	DeviceName                  *string  `url:"device_name,omitempty" json:"device_name,omitempty"`
	GMPExemptFlag               *string  `url:"gmp_exempt_flag,omitempty" json:"gmp_exempt_flag,omitempty"`
	ImplantFlag                 *string  `url:"implant_flag,omitempty" json:"implant_flag,omitempty"`
	LifeSustainSupportFlag      *string  `url:"life_sustain_support_flag,omitempty" json:"life_sustain_support_flag,omitempty"`
	MedicalSpecialty            *string  `url:"medical_specialty,omitempty" json:"medical_specialty,omitempty"`
	MedicalSpecialtyDescription *string  `url:"medical_specialty_description,omitempty" json:"medical_specialty_description,omitempty"`
	OpenFDA                     *OpenFDA `url:"openfda,omitempty" json:"openfda,omitempty"`
	ProductCode                 *string  `url:"product_code,omitempty" json:"product_code,omitempty"`
	RegulationNumber            *string  `url:"regulation_number,omitempty" json:"regulation_number,omitempty"`
	ReviewCode                  *string  `url:"review_code,omitempty" json:"review_code,omitempty"`
	ReviewPanel                 *string  `url:"review_panel,omitempty" json:"review_panel,omitempty"`
	SubmissionTypeID            *string  `url:"submission_type_id,omitempty" json:"submission_type_id,omitempty"`
	SummaryMalfunctionReporting *string  `url:"summary_malfunction_reporting,omitempty" json:"summary_malfunction_reporting,omitempty"`
	ThirdPartyFlag              *string  `url:"third_party_flag,omitempty" json:"third_party_flag,omitempty"`
	UnclassifiedReason          *string  `url:"unclassified_reason,omitempty" json:"unclassified_reason,omitempty"`
}

type ClassificationResponse struct {
	Results []*Classification `json:"results,omitempty"`
	Meta    *Meta             `json:"meta,omitempty"`
}

func (s *ClassificationService) GetClassification(opt *ClassificationOptions, options ...RequestOptionFunc) (ClassificationResponse, *Response, error) {
	var result ClassificationResponse
	u := devicePath + classificationRoute

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
