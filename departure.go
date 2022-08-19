package ebct

import "time"

const departure = DefaultEndPoint + Version + "cn38request/departure"

type Departure struct {
	Cn38CodeList []string  `json:"cn38CodeList"` // Invoice numbers that will be boarded on the flight informed in the "flightList" element. One or more invoices can be included, as long as the flight details ("flightList") are the same.
	FlightList   []Flights `json:"flightList"`   // JSON element that will receive the keys and values with the flight data
}

type Flights struct {
	FlightNumber         *string    `json:"flightNumber"`         // Flight number
	AirlineCode          *string    `json:"airlineCode"`          // Airline code. This code must match the key ""code" contained in the return of the "Get Airline List" method
	DepartureDate        *time.Time `json:"departureDate"`        // Date time in the format "YYYY-MM-DDTHH:MM:SSZ", where the letters T and Z are fixed values
	DepartureAirportCode *string    `json:"departureAirportCode"` // Departure airport code/acronym
	ArrivalDate          *time.Time `json:"arrivalDate"`          // Date time in the format "YYYY-MM-DDTHH:MM:SSZ", where the letters T and Z are fixed values.
	ArrivalAirportCode   *string    `json:"arrivalAirportCode"`   // Airport code/acronym of arrival
}

func (n Departure) Error() string {
	panic("implement me")
}

func (c *Client) PutDepartureConfirmation(dep *Departure) (*Departure, error) {
	err := c.Put(departure, dep, nil, departure)
	return dep, err
}
