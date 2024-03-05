package services

import (
	"context"
	"errors"
	"log"
	"medcare/database"
	"medcare/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CustomerSignUp(request *models.CustomerSignUp) error {
	ctx := context.Background()
	request.CustomerID = GenerateIDPA()
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"customerid", request.CustomerID}},
			bson.D{{"email", request.EmailID}},
		}},
	}

	existingCustomer := models.CustomerSignUp{}
	err := database.Customer_Collection.FindOne(ctx, filter).Decode(&existingCustomer)
	if err == nil {
		log.Println("Account Already Exists", err)
		return errors.New("account elready Exists")
	} else if err != mongo.ErrNoDocuments {
		return err
	}
	request.CreatedTime = CurrentTime()

	result, err := database.Customer_Collection.InsertOne(ctx, request)
	if err != nil {
		log.Printf("Failed to create customer: %v", err)
		return errors.New("failed to create customer")
	}
	// Find and return the newly inserted customer
	newCustomer := &models.CustomerSignUp{}
	err = database.Customer_Collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(newCustomer)
	if err != nil {
		return err
	}

	return nil
}

func IsValidUser(request *models.CustomerSignIn) (bool, models.CustomerSignUp) {
	ctx := context.Background()
	customer := models.CustomerSignUp{}
	query := bson.M{"emailid": request.EmailId}
	err := database.Customer_Collection.FindOne(ctx, query).Decode(&customer)
	if err != nil {
		log.Println("Error in fetching user details : ", err)
		return false, customer
	}
	if request.Password != customer.Password {
		return false, customer
	}
	return true, customer
}

func InsertToken(customerid, email, token string) (string, error) {
	dbToken := models.PatientToken{EmailID: email, Token: token, CustomerID: customerid}
	result, err := database.Customer_Token.InsertOne(context.Background(), dbToken)
	if err != nil {
		log.Printf("Error inserting token: %v", err)
		return "", err
	}
	token1 := models.PatientToken{}
	query := bson.M{"_id": result.InsertedID}
	err = database.Customer_Token.FindOne(context.Background(), query).Decode(&token1)
	if err != nil {
		log.Println("Error in fetching token : ", err)
		return "", err
	}
	return token1.Token, nil
}

func BookAppointment(request *models.BookAppointment) error {
	ctx := context.Background()
	_, err := database.BookAppointment.InsertOne(ctx, request)
	if err != nil {
		log.Println("error while inserting  patient appointment", err)
		return err
	}
	return nil
}




