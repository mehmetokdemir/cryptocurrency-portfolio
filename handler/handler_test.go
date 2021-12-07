package handler

import (
	// Go imports
	"testing"

	// External imports
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestConvertPrice(t *testing.T) {
	expected := 153834.14
	calculated := convertPrice(51278.04789, 3)
	if expected != calculated {
		t.Errorf("Test Fail : Calculated [%f]\tExptected [%f]\n", calculated, expected)
	}
}

func TestPrimitiveObjectId(t *testing.T) {
	id := "61af5eb8d5c4a5ec47631243"
	calculated := true
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		calculated = false
	}
	if objectId.Hex() != id {
		t.Errorf("Test Fail : Calculated [%t]\tExptected [%t]\n", calculated, true)
	}
}
