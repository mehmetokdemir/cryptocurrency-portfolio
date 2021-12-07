package handler

import (
	"context"
	. "cryptocurrency-portfolio/model/common"
	"cryptocurrency-portfolio/model/database"
	"cryptocurrency-portfolio/model/response"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func (h *Handler) GetCurrencies() ApiResponse {
	ctx := context.TODO()
	cur, err := h.MongoCollection.Find(ctx, bson.M{})
	if err != nil {
		// TODO: ERR
		return GenerateResponse(1, "", nil)
	}

	var cryptoCurrencies database.Currencies
	if err := cur.All(ctx, &cryptoCurrencies); err != nil {
		// TODO: ERR
		return GenerateResponse(1, "", nil)
	}

	var respondCurrencies response.List
	for _, cryptoCurrency := range cryptoCurrencies {
		var history []response.History
		availableCurrentPrice, err := calculatePrice(cryptoCurrency.Amount, cryptoCurrency.Code)
		if err != nil {
			return GenerateResponse(http.StatusInternalServerError, DescriptionEnumServerError, err.Error())
		}

		history = append(history, response.History{
			Amount:            cryptoCurrency.Amount,
			CurrentlyInWallet: true,
			Price: response.CurrencyPrice{
				Old:     cryptoCurrency.Price,
				Current: availableCurrentPrice,
			},
		})

		for _, past := range cryptoCurrency.History {
			price, err := calculatePrice(cryptoCurrency.Amount, cryptoCurrency.Code)
			if err != nil {
				return GenerateResponse(http.StatusInternalServerError, DescriptionEnumServerError, err.Error())
			}
			history = append(history, response.History{
				Amount: past.Amount,
				Price: response.CurrencyPrice{
					Old:     past.Price,
					Current: price,
				},
			})
		}

		respondCurrencies = append(respondCurrencies, response.Read{
			Id:      cryptoCurrency.Id.Hex(),
			Code:    cryptoCurrency.Code,
			History: history,
		})
	}
	return GenerateResponse(http.StatusOK, DescriptionEnumSuccess, respondCurrencies)
}