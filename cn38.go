package ebct

import (
	"errors"
	"time"
)

const cn38async = DefaultEndPoint + Version + "cn38request"
const cn38Units = DefaultEndPoint + Version + "cn38request/units"
const cn38ListsPending = DefaultEndPoint + Version + "cn38request/departure/pending"
const cn38ListsConfirmed = DefaultEndPoint + Version + "cn38request/departure/confirmed"
const cn38ListsGenerated = DefaultEndPoint + Version + "cn38request/generated"

type Cn38RequestAsyncInput struct {
	DispatchNumbers []int `json:"dispatchNumbers"` // This field should receive a list with all the numbers of orders that will make up this delivery note, and can have up to 100 orders on the same delivery note. The dispatchNumbers entered in this list must be the same as those passed in the "dispatchNumbers" parameter of the request for the Units Request method.
}

type Cn38RequestReturn struct {
	RequestId      *string `json:"requestId,omitempty"`     // Value generated to identify the request. This data must be passed as the value of the key "requestid" of the GET method "Get CN38 async request info", used to obtain the delivery invoice number.
	DispatchNumber []int   `json:"dispatchNumber"`          // Contains the list of dispatchNumbers informed in the method request.
	RequestStatus  *string `json:"requestStatus,omitempty"` // Contains the status of the request. Pending: Pending processing; Processing: In processing; Success: Processing was completed successfully. Error: An error has occurred in the processing of the information. The error message is displayed in the “errorMessage” field.
	Cn38Code       *string `json:"cn38Code,omitempty"`      // Invoice number to be informed on form CN38 / Delivery Invoice maximum size = 12 characters
}

type Cn38RequestGetAsync struct {
	RequestId *string `json:"requestId,omitempty"`
}

func (n Cn38RequestReturn) Error() string {
	panic("implement me")
}

type Cn38UnitList struct {
	UnitList []Cn38UnitCodes `json:"unitList,omitempty"`
}

type Cn38UnitCodes struct {
	UnitCode        *string  `json:"unitCode,omitempty"`        // Receives a list of unitizers/bags and the packages contained in each one
	TrackingNumbers []string `json:"trackingNumbers,omitempty"` // List of packages (trackingNumbers) contained in each bag (unitCode)
}

type Cn38UnitCodesReturn struct {
	RequestId     *string `json:"requestId,omitempty"`     // Receives a list of unitizers/bags and the packages contained in each one
	RequestStatus *string `json:"requestStatus,omitempty"` // List of packages (trackingNumbers) contained in each bag (unitCode)
}

func (n Cn38UnitCodesReturn) Error() string {
	panic("implement me")
}

type Cn38ListDeparture struct {
	Page *int `json:"page,omitempty"`
}

type Cn38Itens struct {
	Cn38Code                  string    `json:"cn38Code"`                  // Bill number
	GenerationDate            time.Time `json:"generationDate"`            // Generation date
	DepartureConfirmed        bool      `json:"departureConfirmed"`        // Boarding status
	DepartureConfirmationDate time.Time `json:"departureConfirmationDate"` // Departure Confirmation Date
	FlightList                []Flights `json:"flightList"`
}

type Cn38Page struct {
	Size           int  `json:"size"`           // Default page size
	NumberElements int  `json:"numberElements"` // Effective number of elements on the current page
	Number         int  `json:"number"`         // Page number. The first is 0.
	Next           bool `json:"next"`           // Indicates if there is next page
	Previous       bool `json:"previous"`       // Indicates if there is a previous page
	First          bool `json:"first"`          // Indicates if this is the first page
	Last           bool `json:"last"`           // Indicates if it is the last page
}

type Cn38DepartureList struct {
	Itens []Cn38Itens `json:"itens"`
	Page  Cn38Page    `json:"page"`
}

func (n Cn38DepartureList) Error() string {
	panic("implement me")
}

func (c *Client) PostCn38Request(dispatches *Cn38RequestAsyncInput) (*Cn38RequestReturn, error) {

	newcn38RequestReturn := &Cn38RequestReturn{}
	err := c.Post(cn38async, dispatches, nil, newcn38RequestReturn)
	return newcn38RequestReturn, err
}

func (c *Client) GetCn38Async(requestId string) (*Cn38RequestReturn, error) {

	sendRequest := &Cn38RequestGetAsync{}
	var err error

	if !IsUUID(requestId) {
		err = errors.New("the requestId must be a valid UUID")
		return nil, err
	}
	sendRequest.RequestId = StringPtr(requestId)

	newCn38RequestReturn := &Cn38RequestReturn{}
	err = c.Get(cn38async, sendRequest, nil, newCn38RequestReturn)
	return newCn38RequestReturn, err
}

func (c *Client) PostCn38UnitCode(unitcodes *Cn38UnitList) (*Cn38UnitCodesReturn, error) {

	newcn38RequestReturn := &Cn38UnitCodesReturn{}
	err := c.Post(cn38Units, unitcodes, nil, newcn38RequestReturn)
	return newcn38RequestReturn, err
}

func (c *Client) GetCn38ListPending(page int) (*Cn38DepartureList, error) {

	sendRequest := &Cn38ListDeparture{}
	var err error
	newCn38RequestReturn := &Cn38DepartureList{}

	if page > 0 {
		sendRequest.Page = IntPtr(page)
		err = c.Get(cn38ListsPending, sendRequest, nil, newCn38RequestReturn)
	} else {
		err = c.Get(cn38ListsPending, nil, nil, newCn38RequestReturn)
	}

	return newCn38RequestReturn, err
}

func (c *Client) GetCn38ListConfirmed(page int) (*Cn38DepartureList, error) {

	sendRequest := &Cn38ListDeparture{}
	var err error
	newCn38RequestReturn := &Cn38DepartureList{}

	if page > 0 {
		sendRequest.Page = IntPtr(page)
		err = c.Get(cn38ListsConfirmed, sendRequest, nil, newCn38RequestReturn)
	} else {
		err = c.Get(cn38ListsConfirmed, nil, nil, newCn38RequestReturn)
	}

	return newCn38RequestReturn, err
}

func (c *Client) GetCn38ListGenerated(page int) (*Cn38DepartureList, error) {

	sendRequest := &Cn38ListDeparture{}
	var err error
	newCn38RequestReturn := &Cn38DepartureList{}

	if page > 0 {
		sendRequest.Page = IntPtr(page)
		err = c.Get(cn38ListsGenerated, sendRequest, nil, newCn38RequestReturn)
	} else {
		err = c.Get(cn38ListsGenerated, nil, nil, newCn38RequestReturn)
	}

	return newCn38RequestReturn, err
}

func (c *Client) GetCn38ListGeneratedByInvoice(cn38Number string) (*Cn38Itens, error) {

	var err error
	newCn38RequestReturn := &Cn38Itens{}
	err = c.Get(cn38ListsGenerated+"/"+cn38Number, nil, nil, newCn38RequestReturn)
	return newCn38RequestReturn, err
}
