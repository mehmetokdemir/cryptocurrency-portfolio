package handler

import (
	"context"
	"cryptocurrency-portfolio/model/database"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func (h *Handler) TestAddCurrency(t *testing.T) {
	fmt.Println("girdi")
	id := primitive.NewObjectID()
	newCurrency := database.Currency{
		Id:      id,
		Code:    "IKZ",
		Amount:  2,
		Price:   100,
		History: nil,
	}

	h.MongoCollection.InsertOne(context.TODO(), newCurrency)

	retrievedCurrency := database.Currency{}
	err := h.MongoCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&retrievedCurrency)
	assert.Nil(t, err)
	assert.EqualValues(t, newCurrency.Code, retrievedCurrency.Code)
}

func TestCalculatePrice(t *testing.T) {

}
