package store

import (
	"context"
	"url-shortener/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func getVisitCollection() *mongo.Collection {
	return DB.Collection("visits")
}

func SaveVisit(visit models.Visit) error {
	_, err := getVisitCollection().InsertOne(context.Background(), visit)
	return err
}

func GetVisitsByID(id string) ([]models.Visit, error) {
	var visits []models.Visit

	cursor, err := getVisitCollection().Find(context.Background(), bson.M{"id": id})
	if err != nil {
		return visits, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var visit models.Visit
		cursor.Decode(&visit)
		visits = append(visits, visit)
	}

	return visits, nil
}