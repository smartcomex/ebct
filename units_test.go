package ebct

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

const UnitAsync = `{
    "dispatchNumber": 953,
    "originCountry": "BR",
    "originOperatorName": "FOXB",
    "destinationOperatorName": "SAOD",
    "postalCategoryCode": "A",
    "serviceSubclassCode": "NX",
    "unitList": [
        {
            "sequence": 1,
            "unitType": 1,
            "trackingNumbers": [
                "NX006288545BR"
            ]
        }
    ]
}`

const Unit = `{
    "dispatchNumber": 61523,
    "originCountry": "CN",
    "originOperatorName": "FOXB",
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

const UnitCodes = `{
    "originCountry": "CN",
    "postalCategoryCode": "A",
    "serviceSubclassCode": "NL",
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
	var packList *UnitsRequestReturn

	//for i := 1; i < 999999; i++ {
	//	aUnitAsync.DispatchNumber = IntPtr(i)
	packList, err = client.PostUnitsRequest(aUnitAsync)
	//	if err == nil {
	//		break
	//	}
	//}
	spew.Dump(packList)
	assert.NoError(t, err)
	spew.Dump(aUnitAsync)
}

func TestUnitAsyncGetRequest(t *testing.T) {

	packList, err := client.GetUnitsRequestAsync("7d607b7b-9848-41b1-a07b-7b9848e1b115")
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

	packList, err := client.DeleteUnitsCodesRequest("BRFOXBBRSAODANX20953001000001")
	spew.Dump(packList)
	assert.NoError(t, err)
}

func TestUnitsDispatchDelete(t *testing.T) {

	packList, err := client.DeleteUnitsDispatchRequest("USDASDBRDDDDDIX00321001000024")
	spew.Dump(packList)
	assert.NoError(t, err)
}
