package ebct

import (
	"errors"
)

const unitsAsync = DefaultEndPoint + Version + "units/async"
const units = DefaultEndPoint + Version + "units"

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
	Sequence *int `json:"sequence,omitempty"` // Sequencial number of the unit in the requisition.
	UnitCode *int `json:"unitCode,omitempty"` // Unit number code generated
}

type UnitsRequestGetAsync struct {
	RequestId *string `json:"requestId,omitempty"`
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
