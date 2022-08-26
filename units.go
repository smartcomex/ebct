package ebct

import (
	"errors"
)

const unitsAsync = DefaultEndPoint + Version + "units/async"
const units = DefaultEndPoint + Version + "units"
const unitsCodes = DefaultEndPoint + Version + "units/unit-codes"
const unitsDispatch = DefaultEndPoint + Version + "units/dispatch"

type UnitsRequestAsync struct {
	DispatchNumber          *int            `json:"dispatchNumber,omitempty"`          // Customer control number to reference the dispatch of that container. This same number should be used as a parameter for generating the Delivery Bill (Similar to CN38 postal form), method CN38 Request (Async). This value cannot be repeated for different shipments
	OriginCountry           *string         `json:"originCountry,omitempty"`           // Inform the origin country abbreviation. The abbreviation should meet the ISO 3166-2.
	OriginOperatorName      *string         `json:"originOperatorName,omitempty"`      // Inform the origin operator name. Use the first 4 letters of your company name.
	DestinationOperatorName *string         `json:"destinationOperatorName,omitempty"` // Inform the destination operator name. CRBA - For Curitiba
	PostalCategoryCode      *string         `json:"postalCategoryCode,omitempty"`      // Inform the postal category code. A – Airmail or Priority Mail; B – S.A.L Mail or Non-Priority Mail; C – Surface Mail or Non-Priority Mail; D – Priority Mail sent by surface transportation.
	ServiceSubclassCode     *string         `json:"serviceSubclassCode,omitempty"`     // Inform the service subclass code. It should be one of the codes below: NX –Standard service IX – Express service XP – Mini/economy service There must be only one type of postal service per container
	UnitList                []UnitsListCntr `json:"unitList,omitempty"`                // Element where the objects contained in the container will be listed.
}

type UnitsListCntr struct {
	Sequence        *int     `json:"sequence,omitempty"`        // Inform the unit sequence in the current dispatch. The maximum number of units inside a dispatch is 500. minimum: 1 maximum: 500
	UnitType        *int     `json:"unitType,omitempty"`        // Inform the unit type: 1 - Bag up to 30kg 2 – Box with base pallet up to 500kg
	TrackingNumbers []string `json:"trackingNumbers,omitempty"` // List of all objects that are contained in the pool. It can hold up to 500 objects. These numbers are the values of the "trackingNumber" key of the Package Request () method return Up to 500 objects at a time, listed in the JSON standard
}

type UnitsRequestReturn struct {
	RequestId        *string             `json:"requestId,omitempty"`     // Value generated to identify the request. This data must be passed as the value of the key "requestid" of the GET method "Get unit async request info", used to obtain the unit number.
	RequestStatus    *string             `json:"requestStatus,omitempty"` // Contains the status of the request. Pending: Pending processing; Processing: In processing; Success: Processing was completed successfully. Error: An error has occurred in the processing of the information. The error message is displayed in the “errorMessage” field.
	Request          *UnitsRequestAsync  `json:"request,omitempty"`       // Same as request Sent
	UnitResponseList []UnitsResponseList `json:"unitResponseList,omitempty"`
}

type UnitsResponseList struct {
	Sequence *int    `json:"sequence,omitempty"` // Sequencial number of the unit in the requisition.
	UnitCode *string `json:"unitCode,omitempty"` // Unit number code generated
}

type UnitsRequestGetAsync struct {
	RequestId *string `json:"requestId,omitempty"`
}

type UnitsCodesReturn struct {
	RequestedQuantity *string `json:"requestedQuantity,omitempty"` // Request number returned by the API when calling the "Unit codes request" method.
	AvailableQuantity *int    `json:"availableQuantity,omitempty"` // Number of unitCodes contained in the request
	TotalUsed         *int    `json:"totalUsed,omitempty"`         // Number of unitCodes contained in the request and which have already been used
	TotalUnused       *int    `json:"totalUnused,omitempty"`       // Number of unitCodes contained in the request that have not yet been used
	Links             []Link  `json:"links,omitempty"`             // This link can be sent in a GET request to obtain details of the request whose identifier appears in the key "requestId".
}

type Link struct {
	Rel  *string `json:"rel,omitempty"`
	Href *string `json:"href,omitempty"`
}

type UnitCodeRequest struct {
	OriginCountry       string `json:"originCountry"`       // Inform the origin country abbreviation. The abbreviation should meet the ISO 3166-2.
	PostalCategoryCode  string `json:"postalCategoryCode"`  // Inform the postal category code. A – Airmail or Priority Mail; B – S.A.L Mail or Non-Priority Mail; C – Surface Mail or Non-Priority Mail; D – Priority Mail sent by surface transportation.
	ServiceSubclassCode string `json:"serviceSubclassCode"` // Inform the service subclass code. It should be one of the codes below: NX –Standard service IX – Express service XP – Mini/economy service There must be only one type of postal service per container
	UnitType            int    `json:"unitType"`            // Inform the unit type:  1 - Bag up to 30kg 2 – Box with base pallet up to 500kg
	Quantity            int    `json:"quantity"`            // Amount of unit codes desired
}

type UnitCodeReturn struct {
	UnitCodes []string `json:"unitCodes,omitempty"`
}

func (n UnitCodeReturn) Error() string {
	panic("implement me")
}

type UnitCodeCompleteReturn struct {
	RequestId           *string             `json:"requestId"`           // Request number returned by the API when calling the "Unit codes request" method.
	RequestDate         *string             `json:"requestDate"`         // Date and time the request was created. It is in the pattern "YYYY-MM-DDThh:mm:ss.sss".
	OriginCountry       *string             `json:"originCountry"`       // Abbreviation of the country of origin of the shipment, entered when the requisition was created
	PostalCategoryCode  *string             `json:"postalCategoryCode"`  // Postal category code (A, B, C or D)
	ServiceSubclassCode *string             `json:"serviceSubclassCode"` // Category of service used: NX –Standard service IX – Express service XP – Mini/economy service
	UnitType            *int                `json:"unitType"`            // Unit type: 1 - Bag up to 30kg 2 – Box with base pallet up to 500kg
	RequestedQuantity   *int                `json:"requestedQuantity"`   // Number of unitizer (unitCode) numbers generated in the request.
	TotalUsed           *int                `json:"totalUsed"`           // Number of unitCodes contained in the request and which have already been used.
	TotalUnused         *int                `json:"totalUnused"`         // Number of unitCodes contained in the request that have not yet been used.
	UnusedCodes         []UnitsResponseList `json:"unusedCodes"`         // JSON arrays containing all unused unitCodes and their respective sequence
	UsedCodes           []UnitsResponseList `json:"usedCodes"`           // JSON arrays containing all the unitCodes already used and their respective sequence
}

func (n UnitCodeCompleteReturn) Error() string {
	panic("implement me")
}

type UnitsDispatch struct {
	DispatchNumber *int `json:"dispatchNumber,omitempty"` // Inform the dispatch number of the unit.
}

type UnitsListCodeNumber struct {
	UnitCode        *string  `json:"unitCode,omitempty"`        // Inform the unit type: 1 - Bag up to 30kg 2 – Box with base pallet up to 500kg
	TrackingNumbers []string `json:"trackingNumbers,omitempty"` // List of all objects that are contained in the pool. It can hold up to 500 objects. These numbers are the values of the "trackingNumber" key of the Package Request () method return Up to 500 objects at a time, listed in the JSON standard
}

type UnitsDispatchReturn struct {
	DispatchNumber *int                  `json:"dispatchNumber,omitempty"` // Customer control number to reference the dispatch of that container. This same number should be used as a parameter for generating the Delivery Bill (Similar to CN38 postal form), method CN38 Request (Async). This value cannot be repeated for different shipments
	UnitList       []UnitsListCodeNumber `json:"unitList,omitempty"`       // Element where the objects contained in the container will be listed.
}

func (n UnitsDispatchReturn) Error() string {
	panic("implement me")
}

//type UnitsRequestGetReturn struct {
//	RequestId        *string             `json:"requestId"`     // Value generated to identify the request. This data must be passed as the value of the key "requestid" of the GET method "Get unit async request info", used to obtain the unit number.
//	RequestStatus    *string             `json:"requestStatus"` // Contains the status of the request. Pending: Pending processing; Processing: In processing; Success: Processing was completed successfully. Error: An error has occurred in the processing of the information. The error message is displayed in the “errorMessage” field.
//	Request          *UnitsRequestAsync  `json:"request"`       // Same as request Sent
//	UnitResponseList []UnitsResponseList `json:"unitResponseList"`
//}

func (n UnitsRequestReturn) Error() string {
	panic("implement me")
}

func (c *Client) PostUnitsRequestAsync(UnitRequest *UnitsRequestAsync) (*UnitsRequestReturn, error) {
	newUnitsRequestReturn := &UnitsRequestReturn{}
	err := c.Post(unitsAsync, UnitRequest, nil, newUnitsRequestReturn)
	return newUnitsRequestReturn, err
}

func (c *Client) PostUnitsRequest(UnitRequest *UnitsRequestAsync) (*UnitsRequestReturn, error) {
	newUnitsRequestReturn := &UnitsRequestReturn{}
	err := c.Post(units, UnitRequest, nil, newUnitsRequestReturn)
	return newUnitsRequestReturn, err
}

func (c *Client) GetUnitsRequestAsync(requestId string) (*UnitsRequestReturn, error) {

	sendRequest := &UnitsRequestGetAsync{}
	var err error

	if !IsUUID(requestId) {
		err = errors.New("the requestId must be a valid UUID")
		return nil, err
	}
	sendRequest.RequestId = StringPtr(requestId)

	newUnitsRequestReturn := &UnitsRequestReturn{}
	err = c.Get(unitsAsync, sendRequest, nil, newUnitsRequestReturn)
	return newUnitsRequestReturn, err
}

func (c *Client) GetUnitsCodesRequest() (*UnitsCodesReturn, error) {

	newUnitsRequestReturn := &UnitsCodesReturn{}
	err := c.Get(unitsCodes, nil, nil, newUnitsRequestReturn)
	return newUnitsRequestReturn, err
}

func (c *Client) PostUnitCodesRequest(UnitRequest *UnitCodeRequest) (*UnitCodeReturn, error) {
	newUnitsRequestReturn := &UnitCodeReturn{}
	err := c.Post(unitsCodes, UnitRequest, nil, newUnitsRequestReturn)
	return newUnitsRequestReturn, err
}

func (c *Client) GetUnitsCodesSingleRequest(requestId string) (*UnitCodeCompleteReturn, error) {

	newUnitsRequestReturn := &UnitCodeCompleteReturn{}
	err := c.Get(unitsCodes+"/"+requestId, nil, nil, newUnitsRequestReturn)
	return newUnitsRequestReturn, err
}

func (c *Client) GetUnitsRequest(dispatchNumber int) (*UnitsDispatchReturn, error) {

	dispatch := &UnitsDispatch{
		DispatchNumber: IntPtr(dispatchNumber),
	}

	newUnitsRequestReturn := &UnitsDispatchReturn{}
	err := c.Get(units, dispatch, nil, newUnitsRequestReturn)
	return newUnitsRequestReturn, err
}

func (c *Client) DeleteUnitsCodesRequest(unitCode string) (*UnitsResponseList, error) {

	newUnitsRequestReturn := &UnitsResponseList{}
	err := c.Delete(units+"/"+unitCode, nil, nil, newUnitsRequestReturn)
	return newUnitsRequestReturn, err
}

func (c *Client) DeleteUnitsDispatchRequest(dispatch string) (*UnitsDispatch, error) {

	newUnitsRequestReturn := &UnitsDispatch{}
	err := c.Delete(unitsDispatch+"/"+dispatch, nil, nil, newUnitsRequestReturn)
	return newUnitsRequestReturn, err
}
