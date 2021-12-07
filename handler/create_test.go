package handler

import (
	// Go imports
	"context"
	"testing"

	// External imports
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// Internal imports
	"cryptocurrency-portfolio/model/database"
)

func (m *mockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	c := &mongo.InsertOneResult{}
	return c, nil
}

func insertData(collection collectionApi, currency database.Currency) (*mongo.InsertOneResult, error) {
	res, err := collection.InsertOne(context.Background(), currency)
	if err != nil {
		return res, err
	}
	return res, nil
}

func TestCreateCurrency(t *testing.T) {
	res, err := insertData(mockCol, database.Currency{
		Code:   "OKC",
		Amount: 3,
	})
	assert.Nil(t, err)
	assert.IsType(t, &mongo.InsertOneResult{}, res)
}
