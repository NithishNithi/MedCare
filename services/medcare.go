package services

import (
	"context"
	"log"
	"medcare/database"
	"medcare/models"

	"go.mongodb.org/mongo-driver/bson"
)

func ListBlogs() ([]models.ListBlogs, error) {
	ctx := context.Background()
	cursor, err := database.BlogCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var results []models.ListBlogs
	err = cursor.All(ctx, &results)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return results, nil
}
