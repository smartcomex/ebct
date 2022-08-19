package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient(
		"eyJhbGciOiJSUzUxMiJ9.eyJhbWJpZW50ZSI6IkhPTU9MT0dBQ0FPIiwiaWQiOiJ0ZXN0ZWludCIsInBmbCI6IlBKIiwicGpJbnRlcm5hY2lvbmFsIjo0MjM3NDg0MCwiYXBpIjpbODAsMzYsNTddLCJjYXJ0YW8tcG9zdGFnZW0iOnsibnVtZXJvIjoiMDA3MzQyMzc3NyIsImNvbnRyYXRvIjoiOTk5MjE1Nzk3MCIsImRyIjozNiwiYXBpIjpbODAsNTcsODNdfSwiaXAiOiIxOTEuMTYyLjEyNy42NywxOTEuMTYyLjEyNy42NyIsImlhdCI6MTY2MDkxNTEzNCwiaXNzIjoidG9rZW4tc2VydmljZSIsImV4cCI6MTY2MTAwMTUzNCwianRpIjoiYTAxNDM5NzgtM2FkNS00NTliLWEwOTUtZWRlMTk4ZjY1YmQyIn0.G4RM38d5Zvt00q2Qw0vxLphN_drIl1smVlveBtrPCQqZr-3qHOWtLE3V3vPmOZ1tVkJmJMezm2fH-ZH-x4iq4DMkAMd-o3qs8Cgwu7qmYsDmETTYdoP9CeEBgKL04GVUaosJkhWR0CZkqcGmshLARyerYOWaWTUY_4ao2lDZsD8SvkpxrwfbkWt6zlkh_ecBg-xaOoA1XQHOsBu5Gy1BGL49D7_Hnwj0W4jE4wmDaCDnsdkbD9glAEpspFSIQXA0TlrCQRX4sTrh4Ei0iZZFiGDofQ-g1frq394rqAv9_KlJwdfCTfnaDO4Ovt_r5luu7VnSf9z0q6Kog6XpyiHDYQ",
		true,
	)
}

func TestLogin(t *testing.T) {
	var err error

	client, err = NewClientToken("testeint", "NppSSY:4d:%x_b$)GQ7j-^}I5JF3MDY7i|2]yoz9", "0073423777", true)
	spew.Dump(client)
	assert.NoError(t, err)
}
