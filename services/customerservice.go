package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"medcare/database"
	"medcare/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CustomerSignUp(request *models.CustomerSignUp) error {
	ctx := context.Background()
	request.CustomerID = GenerateID()
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"customerid", request.CustomerID}},
			bson.D{{"email", request.EmailID}},
		}},
	}

	existingCustomer := models.CustomerSignUp{}
	err := database.Customer_Collection.FindOne(ctx, filter).Decode(&existingCustomer)
	if err == nil {
		fmt.Println(err)
		return errors.New("Account Already Exists")
	} else if err != mongo.ErrNoDocuments {
		return err
	}
	request.CreatedTime = CurrentTime()

	result, err := database.Customer_Collection.InsertOne(ctx, request)
	if err != nil {
		log.Printf("Failed to create customer: %v", err)
		return errors.New("Failed to create customer")
	}
	// Find and return the newly inserted customer
	newCustomer := &models.CustomerSignUp{}
	err = database.Customer_Collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(newCustomer)
	if err != nil {
		return err
	}

	return nil
}
