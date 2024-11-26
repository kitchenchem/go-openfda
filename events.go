package fda

import "time"

type DeviceEventReport struct {
	AdverseEventFlag         string     `json:"adverse_event_flag,omitempty"`
	DateFacilityAware        *time.Time `json:"date_facility_aware,omitempty"`
	DateManufacturerReceived *time.Time `json:"date_manufacturer_received,omitempty"`
	DateOfEvent              *time.Time `json:"date_of_event,omitempty"`
	DateReceived             *time.Time `json:"date_received,omitempty"`
	DateReport               *time.Time `json:"date_report,omitempty"`
	DateReportToFDA          *time.Time `json:"date_report_to_fda,omitempty"`
	DateReportToManufacturer *time.Time `json:"date_report_to_manufacturer,omitempty"`
	Devices                  []*Device  `json:"device,omitempty"`
	DeviceDateOfManufacturer *time.Time `json:"device_date_of_manufacturer,omitempty"`
	DistributorAddress1      string     `json:"distributor_address_1,omitempty"`
	DistributorAddress2      string     `json:"distributor_address_2,omitempty"`
	DistributorCity          string     `json:"distributor_city,omitempty"`
	DistributorName          string     `json:"distributor_name,omitempty"`
	DistributorState         string     `json:"distributor_state,omitempty"`
	DistributorZipCode       string     `json:"distributor_zip_code,omitempty"`
	DistributorZipCodeExt    string     `json:"distributor_zip_code_ext,omitempty"`
	EventKey                 string     `json:"event_key,omitempty"`
	EventLocation            string     `json:"event_location,omitempty"`
	EventType                string     `json:"event_type,omitempty"`
	ExpirationDateOfDevice   *time.Time `json:"expiration_date_of_device,omitempty"`
	HealthProfessional       string     `json:"health_professional,omitempty"`
	InitialReportToFDA       string     `json:"initial_report_to_fda,omitempty"`
	ManufacturerAddress1     string     `json:"manufacturer_address_1,omitempty"`
	ManufacturerAddress2     string     `json:"manufacturer_address_2,omitempty"`
	ManufacturerCity         string     `json:"manufacturer_city,omitempty"`
	ManufacturerName         string     `json:"manufacturer_name,omitempty"`
	ManufacturerPostalCode   string     `json:"manufacturer_postal_code,omitempty"`
	ManufacturerState        string     `json:"manufacturer_state,omitempty"`
	ManufacturerZipCode      string     `json:"manufacturer_zip_code,omitempty"`
	ManufacturerZipCodeExt   string     `json:"manufacturer_zip_code_ext,omitempty"`
	MDRReportKey             string     `json:"mdr_report_key,omitempty"`
	MDRText                  []*MDRText `json:"mdr_text,omitempty"`
	NumberDevicesInEvent     string     `json:"number_devices_in_event,omitempty"`
	NumberPatientsInEvent    string     `json:"number_patients_in_event,omitempty"`
	Patients                 []*Patient `json:"patient,omitempty"`
	PreviousUseCode          string     `json:"previous_use_code,omitempty"`
	ProductProblems          []string   `json:"product_problems,omitempty"`
	ProductProblemFlag       string     `json:"product_problem_flag,omitempty"`
	RemedialAction           []string   `json:"remedial_action,omitempty"`
	RemovalCorrectionNumber  string     `json:"removal_correction_number,omitempty"`
	ReportDate               *time.Time `json:"report_date,omitempty"`
	ReportNumber             string     `json:"report_number,omitempty"`
	ReportSourceCode         string     `json:"report_source_code,omitempty"`
	ReportToFDA              string     `json:"report_to_fda,omitempty"`
	ReportToManufacturer     string     `json:"report_to_manufacturer,omitempty"`
	ReporterOccupationCode   string     `json:"reporter_occupation_code,omitempty"`
	ReprocessedAndReusedFlag string     `json:"reprocessed_and_reused_flag,omitempty"`
	SingleUseFlag            string     `json:"single_use_flag,omitempty"`
	SourceType               []string   `json:"source_type,omitempty"`
	TypeOfReport             []string   `json:"type_of_report,omitempty"`
}

type Device struct {
	BrandName                     string     `json:"brand_name,omitempty"`
	CatalogNumber                 string     `json:"catalog_number,omitempty"`
	DateReceived                  *time.Time `json:"date_received,omitempty"`
	DateRemovedFlag               string     `json:"date_removed_flag,omitempty"`
	DateReturnedToManufacturer    *time.Time `json:"date_returned_to_manufacturer,omitempty"`
	DeviceAgeText                 string     `json:"device_age_text,omitempty"`
	DeviceAvailability            string     `json:"device_availability,omitempty"`
	DeviceEvaluatedByManufacturer string     `json:"device_evaluated_by_manufacturer,omitempty"`
	DeviceEventKey                string     `json:"device_event_key,omitempty"`
	DeviceOperator                string     `json:"device_operator,omitempty"`
	DeviceReportProductCode       string     `json:"device_report_product_code,omitempty"`
	DeviceSequenceNumber          string     `json:"device_sequence_number,omitempty"`
	ExpirationDateOfDevice        *time.Time `json:"expiration_date_of_device,omitempty"`
	GenericName                   string     `json:"generic_name,omitempty"`
	UDIDI                         string     `json:"udi_di,omitempty"`
	ImplantFlag                   string     `json:"implant_flag,omitempty"`
	LotNumber                     string     `json:"lot_number,omitempty"`
	ManufacturerDAddress1         string     `json:"manufacturer_d_address_1,omitempty"`
	ManufacturerDAddress2         string     `json:"manufacturer_d_address_2,omitempty"`
	ManufacturerDCity             string     `json:"manufacturer_d_city,omitempty"`
	ManufacturerDCountry          string     `json:"manufacturer_d_country,omitempty"`
	ManufacturerDName             string     `json:"manufacturer_d_name,omitempty"`
	ManufacturerDPostalCode       string     `json:"manufacturer_d_postal_code,omitempty"`
	ManufacturerDState            string     `json:"manufacturer_d_state,omitempty"`
	ManufacturerDZipCode          string     `json:"manufacturer_d_zip_code,omitempty"`
	ManufacturerDZipCodeExt       string     `json:"manufacturer_d_zip_code_ext,omitempty"`
	ModelNumber                   string     `json:"model_number,omitempty"`
	OpenFDA                       *OpenFDA   `json:"openfda,omitempty"`
	OtherIDNumber                 string     `json:"other_id_number,omitempty"`
	UDIPublic                     string     `json:"udi_public,omitempty"`
}

type MDRText struct {
	DateReport            *time.Time `json:"date_report,omitempty"`
	MDRTextKey            string     `json:"mdr_text_key,omitempty"`
	PatientSequenceNumber string     `json:"patient_sequence_number,omitempty"`
	Text                  string     `json:"text,omitempty"`
	TextTypeCode          string     `json:"text_type_code,omitempty"`
}

type Patient struct {
	DateReceived            *time.Time `json:"date_received,omitempty"`
	PatientSequenceNumber   string     `json:"patient_sequence_number,omitempty"`
	PatientAge              string     `json:"patient_age,omitempty"`
	PatientSex              string     `json:"patient_sex,omitempty"`
	PatientWeight           string     `json:"patient_weight,omitempty"`
	PatientEthnicity        string     `json:"patient_ethnicity,omitempty"`
	PatientRace             string     `json:"patient_race,omitempty"`
	PatientProblems         []string   `json:"patient_problems,omitempty"`
	SequenceNumberOutcome   []string   `json:"sequence_number_outcome,omitempty"`
	SequenceNumberTreatment []string   `json:"sequence_number_treatment,omitempty"`
}
