package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"last_try_rest/models"
	"log"
	"time"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s"
)

// GetConnection - Retrieves a client to the DocumentDB
func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	host := "localhost"
	port := "27017"

	connectionURI := fmt.Sprintf(connectionStringTemplate, host, port)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, ctx, cancel
}

func AddTrickTips(trickTips *models.TrickTips) (*primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
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

func GetAllTrickTips() ([]*models.TrickTips, error) {
	var trickTips []*models.TrickTips

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database("last_try")
	collection := db.Collection("trickTips")
	cursor, err := collection.Find(ctx, bson.D{})
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
