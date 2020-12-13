package repository

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"last_try_rest/db"
	"last_try_rest/models"
	"log"
)

func CreateTrickTips(trickTips *models.TrickTips) (*primitive.ObjectID, error) {
	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	trickTips.ID = primitive.NewObjectID()

	result, err := client.Database("last_try").Collection("trickTips").InsertOne(ctx, trickTips)
	if err != nil {
		log.Printf("Could not create trickTips: %v", err)
		return nil, err
	}
	oid := result.InsertedID.(primitive.ObjectID)

	return &oid, nil
}

func GetTrickTipsByID(id primitive.ObjectID) (*models.TrickTips, error) {
	trickTips := &models.TrickTips{}

	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("last_try").Collection("trickTips").FindOne(ctx, bson.M{"_id": id}).Decode(trickTips)
	if err != nil {
		return nil, err
	}

	return trickTips, nil
}

func GetAllTrickTips(query *models.Query) ([]*models.TrickTips, error) {
	var trickTips []*models.TrickTips

	findOptions := options.Find()
	findOptions.SetSort(query.Sort)

	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database("last_try")
	collection := db.Collection("trickTips")
	cursor, err := collection.Find(ctx, query.Filters, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &trickTips)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}

	return trickTips, nil
}

func UpdateTrickTips(trickTips *models.TrickTips) error {
	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("last_try").Collection("trickTips").UpdateOne(ctx,
		bson.M{
			"_id": trickTips.ID,
		},
		bson.D{
			{"$set", bson.D{{"title", trickTips.Title}}},
			{"$set", bson.D{{"thumbnail", trickTips.Thumbnail}}},
			{"$set", bson.D{{"url", trickTips.Url}}},
			{"$set", bson.D{{"category", trickTips.Category}}},
			{"$set", bson.D{{"level", trickTips.Level}}},
			{"$set", bson.D{{"sequence", trickTips.Sequence}}},
			{"$set", bson.D{{"descriptionStep", trickTips.DescriptionStep}}},
		})

	if err != nil {
		return errors.New("Could not update tricksTips")
	}

	return nil
}
