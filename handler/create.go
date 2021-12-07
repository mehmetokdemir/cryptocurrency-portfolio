package handler

import (
	// Go imports
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	// External imports
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"github.com/shopspring/decimal"

	// Internal imports
	. "cryptocurrency-portfolio/model/common"
	"cryptocurrency-portfolio/model/database"
	"cryptocurrency-portfolio/model/request"
	"cryptocurrency-portfolio/model/response"
)

// PutCurrency godoc
// @Summary      Create Cryptocurrency
// @Description  Create cryptocurrency portfolio
// @Tags         Cryptocurrency
// @Produce      json
// @Param request body request.CreateAndUpdate true "Example Request"
// @Success      200  {object}  ApiResponse{data=response.CreateAndUpdate} "Success"
// @Router       /currency [put]
func (h *Handler) PutCurrency() ApiResponse {
	body, err := ioutil.ReadAll(h.Ctx.Request().Body)
	if err != nil {
		return GenerateResponse(http.StatusBadRequest, DescriptionEnumBodyError, err.Error())
	}

	var create request.CreateAndUpdate
	if err := json.Unmarshal(body, &create); err != nil {
		return GenerateResponse(http.StatusBadRequest, DescriptionEnumBodyError, err.Error())
	}

	validate := validator(create)
	if len(validate) > 0 {
		return GenerateResponse(http.StatusBadRequest, DescriptionEnumBodyError, validate)
	}

	create.Code = strings.ToUpper(create.Code)

	// TODO: Make local cache to currencies list
	currencyCodes, err := getCurrencyCodes()
	if err != nil {
		return GenerateResponse(
			http.StatusInternalServerError,
			DescriptionEnumServerError,
			map[string]string{"error": err.Error()})
	}

	var currencyInMarketCamp string
	for _, currency := range currencyCodes {
		// funk package can be used for this check
		if currency == create.Code {
			currencyInMarketCamp = create.Code
			break
		}
	}

	if currencyInMarketCamp == "" {
		return GenerateResponse(
			http.StatusInternalServerError,
			DescriptionEnumServerError,
			map[string]string{"error": fmt.Sprintf("currency %s is not found in market camp coin list", create.Code)})
	}

	count, err := h.MongoCollection.CountDocuments(context.TODO(), bson.M{"code": currencyInMarketCamp})
	if err != nil {
		return GenerateResponse(http.StatusInternalServerError, DescriptionEnumServerError, map[string]string{"error": "Server error"})
	}

	if count > 0 {
		return GenerateResponse(http.StatusForbidden, DescriptionEnumCurrencyAlreadyExist, map[string]string{"error": "Currency already exist"})
	}

	price, err := calculatePrice(create.Amount, create.Code)
	if err != nil {
		return GenerateResponse(http.StatusNotFound, DescriptionEnumCurrencyNotFound, map[string]string{"error": err.Error()})
	}

	cryptoCurrency := database.Currency{
		Code:   currencyInMarketCamp,
		Amount: create.Amount,
		Price:  price,
	}

	insertResult, err := h.MongoCollection.InsertOne(context.TODO(), cryptoCurrency)
	if err != nil {
		return GenerateResponse(http.StatusInternalServerError, DescriptionEnumServerError, map[string]string{"error": "Server error"})
	}

	var insertedId string
	if id, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		insertedId = id.Hex()
	}

	return GenerateResponse(http.StatusOK, DescriptionEnumSuccess, response.CreateAndUpdate{
		Id:     insertedId,
		Code:   currencyInMarketCamp,
		Amount: create.Amount,
		Price:  price,
	})
}
