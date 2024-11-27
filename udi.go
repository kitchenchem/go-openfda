package fda

// https://open.fda.gov/apis/device/udi/

type UdiService struct {
	client *Client
}

type UDI struct {
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
