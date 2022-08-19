package ebct

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

const TrackingInput = `{
    "distributionModality": 33170,
    "quantity": 5
}`

func TestTrackingNumbersGet(t *testing.T) {

	packList, err := client.GetTrackingNumbers()
	spew.Dump(packList)
	assert.NoError(t, err)
}

func TestTrackingNumbersPost(t *testing.T) {

	trackingNumbers := &TrackingNumbersInput{}
	var bPack = []byte(TrackingInput)
	var err error
	err = json.Unmarshal(bPack, trackingNumbers)
	if err != nil {
		spew.Dump(err)
	}

	packList, err := client.PostTrackingNumbers(trackingNumbers)
	spew.Dump(packList)
	assert.NoError(t, err)
}
