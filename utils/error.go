package utils

import (
	"go.mongodb.org/mongo-driver/bson"
)

func NewError(err error) bson.M {
	return bson.M{"error": err.Error()}
}
