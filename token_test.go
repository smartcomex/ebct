package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient(
		"eyJhbGciOiJSUzUxMiJ9.eyJhbWJpZW50ZSI6IkhPTU9MT0dBQ0FPIiwiaWQiOiJ0ZXN0ZWludCIsInBmbCI6IlBKIiwicGpJbnRlcm5hY2lvbmFsIjo0MjM3NDg0MCwiYXBpIjpbODAsMzYsNTcsNTY0XSwiY2FydGFvLXBvc3RhZ2VtIjp7Im51bWVybyI6IjAwNzM0MjM3NzciLCJjb250cmF0byI6Ijk5OTIxNTc5NzAiLCJkciI6MzYsImFwaSI6WzgwLDU3LDgzXX0sImlwIjoiMTAuMC4yMDQuMTAsMTAuMC4yMDQuMTAiLCJpYXQiOjE2NjQ4OTE2MDMsImlzcyI6InRva2VuLXNlcnZpY2UiLCJleHAiOjE2NjQ5NzgwMDMsImp0aSI6IjFmOTIyOTYwLTUzMTQtNGI2ZC04OTNiLWExNzIzMzNhZmM4ZSJ9.aC1OcdYExueoegPFFD20_JRsOm2SatB4sNtiqEvmj3uN94_26SfP2k0qVG6YxLNA0msR4tAKTD-shV00B_eUrs7P7d6Iwoj9znAYCV-yemSzaEatsr2FNklS8oXqgfZMtJquFfbmE2O9nKJI03QwHgUdOqs_1UOOBxdT0I7ThEUH10k7-S3XXODywQQ1_1Hqdb5POIu49FTx2T7BEU56ho7Vdu9Hk7PgzAijwT9uJqQW50Q7YEyYsKpTqckAYWehLqStio6ralYXXdj3cLeNH-Q0OYJ-z_h27qTF0JyjnFjICluMeKb8SVp3sqGAIc07t8fu0nQiw3M1ym9FuA4rUA",
		true,
	)
}

func TestLogin(t *testing.T) {
	var err error

	client, err = NewClientToken("testeint", "NppSSY:4d:%x_b$)GQ7j-^}I5JF3MDY7i|2]yoz9", "0073423777", true)
	spew.Dump(client)
	assert.NoError(t, err)
}
