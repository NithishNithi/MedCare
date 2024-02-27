package database

import (
	"context"
	"medcare/constants"

	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Customer_Collection *mongo.Collection

func init() {
	clientoption := options.Client().ApplyURI(constants.MongoURI)
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	client, err := mongo.Connect(ctx, clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected-(patient)")
	Customer_Collection = client.Database("MedCare").Collection("CustomerProfile")

	fmt.Println("Collection Connected")
}
