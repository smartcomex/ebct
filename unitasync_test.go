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
