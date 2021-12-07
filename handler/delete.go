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
)

// DeleteCurrencyBy godoc
// @Summary      Delete Cryptocurrency
// @Description  Delete cryptocurrency portfolio
// @Tags         Cryptocurrency
// @Produce      json
// @Param id path string true "Cryptocurrency id"
// @Success      200  {object}  ApiResponse "Success"
// @Router       /currency/{id} [delete]
func (h *Handler) DeleteCurrencyBy(id string) ApiResponse {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return GenerateResponse(http.StatusBadRequest, DescriptionEnumBodyError, map[string]string{"error": "Bad request"})
	}

	if err := h.MongoCollection.FindOneAndDelete(context.TODO(), bson.M{"_id": objectId}).Err(); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return GenerateResponse(http.StatusNotFound, DescriptionEnumCurrencyNotFound, map[string]string{"error": "Currency with that id does not exist"})
		default:
			return GenerateResponse(http.StatusInternalServerError, DescriptionEnumServerError, map[string]string{"error": "Server error"})
		}
	}

	return GenerateResponse(http.StatusOK, DescriptionEnumSuccess, nil)
}
