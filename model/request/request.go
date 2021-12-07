package request

type CreateAndUpdate struct {
	Code   string `json:"code" extensions:"x-order-1" example:"BTC" validate:"required" valid:"required~code|invalid"`   // Symbol of the cryptocurrency
	Amount int64  `json:"amount" extensions:"x-order-2" example:"3" validate:"required" valid:"required~amount|invalid"` // Amount of the code
}
