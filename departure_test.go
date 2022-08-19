package ebct

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const departureConfirmation = `{
    "cn38CodeList": [
        "10432021"
    ],
    "flightList": [
        {
            "flightNumber": "9010",
            "airlineCode": "AE",
            "departureDate": "2021-06-15T21:05:00Z",
            "departureAirportCode": "MI",
            "arrivalDate": "2021-06-16T22:05:00Z",
            "arrivalAirportCode": "GRU"
        }
    ]
}`

func TestDepartureConfirmation(t *testing.T) {

	dep := &Departure{}
	var bPack = []byte(departureConfirmation)
	var err error
	err = json.Unmarshal(bPack, dep)
	if err != nil {
		spew.Dump(err)
	}

	dep.FlightList[0].DepartureDate = TimePtr(time.Now().AddDate(0, 0, -1))
	dep.FlightList[0].ArrivalDate = TimePtr(time.Now().AddDate(0, 0, -1))

	packList, err := client.PutDepartureConfirmation(dep)
	spew.Dump(packList)
	assert.NoError(t, err)
}
