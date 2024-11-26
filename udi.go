package fda

type MedicalDevice struct {
	BrandName                     string                 `json:"brand_name,omitempty"`
	CatalogNumber                 string                 `json:"catalog_number,omitempty"`
	CommercialDistributionEndDate string                 `json:"commercial_distribution_end_date,omitempty"`
	CommercialDistributionStatus  string                 `json:"commercial_distribution_status,omitempty"`
	CompanyName                   string                 `json:"company_name,omitempty"`
	CustomerContacts              []*CustomerContact     `json:"customer_contacts,omitempty"`
	DeviceCountInBasePackage      int                    `json:"device_count_in_base_package,omitempty"`
	DeviceDescription             string                 `json:"device_description,omitempty"`
	DeviceSizes                   []*DeviceSize          `json:"device_sizes,omitempty"`
	GMDNTerms                     []*GMDNTerm            `json:"gmdn_terms,omitempty"`
	HasDonationIDNumber           bool                   `json:"has_donation_id_number,omitempty"`
	HasExpirationDate             bool                   `json:"has_expiration_date,omitempty"`
	HasLotOrBatchNumber           bool                   `json:"has_lot_or_batch_number,omitempty"`
	HasManufacturingDate          bool                   `json:"has_manufacturing_date,omitempty"`
	HasSerialNumber               bool                   `json:"has_serial_number,omitempty"`
	Identifiers                   []*Identifier          `json:"identifiers,omitempty"`
	IsCombinationProduct          bool                   `json:"is_combination_product,omitempty"`
	IsDirectMarkingExempt         bool                   `json:"is_direct_marking_exempt,omitempty"`
	IsHCTP                        bool                   `json:"is_hct_p,omitempty"`
	IsKit                         bool                   `json:"is_kit,omitempty"`
	IsLabeledAsNoNRL              bool                   `json:"is_labeled_as_no_nrl,omitempty"`
	IsLabeledAsNRL                bool                   `json:"is_labeled_as_nrl,omitempty"`
	IsOTC                         bool                   `json:"is_otc,omitempty"`
	IsPMExempt                    bool                   `json:"is_pm_exempt,omitempty"`
	IsRX                          bool                   `json:"is_rx,omitempty"`
	IsSingleUse                   bool                   `json:"is_single_use,omitempty"`
	LabelerDUNSNumber             string                 `json:"labeler_duns_number,omitempty"`
	MRISafety                     string                 `json:"mri_safety,omitempty"`
	PremarketSubmissions          []*PremarketSubmission `json:"premarket_submissions,omitempty"`
	ProductCodes                  []*ProductCode         `json:"product_codes,omitempty"`
	PublishDate                   string                 `json:"publish_date,omitempty"`
	PublicVersionDate             string                 `json:"public_version_date,omitempty"`
	PublicVersionNumber           string                 `json:"public_version_number,omitempty"`
	PublicVersionStatus           string                 `json:"public_version_status,omitempty"`
	RecordKey                     string                 `json:"record_key,omitempty"`
	RecordStatus                  string                 `json:"record_status,omitempty"`
	Sterilization                 *Sterilization         `json:"sterilization,omitempty"`
	Storage                       []*Storage             `json:"storage,omitempty"`
	VersionOrModelNumber          string                 `json:"version_or_model_number,omitempty"`
}

type CustomerContact struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type DeviceSize struct {
	Text  string `json:"text,omitempty"`
	Type  string `json:"type,omitempty"`
	Unit  string `json:"unit,omitempty"`
	Value string `json:"value,omitempty"`
}

type GMDNTerm struct {
	Code        string `json:"code,omitempty"`
	Definition  string `json:"definition,omitempty"`
	Name        string `json:"name,omitempty"`
	Implantable bool   `json:"implantable,omitempty"`
	CodeStatus  string `json:"code_status,omitempty"`
}

type Identifier struct {
	ID                     string `json:"id,omitempty"`
	IssuingAgency          string `json:"issuing_agency,omitempty"`
	PackageDiscontinueDate string `json:"package_discontinue_date,omitempty"`
	PackageStatus          string `json:"package_status,omitempty"`
	PackageType            string `json:"package_type,omitempty"`
	QuantityPerPackage     int    `json:"quantity_per_package,omitempty"`
	Type                   string `json:"type,omitempty"`
	UnitOfUseID            string `json:"unit_of_use_id,omitempty"`
}

type PremarketSubmission struct {
	SubmissionNumber string `json:"submission_number,omitempty"`
	SupplementNumber string `json:"supplement_number,omitempty"`
	SubmissionType   string `json:"submission_type,omitempty"`
}

type ProductCode struct {
	Code    string   `json:"code,omitempty"`
	Name    string   `json:"name,omitempty"`
	OpenFDA *OpenFDA `json:"openfda,omitempty"`
}

type Sterilization struct {
	IsSterile               bool   `json:"is_sterile,omitempty"`
	IsSterilizationPriorUse bool   `json:"is_sterilization_prior_use,omitempty"`
	SterilizationMethods    string `json:"sterilization_methods,omitempty"`
}

type Storage struct {
	High struct {
		Unit  string `json:"unit,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"high,omitempty"`
	Low struct {
		Unit  string `json:"unit,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"low,omitempty"`
	SpecialConditions string `json:"special_conditions,omitempty"`
	Type              string `json:"type,omitempty"`
}
