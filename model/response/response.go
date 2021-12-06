package response

type CreateAndUpdate struct {
	Id     string  `json:"id"`
	Code   string  `json:"code"`
	Amount int64   `json:"amount"`
	Price  float64 `json:"price"`
}

type CurrencyPrice struct {
	Old     float64 `json:"old"`
	Current float64 `json:"current"`
}

type History struct {
	Amount            int64         `json:"amount"`
	CurrentlyInWallet bool          `json:"currently_in_wallet,omitempty"` // Shows which data is currently in the wallet
	Price             CurrencyPrice `json:"price"`
}

type Read struct {
	Id      string    `json:"id"`
	Code    string    `json:"code"`
	History []History `json:"history"`
}

type List []Read
