package repository

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"last_try_rest/db"
	"last_try_rest/models"
)

func GetFavorite(userID *primitive.ObjectID, trickTipsID *primitive.ObjectID) (*models.Favorite, error) {
	favorite := &models.Favorite{}

	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("last_try").Collection("favorite").FindOne(ctx,
		bson.M{
			"user_id":      userID,
			"trickTips_id": trickTipsID,
		}, nil).Decode(favorite)

	if err != nil {
		return nil, err
	}

	return favorite, nil
}

func UpdateFavorite(favorite *models.Favorite) error {
	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("last_try").Collection("favorite").UpdateOne(ctx,
		bson.M{
			"user_id":      favorite.UserID,
			"trickTips_id": favorite.TrickTipsID,
		},
		bson.D{
			{"$set", bson.D{{"favorite", favorite.IsFavorite}}},
			{"upsert", true},
		})

	if err != nil {
		return errors.New("Could not update user favorite trickTips")
	}

	return nil
}
