package ebct

import (
	"encoding/base64"
	"fmt"
)

type Client struct {
	Token      string
	Sandbox    bool
	BasicAuth  string
	Package    string
	ExpiryDate string
}

func NewClientToken(user, pass, contract string, sandbox bool) (*Client, error) {

	c := &Client{
		BasicAuth: fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(user+":"+pass))),
		Package:   contract,
		Sandbox:   sandbox,
	}

	tmpLoginToken := LoginToken{
		Numero: contract,
	}

	model := new(TokenResponse)

	err := c.Post(tokenAuth, tmpLoginToken, nil, model)

	if err == nil {
		if *model.Token != "" {
			c.Token = *model.Token
			c.ExpiryDate = *model.ExpiraEm
		}
	}

	return c, err
}

func NewClient(token string, sandbox bool) *Client {
	return &Client{
		Token:   token,
		Sandbox: sandbox,
	}
}

func (c *Client) GetToken() string {
	if c.Token != "" {
		return fmt.Sprintf("Bearer %s", c.Token)
	}
	return ""
}

func (c *Client) GetEndpoint() string {

	if c.Sandbox {
		return SandBoxUrl
	}

	return ProductionUrl

}
