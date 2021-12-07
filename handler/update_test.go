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

func (m *mockCollection) ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	c := &mongo.UpdateResult{}
	return c, nil
}

func updateData(collection collectionApi, currency database.Currency) (*mongo.UpdateResult, error) {
	res, err := collection.ReplaceOne(context.Background(), nil, currency)
	if err != nil {
		return res, err
	}
	return res, nil
}

func TestUpdateCurrency(t *testing.T) {
	res, err := updateData(mockCol, database.Currency{
		Code:   "OKC",
		Amount: 5,
	})
	assert.Nil(t, err)
	assert.IsType(t, &mongo.UpdateResult{}, res)
}
