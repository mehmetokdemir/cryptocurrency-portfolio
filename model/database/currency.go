package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type History struct {
	Amount int64   `json:"amount" bson:"amount"` // Amount of the code
	Price  float64 `json:"price" bson:"price"`   // Multiply amount and usd value of code
}

type Currency struct {
	Id      primitive.ObjectID `json:"id" bson:"_id,omitempty"` // ID of the record
	Code    string             `json:"code" bson:"code"`        // Code of the currency
	Amount  int64              `json:"amount" bson:"amount"`    // Amount of the code
	Price   float64            `json:"price" bson:"price"`      // Multiply amount and usd value of code
	History []History          `json:"history" bson:"history"`  //
}

type Currencies []Currency
