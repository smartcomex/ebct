package ebct

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

const Cn38UnitCode = `{"unitList": [
	{
		"unitCode": "CNTICPBRSAODAIX12374001000030",
		"trackingNumbers": [
		"IX000682851BR","IX000682865BR"
	] },
	{
		"unitCode": "CNTICPBRSAODAIX12375001000030",
		"trackingNumbers": [
		"IX000682879BR"
	]
	}]}`

func TestCn38Request(t *testing.T) {

	aCn38Request := &Cn38RequestAsyncInput{
		DispatchNumbers: []int{953},
	}

	requestReturn, err := client.PostCn38Request(aCn38Request)
	spew.Dump(requestReturn)
	assert.NoError(t, err)
}

func TestCn38AsyncGetRequest(t *testing.T) {

	requestReturn, err := client.GetCn38Async("85092678-4a81-410b-8926-784a81910b29")
	spew.Dump(requestReturn)
	assert.NoError(t, err)
}

func TestCn38UnitCode(t *testing.T) {

	aCn38Request := &Cn38UnitList{}
	var bPack = []byte(Cn38UnitCode)
	var err error
	err = json.Unmarshal(bPack, aCn38Request)
	if err != nil {
		spew.Dump(err)
	}

	requestReturn, err := client.PostCn38UnitCode(aCn38Request)
	spew.Dump(requestReturn)
	assert.NoError(t, err)
}

func TestCn38ListPendingDepartureGet(t *testing.T) {

	requestReturn, err := client.GetCn38ListPending(0)
	spew.Dump(requestReturn)
	assert.NoError(t, err)
}

func TestCn38ListConfirmedDepartureGet(t *testing.T) {

	requestReturn, err := client.GetCn38ListConfirmed(0)
	spew.Dump(requestReturn)
	assert.NoError(t, err)
}

func TestCn38ListGeneratedDepartureGet(t *testing.T) {

	requestReturn, err := client.GetCn38ListGenerated(0)
	spew.Dump(requestReturn)
	assert.NoError(t, err)
}

func TestCn38ListGeneratedByInvoiceGet(t *testing.T) {

	requestReturn, err := client.GetCn38ListGeneratedByInvoice("9012021")
	spew.Dump(requestReturn)
	assert.NoError(t, err)
}
