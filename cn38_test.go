package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

const Cn38 = `{
	"dispatchNumbers": [151, 4587]
	}`

func TestCn38Request(t *testing.T) {

	aCn38Request := &Cn38RequestAsyncInput{
		DispatchNumbers: []int{151, 4587},
	}

	requestReturn, err := client.PostCn38Request(aCn38Request)
	spew.Dump(requestReturn)
	assert.NoError(t, err)
}

func TestCn38AsyncGetRequest(t *testing.T) {

	requestReturn, err := client.GetCn38Async("b34fdc7c-d9fc-4a99-8fdc-7cd9fcfa998e")
	spew.Dump(requestReturn)
	assert.NoError(t, err)
}
