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

func CreateUser(user *models.User) (*primitive.ObjectID, error) {
	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	user.ID = primitive.NewObjectID()

	result, err := client.Database("last_try").Collection("user").InsertOne(ctx, user)
	if err != nil {
		log.Printf("Could not create trickTips: %v", err)
		return nil, err
	}
	oid := result.InsertedID.(primitive.ObjectID)

	return &oid, nil
}

func GetAllUser(query *models.Query) ([]*models.User, error) {
	var users []*models.User

	findOptions := options.Find()
	findOptions.SetSort(query.Sort)

	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database("last_try")
	collection := db.Collection("user")
	cursor, err := collection.Find(ctx, query.Filters, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &users)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}

	return users, nil
}

func GetUserByID(id primitive.ObjectID) (*models.User, error) {
	user := &models.User{}

	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("last_try").Collection("user").FindOne(ctx, bson.M{"_id": id}).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(user *models.User) error {
	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("last_try").Collection("user").UpdateOne(ctx,
		bson.M{
			"_id": user.ID,
		},
		bson.D{
			{"$set", bson.D{{"username", user.Username}}},
			{"$set", bson.D{{"email", user.Email}}},
			{"$set", bson.D{{"password", user.Password}}},
			{"$set", bson.D{{"avatar", user.Avatar}}},
			{"$set", bson.D{{"stance", user.Stance}}},
		})

	if err != nil {
		return errors.New("Couldn't update user")
	}

	return nil
}
