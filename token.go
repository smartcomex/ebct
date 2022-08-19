package ebct

const tokenAuth = "token/" + Version + "autentica/cartaopostagem"

type TokenResponse struct {
	Ambiente        *string        `json:"Ambiente,omitempty"`        // “HOMOLOGACAO” = Test environment “PRODUCAO” = Real environment
	Id              *string        `json:"id,omitempty"`              // User idCorreios credential used in method requisition.
	Perfil          *string        `json:"Perfil,omitempty"`          // PJ = Corporate PF = Natural Person
	PjInternacional *int           `json:"pjInternacional,omitempty"` // Customer contract code, with Correios.
	CartaoPostagem  CartaoPostagem `json:"cartaoPostagem,omitempty"`  // Customer contract code, with Correios.
	Api             []*int         `json:"api,omitempty"`             // não tem o que é isso no layout
	Ip              *string        `json:"ip,omitempty"`              // Contains the IP of the machine from which the request originated
	Emissao         *string        `json:"emissao,omitempty"`         // Date when the token was generated.
	ExpiraEm        *string        `json:"expiraEm,omitempty"`        // Date when the token will expire.
	Token           *string        `json:"token,omitempty"`           // The value contained in the key “token” must be used in all requests made during its validity. This token must be inserted in the Request Header, as part of the value of the Authorization key.
}

type CartaoPostagem struct {
	Numero   *string `json:"numero,omitempty"`   // Postage card number used in the requisition.
	Contrato *string `json:"contrato,omitempty"` // Contract number associated with the postage card used.
	Dr       *int    `json:"dr,omitempty"`       // Correios sales team managing the contract and client.
	Api      []*int  `json:"api,omitempty"`      // não tem o que é isso no layout
}

type LoginToken struct {
	Numero string `json:"numero,omitempty"`
}

func (n TokenResponse) Error() string {
	panic("implement me")
}
