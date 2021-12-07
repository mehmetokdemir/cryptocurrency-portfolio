package handler

import (
	// Go imports
	"context"
	"net/http"
	// External imports
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	// Internal imports
	. "cryptocurrency-portfolio/model/common"
	"cryptocurrency-portfolio/model/database"
	"cryptocurrency-portfolio/model/response"
)

func (h *Handler) GetCurrencyBy(id string) ApiResponse {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return GenerateResponse(http.StatusBadRequest, DescriptionEnumBodyError, map[string]string{"error": "Bad request"})
	}

	var crypto database.Currency
	if err := h.MongoCollection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&crypto); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return GenerateResponse(http.StatusNotFound, DescriptionEnumCurrencyNotFound, map[string]string{"error": "Currency with that id does not exist"})
		default:
			return GenerateResponse(http.StatusInternalServerError, DescriptionEnumServerError, map[string]string{"error": "Server error"})
		}
	}

	var history []response.History
	// Add available price to history
	availableCurrentPrice, err := calculatePrice(crypto.Amount, crypto.Code)
	if err != nil {
		return GenerateResponse(http.StatusInternalServerError, DescriptionEnumServerError, err.Error())
	}

	history = append(history, response.History{
		Amount:            crypto.Amount,
		CurrentlyInWallet: true,
		Price: response.CurrencyPrice{
			Old:     crypto.Price,
			Current: availableCurrentPrice,
		},
	})

	// Add old prices to history
	for _, past := range crypto.History {
		currentPrice, err := calculatePrice(past.Amount, crypto.Code)
		if err != nil {
			return GenerateResponse(http.StatusNotFound, DescriptionEnumCurrencyNotFound, err.Error())
		}
		history = append(history, response.History{
			Amount: past.Amount,
			Price: response.CurrencyPrice{
				Old:     past.Price,
				Current: currentPrice,
			},
		})
	}

	return GenerateResponse(http.StatusOK, DescriptionEnumSuccess, response.Read{
		Id:      crypto.Id.Hex(),
		Code:    crypto.Code,
		History: history,
	})
}
