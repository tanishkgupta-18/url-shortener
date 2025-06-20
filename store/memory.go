package store

import (
	"context"
	"url-shortener/models"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func getURLCollection() *mongo.Collection {
	return DB.Collection("urls")
}

func SaveURL(url models.URL) error {
	_, err := getURLCollection().InsertOne(context.Background(), url)
	return err
}

func GetURLByID(id string) (models.URL, error) {
	var result models.URL
	err := getURLCollection().FindOne(context.Background(), bson.M{"id": id}).Decode(&result)
	return result, err
}