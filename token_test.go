package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient(
		"eyJhbGciOiJSUzUxMiJ9.eyJhbWJpZW50ZSI6IkhPTU9MT0dBQ0FPIiwiaWQiOiJ0ZXN0ZWludCIsInBmbCI6IlBKIiwicGpJbnRlcm5hY2lvbmFsIjo0MjM3NDg0MCwiYXBpIjpbODAsMzYsNTddLCJjYXJ0YW8tcG9zdGFnZW0iOnsibnVtZXJvIjoiMDA3MzQyMzc3NyIsImNvbnRyYXRvIjoiOTk5MjE1Nzk3MCIsImRyIjozNiwiYXBpIjpbODAsNTcsODNdfSwiaXAiOiIxMC4wLjIwNC4xMCwxMC4wLjIwNC4xMCIsImlhdCI6MTY2Mzg2ODEwNywiaXNzIjoidG9rZW4tc2VydmljZSIsImV4cCI6MTY2Mzk1NDUwNywianRpIjoiMTljNGMxMzItZWQzMi00ZGNlLWI1NWUtMDAyZTgzM2VjNzczIn0.eRzeDpvTo-J8irkaLTB3OplTxsk4LH-byqYRgARbJ7a9k8N2MsNAZwQPtcCuvPymCl_8taSN3PNXqGS_5pEuIRw34m6PI0ITt_WGRMzE9ixQnNcXLXgobEUmm8khIwAKG7TcBAoeIrKgBUnq9HLqNJ5v2loBnFc7XR3_TeyBRFJLycN2RXbnpTH2Ngaus3cuI3v7B8Z3zxgUGzqSZo5ZfMljCosMqDv6r0mT6WM3MDQ1bvgIIaTwpH2JVlFBl3Rmbxa8c0kw4N-kl9m_LUIKqe7BdsOEbzl-8_rbn9xqwskgHxOmjkWF--XP6RbAG7JvvFfDU4BlEhEx_zJ8OfsvrQ",
		true,
	)
}

func TestLogin(t *testing.T) {
	var err error

	client, err = NewClientToken("testeint", "NppSSY:4d:%x_b$)GQ7j-^}I5JF3MDY7i|2]yoz9", "0073423777", true)
	spew.Dump(client)
	assert.NoError(t, err)
}
