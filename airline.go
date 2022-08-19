package ebct

const airline = DefaultEndPoint + Version + "cn38request/departure/airlines"

type Airline struct {
	Name *string `json:"name,omitempty"`
	Code *string `json:"code,omitempty"`
}

func (n Airline) Error() string {
	panic("implement me")
}

func (c *Client) GetAirline() ([]Airline, error) {
	var aAirlines []Airline
	err := c.Get(airline, nil, nil, aAirlines)
	return aAirlines, err
}
