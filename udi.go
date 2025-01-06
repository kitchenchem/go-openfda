package fda

import "net/http"

// https://open.fda.gov/apis/device/udi/

type UdiService struct {
	client *Client
}

type Udi struct {
	BrandName                     string                 `json:"brand_name"`
	CatalogNumber                 string                 `json:"catalog_number"`
	CommercialDistributionEndDate string                 `json:"commercial_distribution_end_date"`
	CommercialDistributionStatus  string                 `json:"commercial_distribution_status"`
	CompanyName                   string                 `json:"company_name"`
	CustomerContacts              []*CustomerContact     `json:"customer_contacts"`
	DeviceCountInBasePackage      int                    `json:"device_count_in_base_package"`
	DeviceDescription             string                 `json:"device_description"`
	DeviceSizes                   []*DeviceSize          `json:"device_sizes"`
	GMDNTerms                     []*GMDNTerm            `json:"gmdn_terms"`
	HasDonationIDNumber           bool                   `json:"has_donation_id_number"`
	HasExpirationDate             bool                   `json:"has_expiration_date"`
	HasLotOrBatchNumber           bool                   `json:"has_lot_or_batch_number"`
	HasManufacturingDate          bool                   `json:"has_manufacturing_date"`
	HasSerialNumber               bool                   `json:"has_serial_number"`
	Identifiers                   []*Identifier          `json:"identifiers"`
	IsCombinationProduct          bool                   `json:"is_combination_product"`
	IsDirectMarkingExempt         bool                   `json:"is_direct_marking_exempt"`
	IsHCTP                        bool                   `json:"is_hct_p"`
	IsKit                         bool                   `json:"is_kit"`
	IsLabeledAsNoNRL              bool                   `json:"is_labeled_as_no_nrl"`
	IsLabeledAsNRL                bool                   `json:"is_labeled_as_nrl"`
	IsOTC                         bool                   `json:"is_otc"`
	IsPMExempt                    bool                   `json:"is_pm_exempt"`
	IsRX                          bool                   `json:"is_rx"`
	IsSingleUse                   bool                   `json:"is_single_use"`
	LabelerDUNSNumber             string                 `json:"labeler_duns_number"`
	MRISafety                     string                 `json:"mri_safety"`
	PremarketSubmissions          []*PremarketSubmission `json:"premarket_submissions"`
	ProductCodes                  []*ProductCode         `json:"product_codes"`
	PublishDate                   string                 `json:"publish_date"`
	PublicVersionDate             string                 `json:"public_version_date"`
	PublicVersionNumber           string                 `json:"public_version_number"`
	PublicVersionStatus           string                 `json:"public_version_status"`
	RecordKey                     string                 `json:"record_key"`
	RecordStatus                  string                 `json:"record_status"`
	Sterilization                 *Sterilization         `json:"sterilization"`
	Storage                       []*Storage             `json:"storage"`
	VersionOrModelNumber          string                 `json:"version_or_model_number"`
}

type CustomerContact struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type DeviceSize struct {
	Text  string `json:"text"`
	Type  string `json:"type"`
	Unit  string `json:"unit"`
	Value string `json:"value"`
}

type GMDNTerm struct {
	Code        string `json:"code"`
	Definition  string `json:"definition"`
	Name        string `json:"name"`
	Implantable bool   `json:"implantable"`
	CodeStatus  string `json:"code_status"`
}

type Identifier struct {
	ID                     string `json:"id"`
	IssuingAgency          string `json:"issuing_agency"`
	PackageDiscontinueDate string `json:"package_discontinue_date"`
	PackageStatus          string `json:"package_status"`
	PackageType            string `json:"package_type"`
	QuantityPerPackage     int    `json:"quantity_per_package"`
	Type                   string `json:"type"`
	UnitOfUseID            string `json:"unit_of_use_id"`
}

type PremarketSubmission struct {
	SubmissionNumber string `json:"submission_number"`
	SupplementNumber string `json:"supplement_number"`
	SubmissionType   string `json:"submission_type"`
}

type ProductCode struct {
	Code    string   `json:"code"`
	Name    string   `json:"name"`
	OpenFDA *OpenFDA `json:"openfda"`
}

type Sterilization struct {
	IsSterile               bool   `json:"is_sterile"`
	IsSterilizationPriorUse bool   `json:"is_sterilization_prior_use"`
	SterilizationMethods    string `json:"sterilization_methods"`
}

type Storage struct {
	High struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"high"`
	Low struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"low"`
	SpecialConditions string `json:"special_conditions"`
	Type              string `json:"type"`
}

type UdiOptions struct {
	QueryParameters
	BrandName                     *string                 `url:"brand_name" json:"brand_name"`
	CatalogNumber                 *string                 `url:"catalog_number" json:"catalog_number"`
	CommercialDistributionEndDate *string                 `url:"commercial_distribution_end_date" json:"commercial_distribution_end_date"`
	CommercialDistributionStatus  *string                 `url:"commercial_distribution_status" json:"commercial_distribution_status"`
	CompanyName                   *string                 `url:"company_name" json:"company_name"`
	CustomerContacts              *[]*CustomerContact     `url:"customer_contacts" json:"customer_contacts"`
	DeviceCountInBasePackage      *int                    `url:"device_count_in_base_package" json:"device_count_in_base_package"`
	DeviceDescription             *string                 `url:"device_description" json:"device_description"`
	DeviceSizes                   *[]*DeviceSize          `url:"device_sizes" json:"device_sizes"`
	GMDNTerms                     *[]*GMDNTerm            `url:"gmdn_terms" json:"gmdn_terms"`
	HasDonationIDNumber           *bool                   `url:"has_donation_id_number" json:"has_donation_id_number"`
	HasExpirationDate             *bool                   `url:"has_expiration_date" json:"has_expiration_date"`
	HasLotOrBatchNumber           *bool                   `url:"has_lot_or_batch_number" json:"has_lot_or_batch_number"`
	HasManufacturingDate          *bool                   `url:"has_manufacturing_date" json:"has_manufacturing_date"`
	HasSerialNumber               *bool                   `url:"has_serial_number" json:"has_serial_number"`
	Identifiers                   *[]*Identifier          `url:"identifiers" json:"identifiers"`
	IsCombinationProduct          *bool                   `url:"is_combination_product" json:"is_combination_product"`
	IsDirectMarkingExempt         *bool                   `url:"is_direct_marking_exempt" json:"is_direct_marking_exempt"`
	IsHCTP                        *bool                   `url:"is_hct_p" json:"is_hct_p"`
	IsKit                         *bool                   `url:"is_kit" json:"is_kit"`
	IsLabeledAsNoNRL              *bool                   `url:"is_labeled_as_no_nrl" json:"is_labeled_as_no_nrl"`
	IsLabeledAsNRL                *bool                   `url:"is_labeled_as_nrl" json:"is_labeled_as_nrl"`
	IsOTC                         *bool                   `url:"is_otc" json:"is_otc"`
	IsPMExempt                    *bool                   `url:"is_pm_exempt" json:"is_pm_exempt"`
	IsRX                          *bool                   `url:"is_rx" json:"is_rx"`
	IsSingleUse                   *bool                   `url:"is_single_use" json:"is_single_use"`
	LabelerDUNSNumber             *string                 `url:"labeler_duns_number" json:"labeler_duns_number"`
	MRISafety                     *string                 `url:"mri_safety" json:"mri_safety"`
	PremarketSubmissions          *[]*PremarketSubmission `url:"premarket_submissions" json:"premarket_submissions"`
	ProductCodes                  *[]*ProductCode         `url:"product_codes" json:"product_codes"`
	PublishDate                   *string                 `url:"publish_date" json:"publish_date"`
	PublicVersionDate             *string                 `url:"public_version_date" json:"public_version_date"`
	PublicVersionNumber           *string                 `url:"public_version_number" json:"public_version_number"`
	PublicVersionStatus           *string                 `url:"public_version_status" json:"public_version_status"`
	RecordKey                     *string                 `url:"record_key" json:"record_key"`
	RecordStatus                  *string                 `url:"record_status" json:"record_status"`
	Sterilization                 *Sterilization          `url:"sterilization" json:"sterilization"`
	Storage                       *[]*Storage             `url:"storage" json:"storage"`
	VersionOrModelNumber          *string                 `url:"version_or_model_number" json:"version_or_model_number"`
}

type UdiResponse struct {
	Results []*Udi `json:"results,omitempty"`
	Meta    *Meta  `json:"meta,omitempty"`
}

func (s *UdiService) GetUdi(opt *UdiOptions, options ...RequestOptionFunc) (UdiResponse, *Response, error) {
	var result UdiResponse
	u := devicePath + udiRoute
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
