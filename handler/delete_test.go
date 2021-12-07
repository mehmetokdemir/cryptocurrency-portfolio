package handler

import (
	// Go imports
	"context"
	"testing"

	// External imports
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func deleteData(collection collectionApi) *mongo.SingleResult {
	if err := collection.FindOneAndDelete(context.Background(), nil).Err(); err != nil {
		return nil
	}
	return nil
}

func (m *mockCollection) FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
	c := &mongo.SingleResult{}
	return c
}

func TestDeleteCurrency(t *testing.T) {
	res := deleteData(mockCol)
	//assert.Nil(t, err)
	assert.IsType(t, &mongo.SingleResult{}, res)
}
