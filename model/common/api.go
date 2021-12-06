package common

type DescriptionEnum string

const (
	DescriptionEnumSuccess              DescriptionEnum = "Success"
	DescriptionEnumBodyError            DescriptionEnum = "Request body or parameters wrong"
	DescriptionEnumCurrencyAlreadyExist DescriptionEnum = "Currency already exist"
	DescriptionEnumCurrencyNotFound     DescriptionEnum = "Currency not found"
	DescriptionEnumServerError          DescriptionEnum = "Server error"

	DescriptionEnumCannotGetCurrencies DescriptionEnum = "Can not get currencies from coin market cap"
)

type ApiResponse struct {
	StatusCode  int             `json:"status_code"`
	Description DescriptionEnum `json:"description"`
	Data        interface{}     `json:"data"`
}

func GenerateResponse(status int, description DescriptionEnum, data interface{}) ApiResponse {
	return ApiResponse{
		StatusCode:  status,
		Description: description,
		Data:        data,
	}
}
