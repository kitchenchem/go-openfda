package fda

import "net/http"

type Covid19SerologyService struct {
	client *Client
}

type Covid19Serology struct {
	Manufacturer    string `json:"manufacturer,omitempty"`
	Device          string `json:"device,omitempty"`
	DatePerformed   string `json:"date_performed,omitempty"`
	EvaluationID    string `json:"evaluation_id,omitempty"`
	LotNumber       string `json:"lot_number,omitempty"`
	Panel           string `json:"panel,omitempty"`
	SampleNo        int    `json:"sample_no,omitempty"`
	SampleID        string `json:"sample_id,omitempty"`
	Type            string `json:"type,omitempty"`
	Group           string `json:"group,omitempty"`
	DaysFromSymptom string `json:"days_from_symptom,omitempty"`
	IgMResult       string `json:"igm_result,omitempty"`
	IgGResult       string `json:"igg_result,omitempty"`
	IgAResult       string `json:"iga_result,omitempty"`
	PanResult       string `json:"pan_result,omitempty"`
	IgMIgGResult    string `json:"igm_igg_result,omitempty"`
	IgMIgAResult    string `json:"igm_iga_result,omitempty"`
	Control         string `json:"control,omitempty"`
	IgMTiter        int    `json:"igm_titer,omitempty"`
	IgGTiter        int    `json:"igg_titer,omitempty"`
	PanTiter        int    `json:"pan_titer,omitempty"`
	IgMTruth        string `json:"igm_truth,omitempty"`
	IgGTruth        string `json:"igg_truth,omitempty"`
	AntibodyTruth   string `json:"antibody_truth,omitempty"`
	IgMAgree        string `json:"igm_agree,omitempty"`
	IgGAgree        string `json:"igg_agree,omitempty"`
	IgAAgree        string `json:"iga_agree,omitempty"`
	PanAgree        string `json:"pan_agree,omitempty"`
	IgMIgGAgree     string `json:"igm_igg_agree,omitempty"`
	IgMIgAAgree     string `json:"igm_iga_agree,omitempty"`
	AntibodyAgree   string `json:"antibody_agree,omitempty"`
}

type Covid19SerologyOptions struct {
	QueryParameters
	Manufacturer    *string `url:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Device          *string `url:"device,omitempty" json:"device,omitempty"`
	DatePerformed   *string `url:"date_performed,omitempty" json:"date_performed,omitempty"`
	EvaluationID    *string `url:"evaluation_id,omitempty" json:"evaluation_id,omitempty"`
	LotNumber       *string `url:"lot_number,omitempty" json:"lot_number,omitempty"`
	Panel           *string `url:"panel,omitempty" json:"panel,omitempty"`
	SampleNo        *int    `url:"sample_no,omitempty" json:"sample_no,omitempty"`
	SampleID        *string `url:"sample_id,omitempty" json:"sample_id,omitempty"`
	Type            *string `url:"type,omitempty" json:"type,omitempty"`
	Group           *string `url:"group,omitempty" json:"group,omitempty"`
	DaysFromSymptom *string `url:"days_from_symptom,omitempty" json:"days_from_symptom,omitempty"`
	IgMResult       *string `url:"igm_result,omitempty" json:"igm_result,omitempty"`
	IgGResult       *string `url:"igg_result,omitempty" json:"igg_result,omitempty"`
	IgAResult       *string `url:"iga_result,omitempty" json:"iga_result,omitempty"`
	PanResult       *string `url:"pan_result,omitempty" json:"pan_result,omitempty"`
	IgMIgGResult    *string `url:"igm_igg_result,omitempty" json:"igm_igg_result,omitempty"`
	IgMIgAResult    *string `url:"igm_iga_result,omitempty" json:"igm_iga_result,omitempty"`
	Control         *string `url:"control,omitempty" json:"control,omitempty"`
	IgMTiter        *int    `url:"igm_titer,omitempty" json:"igm_titer,omitempty"`
	IgGTiter        *int    `url:"igg_titer,omitempty" json:"igg_titer,omitempty"`
	PanTiter        *int    `url:"pan_titer,omitempty" json:"pan_titer,omitempty"`
	IgMTruth        *string `url:"igm_truth,omitempty" json:"igm_truth,omitempty"`
	IgGTruth        *string `url:"igg_truth,omitempty" json:"igg_truth,omitempty"`
	AntibodyTruth   *string `url:"antibody_truth,omitempty" json:"antibody_truth,omitempty"`
	IgMAgree        *string `url:"igm_agree,omitempty" json:"igm_agree,omitempty"`
	IgGAgree        *string `url:"igg_agree,omitempty" json:"igg_agree,omitempty"`
	IgAAgree        *string `url:"iga_agree,omitempty" json:"iga_agree,omitempty"`
	PanAgree        *string `url:"pan_agree,omitempty" json:"pan_agree,omitempty"`
	IgMIgGAgree     *string `url:"igm_igg_agree,omitempty" json:"igm_igg_agree,omitempty"`
	IgMIgAAgree     *string `url:"igm_iga_agree,omitempty" json:"igm_iga_agree,omitempty"`
	AntibodyAgree   *string `url:"antibody_agree,omitempty" json:"antibody_agree,omitempty"`
}

type Covid19SerologyResponse struct {
	Results []*Covid19Serology `json:"results,omitempty"`
	Meta    *Meta              `json:"meta,omitempty"`
}

func (s *Covid19SerologyService) GetCovid19SerologyResponse(opt *Covid19SerologyOptions) (Covid19SerologyResponse, *Response, error) {
	var result Covid19SerologyResponse
	u := devicePath + covid19SerologyRoute

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
