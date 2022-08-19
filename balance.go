package ebct

const balanceTrackingNumbers = DefaultEndPoint + Version + "packages/tracking-numbers/balance"

type BalanceResponse struct {
	MaxQuantity       *int `json:"maxQuantity,omitempty"`       // Postage card number used in the requisition.
	RequestedQuantity *int `json:"requestedQuantity,omitempty"` // Contract number associated with the postage card used.
	AvailableQuantity *int `json:"availableQuantity,omitempty"` // Correios sales team managing the contract and client.
	TotalUsed         *int `json:"totalUsed,omitempty"`         // não tem o que é isso no layout
	TotalUnused       *int `json:"totalUnused,omitempty"`       // não tem o que é isso no layout
}

func (n BalanceResponse) Error() string {
	panic("implement me")
}

func (c *Client) GetBalance() (*BalanceResponse, error) {
	newBalance := &BalanceResponse{}
	err := c.Get(balanceTrackingNumbers, nil, nil, newBalance)
	return newBalance, err
}
