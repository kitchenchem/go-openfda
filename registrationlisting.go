package fda

type RegistrationListing struct {
	EstablishmentType []string      `json:"establishment_type,omitempty"`
	KNumber           string        `json:"k_number,omitempty"`
	PMANumber         string        `json:"pma_number,omitempty"`
	Products          []*Product    `json:"products,omitempty"`
	ProprietaryName   []string      `json:"proprietary_name,omitempty"`
	Registration      *Registration `json:"registration,omitempty"`
}

type Product struct {
	CreatedDate         string   `json:"created_date,omitempty"`
	Exempt              string   `json:"exempt,omitempty"`
	OpenFDA             *OpenFDA `json:"openfda,omitempty"`
	OwnerOperatorNumber string   `json:"owner_operator_number,omitempty"`
	ProductCode         string   `json:"product_code,omitempty"`
}

type Registration struct {
	AddressLine1        string         `json:"address_line_1,omitempty"`
	AddressLine2        string         `json:"address_line_2,omitempty"`
	City                string         `json:"city,omitempty"`
	FEINumber           string         `json:"fei_number,omitempty"`
	InitialImporterFlag string         `json:"initial_importer_flag,omitempty"`
	ISOCountryCode      string         `json:"iso_country_code,omitempty"`
	Name                string         `json:"name,omitempty"`
	OwnerOperator       *OwnerOperator `json:"owner_operator,omitempty"`
	PostalCode          string         `json:"postal_code,omitempty"`
	RegExpiryDateYear   string         `json:"reg_expiry_date_year,omitempty"`
	RegistrationNumber  string         `json:"registration_number,omitempty"`
	StateCode           string         `json:"state_code,omitempty"`
	StatusCode          string         `json:"status_code,omitempty"`
	USAgent             *USAgent       `json:"us_agent,omitempty"`
	ZipCode             string         `json:"zip_code,omitempty"`
}

type OwnerOperator struct {
	ContactAddress        *ContactAddress  `json:"contact_address,omitempty"`
	FirmName              string           `json:"firm_name,omitempty"`
	OfficialCorrespondent *OfficialContact `json:"official_correspondent,omitempty"`
	OwnerOperatorNumber   string           `json:"owner_operator_number,omitempty"`
}

type ContactAddress struct {
	Address1       string `json:"address_1,omitempty"`
	Address2       string `json:"address_2,omitempty"`
	City           string `json:"city,omitempty"`
	ISOCountryCode string `json:"iso_country_code,omitempty"`
	PostalCode     string `json:"postal_code,omitempty"`
	StateCode      string `json:"state_code,omitempty"`
	StateProvince  string `json:"state_province,omitempty"`
}

type OfficialContact struct {
	FirstName             string `json:"first_name,omitempty"`
	LastName              string `json:"last_name,omitempty"`
	MiddleInitial         string `json:"middle_initial,omitempty"`
	PhoneNumber           string `json:"phone_number,omitempty"`
	SubaccountCompanyName string `json:"subaccount_company_name,omitempty"`
}

type USAgent struct {
	AddressLine1     string `json:"address_line_1,omitempty"`
	AddressLine2     string `json:"address_line_2,omitempty"`
	BusinessName     string `json:"business_name,omitempty"`
	BusPhoneAreaCode string `json:"bus_phone_area_code,omitempty"`
	BusPhoneExtn     string `json:"bus_phone_extn,omitempty"`
	BusPhoneNum      string `json:"bus_phone_num,omitempty"`
	City             string `json:"city,omitempty"`
	EmailAddress     string `json:"email_address,omitempty"`
	FaxAreaCode      string `json:"fax_area_code,omitempty"`
	FaxNum           string `json:"fax_num,omitempty"`
	ISOCountryCode   string `json:"iso_country_code,omitempty"`
	Name             string `json:"name,omitempty"`
	PostalCode       string `json:"postal_code,omitempty"`
	StateCode        string `json:"state_code,omitempty"`
	ZipCode          string `json:"zip_code,omitempty"`
}

type RegistrationListingOptions struct {
	EstablishmentType *[]string     `url:"establishment_type,omitempty" json:"establishment_type,omitempty"`
	KNumber           *string       `url:"k_number,omitempty" json:"k_number,omitempty"`
	PMANumber         *string       `url:"pma_number,omitempty" json:"pma_number,omitempty"`
	Products          *[]*Product   `url:"products,omitempty" json:"products,omitempty"`
	ProprietaryName   *[]string     `url:"proprietary_name,omitempty" json:"proprietary_name,omitempty"`
	Registration      *Registration `url:"registration,omitempty" json:"registration,omitempty"`
}
