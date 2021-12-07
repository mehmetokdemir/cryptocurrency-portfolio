package handler

import (
	// Go imports
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	// External imports
	"github.com/asaskevich/govalidator"
	"github.com/go-resty/resty/v2"
	"github.com/kataras/iris/v12"
	"github.com/pterm/pterm"
	"go.mongodb.org/mongo-driver/mongo"

	// Internal imports
	"cryptocurrency-portfolio/model/service"
)

type Handler struct {
	Ctx             iris.Context
	MongoCollection *mongo.Collection
}

const (
	cmcApiKey            = "bf7a2e3b-3fd0-4b8b-8e10-609673e3cd33"
	currencyMapEndpoint  = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/map"
	quotesLatestEndpoint = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest"
)

func getCurrencyCodes() ([]string, error) {
	// https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest
	// "X-CMC_PRO_API_KEY" = bf7a2e3b-3fd0-4b8b-8e10-609673e3cd33

	restyClient := resty.New()
	outResponse, err := restyClient.R().
		SetHeader("Accepts", "application/json").
		SetHeader("X-CMC_PRO_API_KEY", cmcApiKey).
		Get(currencyMapEndpoint)
	if err != nil {
		pterm.Error.Println("can not call cryptocurrency listing", err.Error())
		return nil, errors.New("server error")
	}

	var out service.Data
	if err := json.Unmarshal(outResponse.Body(), &out); err != nil {
		return nil, errors.New("can not decode data")
	}

	var coins []string
	if out.Status.ErrorCode == 0 {
		for _, code := range out.Data {
			coins = append(coins, code.Symbol)
		}
	}

	return coins, nil
}

func calculatePrice(amount int64, code string) (float64, error) {
	restyClient := resty.New()
	outResponse, err := restyClient.R().
		SetHeader("Accepts", "application/json").
		SetHeader("X-CMC_PRO_API_KEY", cmcApiKey).
		SetQueryParams(map[string]string{
			"symbol":  code,
			"convert": "USD",
		}).
		Get(quotesLatestEndpoint)
	if err != nil {
		pterm.Error.Println("can not call cryptocurrency listing", err.Error())
		return 0, errors.New("server error")
	}

	var out service.ListingLatest
	if err := json.Unmarshal(outResponse.Body(), &out); err != nil {
		pterm.Error.Println("err", err.Error())
		return 0, errors.New("can not decode data")
	}

	data, ok := out.Data[code]
	if !ok {
		return 0, fmt.Errorf("crypto currency %s not found", code)
	}

	quote, ok := data.Quote["USD"]
	if !ok {
		return 0, errors.New("usd currency not found")
	}

	return convertPrice(quote.Price, float64(amount)), nil
}

func convertPrice(price, amount float64) float64 {
	convertedPrice := price * amount
	return math.Round(math.Pow(10, float64(2))*convertedPrice) / math.Pow(10, float64(2))
}

func listingsHistorical() map[string]float64 {
	//var currencies = make(map[string]float64)
	//In order to prevent many requests to the coinmarketcap API in the for loop, I wanted to pull this api and map the symbol and quantity. But I couldn't dwell on it because of pagination
	a, _ := getCurrencyCodes()
	l := len(a)
	restyClient := resty.New()
	outResponse, err := restyClient.R().
		SetHeader("Accepts", "application/json").
		SetHeader("X-CMC_PRO_API_KEY", "bf7a2e3b-3fd0-4b8b-8e10-609673e3cd33").
		SetQueryParam("limit", strconv.FormatInt(time.Now().Unix(), l)).
		Get("https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest")
	if err != nil {
		pterm.Error.Println("can not call cryptocurrency listing", err.Error())
		return nil
	}

	fmt.Println("outres", len(outResponse.Body()))

	return nil
}

func validator(data interface{}) map[string]string {
	var validateError = make(map[string]string)
	if _, err := govalidator.ValidateStruct(data); err != nil {
		switch errs := err.(type) {
		case govalidator.Errors:
			for _, e := range errs {
				parts := strings.Split(e.Error(), "|")
				if len(parts) != 2 {
					continue
				}
				validateError[parts[0]] = parts[1]
			}
		}
	}
	return validateError
}
