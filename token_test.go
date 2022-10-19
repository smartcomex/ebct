package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient(
		"eyJhbGciOiJSUzUxMiJ9.eyJhbWJpZW50ZSI6IkhPTU9MT0dBQ0FPIiwiaWQiOiJ0ZXN0ZWludCIsInBmbCI6IlBKIiwicGpJbnRlcm5hY2lvbmFsIjo0MjM3NDg0MCwiYXBpIjpbODAsMzYsNTcsNTY0XSwiY2FydGFvLXBvc3RhZ2VtIjp7Im51bWVybyI6IjAwNzM0MjM3NzciLCJjb250cmF0byI6Ijk5OTIxNTc5NzAiLCJkciI6MzYsImFwaSI6WzgwLDU3LDgzXX0sImlwIjoiMTAuMC4yMDQuMTAsMTAuMC4yMDQuMTAiLCJpYXQiOjE2NjYxMDA5NDEsImlzcyI6InRva2VuLXNlcnZpY2UiLCJleHAiOjE2NjYxODczNDEsImp0aSI6IjllOWRkNTNkLTFkYTQtNDNhMy1hYzI3LWM2YzQ1NzRhMWFiOSJ9.g4eufn9IBLB8dWd-c-TOoCkcP0-bItgGK_KaJKuuC8FK7nEQv1X34NoZjqqRo0_i-wIOSr_qcdZ0mp8X7Wa86Au4SoEYrXeueN6KWyYLrpEslRaafJGHflnZicvv-uaMb879IBncE2tAdsiOEK8OJoY9fSRfBlg8fnu7wvjO-ONFRrmFsHso02cSdD8HYo30SV8HQtREsNzHkI6geGEWVV_Sp0VRzsUfPC4W5U7vKgHh0UQWtrgNwMH4SXI7runFSQvenFyJDvnrd6nlgsPJxw25DX4DSkgPmT6peayTYoJ7D5t38YdxFZfawPC1jRSLEtAQ94tapcdT7wu-Ft3EDA",
		true,
	)
}

func TestLogin(t *testing.T) {
	var err error

	client, err = NewClientToken("testeint", "NppSSY:4d:%x_b$)GQ7j-^}I5JF3MDY7i|2]yoz9", "0073423777", true)
	spew.Dump(client)
	assert.NoError(t, err)
}
