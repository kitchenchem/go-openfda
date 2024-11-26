package fda

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
