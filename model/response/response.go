package response

type CreateAndUpdate struct {
	Id     string  `json:"id" extensions:"x-order-1" example:"61ae85f3b45c25aa9cdaba99"` // ID of the cryptocurrency portfolio
	Code   string  `json:"code" extensions:"x-order-2" example:"BTC"`                    // Code of the cryptocurrency portfolio
	Amount int64   `json:"amount" extensions:"x-order-3" example:"4"`                    // Amount of the code
	Price  float64 `json:"price" extensions:"x-order-4" example:"4900.01"`               // Price of the code and multiplied amount
}

type CurrencyPrice struct {
	Old     float64 `json:"old" extensions:"x-order-1" example:"4850.13"`    // Old price of the currency
	Current float64 `json:"current"extensions:"x-order-2" example:"4900.01"` // New price of the currency
}

type History struct {
	Amount            int64         `json:"amount" extensions:"x-order-1" example:"4900.01"`                     // Amount of the history
	CurrentlyInWallet bool          `json:"currently_in_wallet,omitempty" extensions:"x-order-2" example:"true"` // Shows which data is currently in the wallet
	Price             CurrencyPrice `json:"price" extensions:"x-order-3" example:"4900.01"`                      // Price of the history
}

type Read struct {
	Id      string    `json:"id" extensions:"x-order-1" example:"61ae85f3b45c25aa9cdaba99"` // ID of the cryptocurrency portfolio
	Code    string    `json:"code" extensions:"x-order-2" example:"BTC"`                    // Code of the cryptocurrency portfolio
	History []History `json:"history" extensions:"x-order-3"`                               // History of the portfolio
}

type List []Read
