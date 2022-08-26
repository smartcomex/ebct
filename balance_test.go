package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBalanceTackingNumber(t *testing.T) {
	vBal, err := client.GetBalance()
	spew.Dump(vBal)
	assert.NoError(t, err)
}

func TestUnitCodeBalance(t *testing.T) {
	vBal, err := client.GetUnitCodeBalance()
	spew.Dump(vBal)
	assert.NoError(t, err)
}

func TestBalanceInput(t *testing.T) {
	vBal, err := client.GetBalanceInput()
	spew.Dump(vBal)
	assert.NoError(t, err)
}
