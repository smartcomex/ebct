package ebct

const packagesRequest = DefaultEndPoint + Version + "packages"

type PackageRequestList struct {
	PackageList *[]PackageRequest `json:"packageList,omitempty"` // List of Packages
}

type PackageRequest struct {
	CustomerControlCode     *string `json:"customerControlCode,omitempty"`     // It can be alphanumeric. This number is a control code established by the customer. This number does not have a specific elaboration rule and can be any alphanumeric value. It can be used to link the order number or any other necessary controls.
	SenderName              *string `json:"senderName,omitempty"`              // Inform the sender’s name, which should be the same as it is printed on the Package label.
	SenderAddress           *string `json:"senderAddress,omitempty"`           // This field must contain the name of the street where the sender is located.
	SenderAddressNumber     *string `json:"senderAddressNumber,omitempty"`     // Inform the sender’s house number.
	SenderAddressComplement *string `json:"senderAddressComplement,omitempty"` // Inform the complement of the sender’s address.
	SenderZipCode           *string `json:"senderZipCode,omitempty"`           // Inform the sender’s zip code.
	SenderCityName          *string `json:"senderCityName,omitempty"`          // Inform the sender’s city.
	SenderState             *string `json:"senderState,omitempty"`             // Inform the state/province/department abbreviation, from which the sender’s city belongs to.
	SenderCountryCode       *string `json:"senderCountryCode,omitempty"`       // Inform the sender’s country abbreviation. The abbreviation should meet the ISO 3166-2. This tag is used to publish the tracking event “posted” related to the country of origin.
	SenderEmail             *string `json:"senderEmail,omitempty"`             // Inform the sender’s e-mail.
	SenderWebsite           *string `json:"senderWebsite,omitempty"`           // Inform the seller web address. This information is Mandatory according with Customs request.
	RecipientName           *string `json:"recipientName,omitempty"`           // Inform the receiver’s name, which must be the same as it is printed on the package label. The Brazilian Customs does not allow general or first names, such as: recipientName: “Maria”.
	RecipientDocumentType   *string `json:"recipientDocumentType,omitempty"`   /* Inform the type of document of the recipient. This field must receive one of the values in bold below:
	CPF - “Registration of Individuals” (for individuals)
	CNPJ - “National Register of Legal Entities” (for companies)
	PASSPORT - For Individuals who choose to use this document.
	Brazilian Post Office are required by Brazilian customs to provide the importer's tax identification number to recipients. */
	RecipientDocumentNumber    *string  `json:"recipientDocumentNumber,omitempty"`    // Recipient's document number. This field must receive the number of the document whose type (CPF, CPNJ or PASSPORT) was informed in the recipientDocumentType field.
	RecipientAddress           *string  `json:"recipientAddress,omitempty"`           // Inform the receiver’s street address, which should be the same as printed on the Package label. This field is for the street name, and not the whole address.
	RecipientAddressNumber     *string  `json:"recipientAddressNumber,omitempty"`     // Inform the receiver’s house number, which should be the same as printed on the package label. This field is essential for delivery.
	RecipientAddressComplement *string  `json:"recipientAddressComplement,omitempty"` // Inform the complement of the receiver’s address, which should be the same as printed on the package label. (i.e. apartment, house, block).
	RecipientCityName          *string  `json:"recipientCityName,omitempty"`          // Inform the receiver’s city, which should be the same as printed on the package label.
	RecipientState             *string  `json:"recipientState,omitempty"`             // Inform the state abbreviation of the receiver, which should be the same as printed on the package label.
	RecipientZipCode           *string  `json:"recipientZipCode,omitempty"`           // Inform the receiver’s zip code, which should be the same as printed on the package label. Essential for delivery. It should be valid; otherwise, it slows customs clearance process.
	RecipientEmail             *string  `json:"recipientEmail,omitempty"`             // Inform the receiver’s e-mail, which should be the same as printed on the package label. If informed, the format should be valid. With this information, the Brazilian Post will be able to be in touch with the receiver in order to ease the customs clearance.
	RecipientPhoneNumber       *string  `json:"recipientPhoneNumber,omitempty"`       // Enter the recipient's phone number. With this information, the Brazilian Post Office will be able to contact the recipient to facilitate customs clearance. This number consists of 10 (landline phone) or 11 (mobile phone) numbers. Must be informed without the international prefix, as examples below: Landline: 3133334444 Mobile phone: 31999999999
	TotalWeight                *int     `json:"totalWeight,omitempty"`                // Inform the Package weight (including the packaging weight). It should be in grams (g).  It cannot be zero (0). It cannot be greater than 30000g.
	PackagingLength            *float32 `json:"packagingLength,omitempty"`            // Inform the packaging length. It should be in centimeters (cm).  It cannot be smaller than 16cm. It cannot be greater than 100cm.
	PackagingWidth             *float32 `json:"packagingWidth,omitempty"`             // Inform the packaging width. It should be in centimeters (cm). It cannot be smaller than 11cm.It cannot be greater than 100cm.
	PackagingHeight            *float32 `json:"packagingHeight,omitempty"`            // Inform the packaging height. It should be in centimeters (cm). It cannot be smaller than 2cm.It cannot be greater than 100cm.
	DistributionModality       *int     `json:"distributionModality,omitempty"`       // Inform the distribution modality in the Brazilian territory. Should receive the values:  33162 - PACKET STANDARD 33170 - PACKET EXPRESS 33197 - PACKET MINI
	TaxPaymentMethod           *string  `json:"taxPaymentMethod,omitempty"`           // Inform the tax payment method. It should be the code below: “DDU” - payment a posteriori. The receiver pays the taxes, after the Brazilian customs clearance.
	Currency                   *string  `json:"currency,omitempty"`                   // Currency of the transaction accordingly with its origin. “USD” - US Dollars USD is the only option at this time.
	FreightPaidValue           *float32 `json:"freightPaidValue,omitempty"`           // Inform the freight value paid by the receiver before arriving in Brazil. minimum: 0.01 maximum: 999999
	InsurancePaidValue         *float32 `json:"insurancePaidValue,omitempty"`         // Inform the insurance value paid by the receiver.  minimum: 0.01 maximum: 999999
	Items                      []Item   `json:"items,omitempty"`                      // Array of one object.  It should be limited to 20 occurrences.
}

type Item struct {
	HsCode      *string  `json:"hsCode,omitempty"`      // Harmonized System Code (HS) for each of the items in the box. Its nomenclature is an internationally standardized system of names and numbers to classify commercialized products. The Post Office of Brazill are not responsible for the tax classification of products. We can help you just by showing the link on the WCO World Customs Organization website, where you can find the General Rules for the interpretation of the Harmonized System.  http://www.wcoomd.org/en/topics/nomenclature/instrument-and-tools/hs-nomenclature-2017-edition/hs-nomenclature-2017-edition.aspx Please, request to your sales consultant the current HS Code used by Brazilian customs system.
	Description *string  `json:"description,omitempty"` // Inform the product/item description
	Quantity    *int     `json:"quantity,omitempty"`    // Inform the item quantity.  It cannot be zero. minimum: 1 maximum: 9999
	Value       *float32 `json:"value,omitempty"`       // Inform the item unitary value.  It cannot be zero. minimum: 0.01 maximum: 3000 Please check the business rules below.
}

func (n PackageRequest) Error() string {
	panic("implement me")
}

type PackageRequestReturn struct {
	PackageResponseList []PackResponse `json:"packageResponseList"`
	ErrorMessage        *string        `json:"errorMessage,omitempty"`
	ErrorDetails        []ErrorDetail  `json:"errorDetails,omitempty"`
}

type PackResponse struct {
	CustomerControlCode *string `json:"customerControlCode"`
	TrackingNumber      *string `json:"trackingNumber"`
}

func (c *Client) PostPackage(PackRequest *PackageRequestList) (*PackageRequestReturn, error) {
	newPackageList := &PackageRequestReturn{}
	err := c.Post(packagesRequest, PackRequest, nil, newPackageList)
	return newPackageList, err
}
