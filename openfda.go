package fda

//extends shared structs to work for any endpoints that use them

type OpenFDA struct {
	ApplicationNumber          []string `json:"application_number,omitempty"`
	BrandName                  []string `json:"brand_name,omitempty"`
	DosageForm                 []string `json:"dosage_form,omitempty"`
	DeviceClass                string   `json:"device_class,omitempty"`
	DeviceName                 string   `json:"device_name,omitempty"`
	FEINumber                  []string `json:"fei_number,omitempty"`
	GenericName                []string `json:"generic_name,omitempty"`
	IsOriginalPackager         bool     `json:"is_original_packager,omitempty"`
	KNumber                    []string `json:"k_number,omitempty"`
	ManufacturerName           []string `json:"manufacturer_name,omitempty"`
	MedicalSpecialtyDesc       string   `json:"medical_specialty_description,omitempty"`
	Nui                        []string `json:"nui,omitempty"`
	OriginalPackagerProductNDC []string `json:"original_packager_product_ndc,omitempty"`
	PackageNDC                 []string `json:"package_ndc,omitempty"`
	PharmClassCS               []string `json:"pharm_class_cs,omitempty"`
	PharmClassEPC              []string `json:"pharm_class_epc,omitempty"`
	PharmClassMOA              []string `json:"pharm_class_moa,omitempty"`
	PharmClassPE               []string `json:"pharm_class_pe,omitempty"`
	ProductNDC                 []string `json:"product_ndc,omitempty"`
	ProductType                []string `json:"product_type,omitempty"`
	RegistrationNumber         []string `json:"registration_number,omitempty"`
	RegulationNumber           string   `json:"regulation_number,omitempty"`
	Route                      []string `json:"route,omitempty"`
	RxCUI                      []string `json:"rxcui,omitempty"`
	RxString                   []string `json:"rxstring,omitempty"`
	RxTTY                      []string `json:"rxtty,omitempty"`
	SPLID                      []string `json:"spl_id,omitempty"`
	SPLSetID                   []string `json:"spl_set_id,omitempty"`
	SubstanceName              []string `json:"substance_name,omitempty"`
	UNII                       []string `json:"unii,omitempty"`
	UPC                        []string `json:"upc,omitempty"`
}

type Meta struct {
	Disclaimer  string     `json:"disclaimer,omitempty"`
	License     string     `json:"license,omitempty"`
	LastUpdated string     `json:"last_updated,omitempty"`
	Results     *MetaStats `json:"results,omitempty"`
}

type Results struct {
	Skip  int64 `json:"skip,omitempty"`
	Limit int64 `json:"limit,omitempty"`
	Total int64 `json:"total,omitempty"`
}

type MetaStats struct {
	Skip  int64 `json:"skip"`
	Limit int64 `json:"limit"`
	Total int64 `json:"total"`
}
