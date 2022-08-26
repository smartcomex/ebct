package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient(
		"eyJhbGciOiJSUzUxMiJ9.eyJhbWJpZW50ZSI6IkhPTU9MT0dBQ0FPIiwiaWQiOiJ0ZXN0ZWludCIsInBmbCI6IlBKIiwicGpJbnRlcm5hY2lvbmFsIjo0MjM3NDg0MCwiYXBpIjpbODAsMzYsNTddLCJjYXJ0YW8tcG9zdGFnZW0iOnsibnVtZXJvIjoiMDA3MzQyMzc3NyIsImNvbnRyYXRvIjoiOTk5MjE1Nzk3MCIsImRyIjozNiwiYXBpIjpbODAsNTcsODNdfSwiaXAiOiIxODkuNzcuOTIuMTQxLDE4OS43Ny45Mi4xNDEiLCJpYXQiOjE2NjE1MzUzMzAsImlzcyI6InRva2VuLXNlcnZpY2UiLCJleHAiOjE2NjE2MjE3MzAsImp0aSI6IjgzYmQ3MDY2LWQ3N2ItNDljMy05YmIyLTMyOTkxMjZlZjE3NCJ9.Ibb9aJ2CM2DoDZZOzMIO_T3v4nKa-EcquVJKWhgSGWWJZz-ID_A-HWPiHRaPWAL3a8r10mGI_13M5LiEJgmR4nVRqnOI9826CjCrCgWgo20HDPq8ANVg6mLc6S3Wq6RYdTUVGm1YIDz-G3naqOBxcI8vB5QIi2gO94nIqVjKIo0kBc1sau57IDAJ0VkIMDkc8akjoYeqpwSfM-dRDd2MboZ381xM13n3eOto6DagRUZ17Fbl6-pV2gJ6te0AGzv49_qX6X_cakqEDeTI1YoFw7bvZwi3VybMfokuZP_ToQ6P2_I9gdr1P9qrOb4fMMltulJFyjtjaVpt_AHkNM7NqA",
		true,
	)
}

func TestLogin(t *testing.T) {
	var err error

	client, err = NewClientToken("testeint", "NppSSY:4d:%x_b$)GQ7j-^}I5JF3MDY7i|2]yoz9", "0073423777", true)
	spew.Dump(client)
	assert.NoError(t, err)
}
