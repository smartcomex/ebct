package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAirline(t *testing.T) {
	vAirlines, err := client.GetAirline()
	spew.Dump(vAirlines)
	assert.NoError(t, err)
}
