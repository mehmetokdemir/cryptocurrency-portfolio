package request

type CreateAndUpdate struct {
	Code   string `json:"code" validate:"required" valid:"required~code|invalid"`     // Symbol of the cryptocurrency
	Amount int64  `json:"amount" validate:"required" valid:"required~amount|invalid"` // Amount of the code
}
