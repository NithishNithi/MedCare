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
var Customer_Token *mongo.Collection
var Doctor_Collection *mongo.Collection
var Doctor_Token *mongo.Collection

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

func init() {
	clientoption := options.Client().ApplyURI(constants.MongoURI)
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	client, err := mongo.Connect(ctx, clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected-(patient)")
	Customer_Token = client.Database("MedCare").Collection("CustomerProfile-JWT-Token")

	fmt.Println("Collection Connected")
}

// =================================================================
// --------------->   Doctor   <----------------------
// =================================================================

func init() {
	clientoption := options.Client().ApplyURI(constants.MongoURI)
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	client, err := mongo.Connect(ctx, clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected-(patient)")
	Doctor_Collection = client.Database("MedCare").Collection("DoctorProfile")

	fmt.Println("Collection Connected")
}

func init() {
	clientoption := options.Client().ApplyURI(constants.MongoURI)
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	client, err := mongo.Connect(ctx, clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected-(patient)")
	Doctor_Token = client.Database("MedCare").Collection("DoctorProfile-JWT-Token")

	fmt.Println("Collection Connected")
}
