package ebct

import "errors"

const cn38async = DefaultEndPoint + Version + "cn38request"

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
