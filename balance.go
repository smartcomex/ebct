package ebct

const balanceTrackingNumbers = DefaultEndPoint + Version + "packages/tracking-numbers/balance"

const balanceUnitsCode = DefaultEndPoint + Version + "units/unit-codes/balance"

const balance = DefaultEndPoint + Version + "balance"

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

type BalanceInput struct {
	CurrentBalance *float64 `json:"currentBalance,omitempty"` // Existing balance in Reais, the Brazilian currency
}

func (n BalanceInput) Error() string {
	panic("implement me")
}

func (c *Client) GetBalance() (*BalanceResponse, error) {
	newBalance := &BalanceResponse{}
	err := c.Get(balanceTrackingNumbers, nil, nil, newBalance)
	return newBalance, err
}

func (c *Client) GetUnitCodeBalance() (*BalanceResponse, error) {
	newBalance := &BalanceResponse{}
	err := c.Get(balanceUnitsCode, nil, nil, newBalance)
	return newBalance, err
}

func (c *Client) GetBalanceInput() (*BalanceInput, error) {
	newBalance := &BalanceInput{}
	err := c.Get(balance, nil, nil, newBalance)
	return newBalance, err
}
