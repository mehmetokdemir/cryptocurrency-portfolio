package handler

import (
	// Go imports
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	// External imports
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	// Internal imports
	. "cryptocurrency-portfolio/model/common"
	"cryptocurrency-portfolio/model/database"
	"cryptocurrency-portfolio/model/request"
	"cryptocurrency-portfolio/model/response"
)

// PatchCurrencyBy godoc
// @Summary      Update Cryptocurrency
// @Description  Update cryptocurrency portfolio by id
// @Tags         Cryptocurrency
// @Produce      json
// @Param id path string true "Cryptocurrency id"
// @Param request body request.CreateAndUpdate true "Example Request"
// @Success      200  {object}  ApiResponse "Success"
// @Router       /currency/{id} [patch]
func (h *Handler) PatchCurrencyBy(id string) ApiResponse {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return GenerateResponse(http.StatusBadRequest, DescriptionEnumBodyError, map[string]string{"error": "Bad request"})
	}

	body, err := ioutil.ReadAll(h.Ctx.Request().Body)
	if err != nil {
		return GenerateResponse(http.StatusBadRequest, DescriptionEnumBodyError, err.Error())
	}

	var update request.CreateAndUpdate
	if err := json.Unmarshal(body, &update); err != nil {
		return GenerateResponse(http.StatusBadRequest, DescriptionEnumBodyError, err.Error())
	}

	validate := validator(update)
	if len(validate) > 0 {
		return GenerateResponse(http.StatusBadRequest, DescriptionEnumBodyError, validate)
	}

	var cryptoCurrency database.Currency
	if err := h.MongoCollection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&cryptoCurrency); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return GenerateResponse(http.StatusNotFound, DescriptionEnumCurrencyNotFound, map[string]string{"error": "Currency with that id does not exist"})
		default:
			return GenerateResponse(http.StatusInternalServerError, DescriptionEnumServerError, map[string]string{"error": "Server error"})
		}
	}

	if strings.ToUpper(update.Code) != cryptoCurrency.Code {
		// Check existing currency and requested currency, if are not equal return error
		return GenerateResponse(http.StatusForbidden, DescriptionEnumCurrencyNotFound, map[string]string{"error": "Current value mismatch with entered value"})
	}

	cryptoCurrency.History = append(cryptoCurrency.History, database.History{
		Amount: cryptoCurrency.Amount,
		Price:  cryptoCurrency.Price,
	})

	price, err := calculatePrice(update.Amount, cryptoCurrency.Code)
	if err != nil {
		return GenerateResponse(http.StatusNotFound, DescriptionEnumCurrencyNotFound, map[string]string{"error": err.Error()})
	}

	cryptoCurrency.Amount = update.Amount
	cryptoCurrency.Price = price

	if _, err := h.MongoCollection.ReplaceOne(context.TODO(), bson.M{"_id": objectId}, cryptoCurrency); err != nil {
		return GenerateResponse(http.StatusInternalServerError, DescriptionEnumServerError, map[string]string{"error": "Server error"})
	}

	return GenerateResponse(http.StatusOK, DescriptionEnumSuccess, response.CreateAndUpdate{
		Id:     id,
		Code:   cryptoCurrency.Code,
		Amount: update.Amount,
		Price:  price,
	})
}
