package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Favorite struct {
	ID          primitive.ObjectID  `json:"id" bson:"_id"`
	UserID      *primitive.ObjectID `json:"user_id" binding:"required" bson:"user_id"`
	TrickTipsID *primitive.ObjectID `json:"trickTips_id" binding:"required" bson:"trickTips_id"`
	IsFavorite  bool                `json:"isFavorite" bson:"isFavorite"`
}

type FavoriteForm struct {
	UserID      *primitive.ObjectID `json:"user_id" binding:"required" bson:"user_id"`
	TrickTipsID *primitive.ObjectID `json:"trickTips_id" binding:"required" bson:"trickTips_id"`
}
