package fda

import (
	"net/http"
	"time"
)

type EventService struct {
	client *Client
}

type Event struct {
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
	BrandName                  string     `json:"brand_name,omitempty"`
	CatalogNumber              string     `json:"catalog_number,omitempty"`
	DateReceived               *time.Time `json:"date_received,omitempty"`
	DateRemovedFlag            string     `json:"date_removed_flag,omitempty"`
	DateReturnedToManufacturer *time.Time `json:"date_returned_to_manufacturer,omitempty"`
	DeviceAgeText              string     `json:"device_age_text,omitempty"`
	DeviceAvailability         string     `json:"device_availability,omitempty"`
	EvaluatedByManufacturer    string     `json:"device_evaluated_by_manufacturer,omitempty"`
	EventKey                   string     `json:"device_event_key,omitempty"`
	DeviceOperator             string     `json:"device_operator,omitempty"`
	DeviceReportProductCode    string     `json:"device_report_product_code,omitempty"`
	DeviceSequenceNumber       string     `json:"device_sequence_number,omitempty"`
	ExpirationDateOfDevice     *time.Time `json:"expiration_date_of_device,omitempty"`
	GenericName                string     `json:"generic_name,omitempty"`
	UDIDI                      string     `json:"udi_di,omitempty"`
	ImplantFlag                string     `json:"implant_flag,omitempty"`
	LotNumber                  string     `json:"lot_number,omitempty"`
	ManufacturerDAddress1      string     `json:"manufacturer_d_address_1,omitempty"`
	ManufacturerDAddress2      string     `json:"manufacturer_d_address_2,omitempty"`
	ManufacturerDCity          string     `json:"manufacturer_d_city,omitempty"`
	ManufacturerDCountry       string     `json:"manufacturer_d_country,omitempty"`
	ManufacturerDName          string     `json:"manufacturer_d_name,omitempty"`
	ManufacturerDPostalCode    string     `json:"manufacturer_d_postal_code,omitempty"`
	ManufacturerDState         string     `json:"manufacturer_d_state,omitempty"`
	ManufacturerDZipCode       string     `json:"manufacturer_d_zip_code,omitempty"`
	ManufacturerDZipCodeExt    string     `json:"manufacturer_d_zip_code_ext,omitempty"`
	ModelNumber                string     `json:"model_number,omitempty"`
	OpenFDA                    *OpenFDA   `json:"openfda,omitempty"`
	OtherIDNumber              string     `json:"other_id_number,omitempty"`
	UDIPublic                  string     `json:"udi_public,omitempty"`
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
type EventOptions struct {
	QueryParameters
	AdverseEventFlag         *string     `url:"adverse_event_flag" json:"adverse_event_flag,omitempty"`
	DateFacilityAware        *time.Time  `url:"date_facility_aware" json:"date_facility_aware,omitempty"`
	DateManufacturerReceived *time.Time  `url:"date_manufacturer_received" json:"date_manufacturer_received,omitempty"`
	DateOfEvent              *time.Time  `url:"date_of_event" json:"date_of_event,omitempty"`
	DateReceived             *time.Time  `url:"date_received" json:"date_received,omitempty"`
	DateReport               *time.Time  `url:"date_report" json:"date_report,omitempty"`
	DateReportToFDA          *time.Time  `url:"date_report_to_fda" json:"date_report_to_fda,omitempty"`
	DateReportToManufacturer *time.Time  `url:"date_report_to_manufacturer" json:"date_report_to_manufacturer,omitempty"`
	Devices                  *[]*Device  `url:"device" json:"device,omitempty"`
	DeviceDateOfManufacturer *time.Time  `url:"device_date_of_manufacturer" json:"device_date_of_manufacturer,omitempty"`
	DistributorAddress1      *string     `url:"distributor_address_1" json:"distributor_address_1,omitempty"`
	DistributorAddress2      *string     `url:"distributor_address_2" json:"distributor_address_2,omitempty"`
	DistributorCity          *string     `url:"distributor_city" json:"distributor_city,omitempty"`
	DistributorName          *string     `url:"distributor_name" json:"distributor_name,omitempty"`
	DistributorState         *string     `url:"distributor_state" json:"distributor_state,omitempty"`
	DistributorZipCode       *string     `url:"distributor_zip_code" json:"distributor_zip_code,omitempty"`
	DistributorZipCodeExt    *string     `url:"distributor_zip_code_ext" json:"distributor_zip_code_ext,omitempty"`
	EventKey                 *string     `url:"event_key" json:"event_key,omitempty"`
	EventLocation            *string     `url:"event_location" json:"event_location,omitempty"`
	EventType                *string     `url:"event_type" json:"event_type,omitempty"`
	ExpirationDateOfDevice   *time.Time  `url:"expiration_date_of_device" json:"expiration_date_of_device,omitempty"`
	HealthProfessional       *string     `url:"health_professional" json:"health_professional,omitempty"`
	InitialReportToFDA       *string     `url:"initial_report_to_fda" json:"initial_report_to_fda,omitempty"`
	ManufacturerAddress1     *string     `url:"manufacturer_address_1" json:"manufacturer_address_1,omitempty"`
	ManufacturerAddress2     *string     `url:"manufacturer_address_2" json:"manufacturer_address_2,omitempty"`
	ManufacturerCity         *string     `url:"manufacturer_city" json:"manufacturer_city,omitempty"`
	ManufacturerName         *string     `url:"manufacturer_name" json:"manufacturer_name,omitempty"`
	ManufacturerPostalCode   *string     `url:"manufacturer_postal_code" json:"manufacturer_postal_code,omitempty"`
	ManufacturerState        *string     `url:"manufacturer_state" json:"manufacturer_state,omitempty"`
	ManufacturerZipCode      *string     `url:"manufacturer_zip_code" json:"manufacturer_zip_code,omitempty"`
	ManufacturerZipCodeExt   *string     `url:"manufacturer_zip_code_ext" json:"manufacturer_zip_code_ext,omitempty"`
	MDRReportKey             *string     `url:"mdr_report_key" json:"mdr_report_key,omitempty"`
	MDRText                  *[]*MDRText `url:"mdr_text" json:"mdr_text,omitempty"`
	NumberDevicesInEvent     *string     `url:"number_devices_in_event" json:"number_devices_in_event,omitempty"`
	NumberPatientsInEvent    *string     `url:"number_patients_in_event" json:"number_patients_in_event,omitempty"`
	Patients                 *[]*Patient `url:"patient" json:"patient,omitempty"`
	PreviousUseCode          *string     `url:"previous_use_code" json:"previous_use_code,omitempty"`
	ProductProblems          *[]string   `url:"product_problems" json:"product_problems,omitempty"`
	ProductProblemFlag       *string     `url:"product_problem_flag" json:"product_problem_flag,omitempty"`
	RemedialAction           *[]string   `url:"remedial_action" json:"remedial_action,omitempty"`
	RemovalCorrectionNumber  *string     `url:"removal_correction_number" json:"removal_correction_number,omitempty"`
	ReportDate               *time.Time  `url:"report_date" json:"report_date,omitempty"`
	ReportNumber             *string     `url:"report_number" json:"report_number,omitempty"`
	ReportSourceCode         *string     `url:"report_source_code" json:"report_source_code,omitempty"`
	ReportToFDA              *string     `url:"report_to_fda" json:"report_to_fda,omitempty"`
	ReportToManufacturer     *string     `url:"report_to_manufacturer" json:"report_to_manufacturer,omitempty"`
	ReporterOccupationCode   *string     `url:"reporter_occupation_code" json:"reporter_occupation_code,omitempty"`
	ReprocessedAndReusedFlag *string     `url:"reprocessed_and_reused_flag" json:"reprocessed_and_reused_flag,omitempty"`
	SingleUseFlag            *string     `url:"single_use_flag" json:"single_use_flag,omitempty"`
	SourceType               *[]string   `url:"source_type" json:"source_type,omitempty"`
	TypeOfReport             *[]string   `url:"type_of_report" json:"type_of_report,omitempty"`
}
type EventResponse struct {
	Results []*Event `json:"results,omitempty"`
	Meta    *Meta    `json:"meta,omitempty"`
}

func (s *EventService) GetEvent(opt *EventOptions, options ...RequestOptionFunc) (EventResponse, *Response, error) {
	var result EventResponse
	u := devicePath + eventRoute

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
