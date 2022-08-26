package ebct

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

const UnitAsync = `{
    "dispatchNumber": 733,
    "originCountry": "US",
    "originOperatorName": "SKYG",
    "destinationOperatorName": "SAOD",
    "postalCategoryCode": "A",
    "serviceSubclassCode": "NX",
    "unitList": [
        {
            "sequence": 1,
            "unitType": 1,
            "trackingNumbers": [
                "NX000823744BR"
            ]
        }
    ]
}`

const Unit = `{
    "dispatchNumber": 12,
    "originCountry": "US",
    "originOperatorName": "SKYG",
    "destinationOperatorName": "SAOD",
    "postalCategoryCode": "A",
    "serviceSubclassCode": "NX",
    "unitList": [
        {
            "sequence": 1,
            "unitType": 1,
            "trackingNumbers": [
                "NX000931205BR",
                "NX000931219BR"
            ]
        }
    ]
}`

const UnitCodes = `{
    "originCountry": "CN",
    "postalCategoryCode": "A",
    "serviceSubclassCode": "XL",
    "unitType": 1,
    "quantity": 1
}`

func TestUnitAsyncRequest(t *testing.T) {

	aUnitAsync := &UnitsRequestAsync{}
	var bPack = []byte(UnitAsync)
	var err error
	err = json.Unmarshal(bPack, aUnitAsync)
	if err != nil {
		spew.Dump(err)
	}

	packList, err := client.PostUnitsRequestAsync(aUnitAsync)
	spew.Dump(packList)
	assert.NoError(t, err)
}

func TestUnitRequest(t *testing.T) {

	aUnitAsync := &UnitsRequestAsync{}
	var bPack = []byte(Unit)
	var err error
	err = json.Unmarshal(bPack, aUnitAsync)
	if err != nil {
		spew.Dump(err)
	}

	packList, err := client.PostUnitsRequest(aUnitAsync)
	spew.Dump(packList)
	assert.NoError(t, err)
}

func TestUnitAsyncGetRequest(t *testing.T) {

	packList, err := client.GetUnitsRequestAsync("f44435ce-1276-4c7d-8435-ce1276dc7d95")
	spew.Dump(packList)
	assert.NoError(t, err)
}

func TestUnitsCodesGetRequest(t *testing.T) {

	packList, err := client.GetUnitsCodesRequest()
	spew.Dump(packList)
	assert.NoError(t, err)
}

func TestUnitCodesRequest(t *testing.T) {

	stUnitCodes := &UnitCodeRequest{}
	var bPack = []byte(UnitCodes)
	var err error
	err = json.Unmarshal(bPack, stUnitCodes)
	if err != nil {
		spew.Dump(err)
	}

	packList, err := client.PostUnitCodesRequest(stUnitCodes)
	spew.Dump(packList)
	assert.NoError(t, err)
}

func TestUnitsCodesSingleRequestGet(t *testing.T) {

	packList, err := client.GetUnitsCodesSingleRequest("618a70b021a8bc472dd69abd")
	spew.Dump(packList)
	assert.NoError(t, err)
}

func TestUnitsRequestGet(t *testing.T) {

	packList, err := client.GetUnitsRequest(953)
	spew.Dump(packList)
	assert.NoError(t, err)
}

func TestUnitsRequestDelete(t *testing.T) {

	packList, err := client.DeleteUnitsCodesRequest("USDASDBRDDDDDIX00321001000024")
	spew.Dump(packList)
	assert.NoError(t, err)
}

func TestUnitsDispatchDelete(t *testing.T) {

	packList, err := client.DeleteUnitsDispatchRequest("USDASDBRDDDDDIX00321001000024")
	spew.Dump(packList)
	assert.NoError(t, err)
}
