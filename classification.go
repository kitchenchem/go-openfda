package fda

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
