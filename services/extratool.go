package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"medcare/constants"
	"medcare/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mailgun/mailgun-go/v4"
)

func CreateToken(email, customerid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      email,
		"customerid": customerid, // Set the customerid claim
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(constants.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractCustomerID(jwtToken string, secretKey string) (string, error) {
	// Parse the JWT token
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err // Handle token parsing errors
	}

	// Check if the token is valid
	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Extract the customer ID from the claims
			customerID, ok := claims["customerid"].(string)
			if ok {
				return customerID, nil
			}
		}
	}

	return "", fmt.Errorf("invalid or expired JWT token")
}

func GenerateIDPA() string {
	length := 4
	randomBytes := make([]byte, length)
	// Use crypto/rand to generate random bytes
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err) // Handle error appropriately
	}
	// Convert random bytes to a hexadecimal string
	id := hex.EncodeToString(randomBytes)

	return "PAT" + id
}
func GenerateIDDOC() string {
	length := 4
	randomBytes := make([]byte, length)
	// Use crypto/rand to generate random bytes
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err) // Handle error appropriately
	}
	// Convert random bytes to a hexadecimal string
	id := hex.EncodeToString(randomBytes)

	return "DOC" + id
}

func CurrentTime() string {
	currentTime := time.Now()
	currentTimeStr := currentTime.Format("2006-01-02 15:04:05")
	return currentTimeStr
}

// //////////////////////////// Send Email - mailGun //////////////////////////////////
func SendSimpleMessage(request *models.MailGunEmail) (string, error) {
	domain := "sandbox567bd12d5b42484e94273fbe408ababa.mailgun.org"
	mg := mailgun.NewMailgun(domain, constants.Mailgun_apikey)

	sender := "Excited User <mailgun@sandbox567bd12d5b42484e94273fbe408ababa.mailgun.org>"
	subject := request.Subject
	message := request.Message

	m := mg.NewMessage(sender, subject, message, request.RecipientEmail)

	ctx := context.Background()

	// Send the message with the context
	resp, id, err := mg.Send(ctx, m)
	if err != nil {
		return "", err
	}
	fmt.Println("MailGun :", resp)
	return id, nil
}

type MeetLinkResponse struct {
	MeetLink string `json:"meet_link"`
	EventID  string `json:"event_id"`
}

type Event struct {
	Summary        string `json:"summary"`
	Start          Date   `json:"start"`
	End            Date   `json:"end"`
	ConferenceData `json:"conferenceData"`
}

type Date struct {
	Date string `json:"date"`
}

type ConferenceData struct {
	CreateRequest `json:"createRequest"`
}

type CreateRequest struct {
	ConferenceSolutionKey `json:"conferenceSolutionKey"`
	RequestId             string `json:"requestId"`
}

type ConferenceSolutionKey struct {
	Type string `json:"type"`
}

// func GetGmmetLink(request *models.BookAppointment)(string,error) {
// 	url := "http://localhost:5001/create-event"

// 	// Event details
// 	summary := "Medcare Appointment" // Replace with the summary you want to send dynamically

// 	event := Event{
// 		Summary: summary,
// 		Start: Date{
// 			Date: "2024-03-06", // Replace with the desired date
// 		},
// 		End: Date{
// 			Date: "2024-03-06", // Replace with the desired date
// 		},
// 		ConferenceData: ConferenceData{
// 			CreateRequest: CreateRequest{
// 				ConferenceSolutionKey: ConferenceSolutionKey{
// 					Type: "hangoutsMeet",
// 				},
// 				RequestId: "abblah",
// 			},
// 		},
// 	}

// 	// Convert event struct to JSON
// 	eventJSON, err := json.Marshal(event)
// 	if err != nil {
// 		fmt.Println("Error marshalling event to JSON:", err)
// 		return
// 	}

// 	// Make POST request to the microservice with event details
// 	resp, err := http.Post(url, "application/json", bytes.NewBuffer(eventJSON))
// 	if err != nil {
// 		fmt.Println("Error making POST request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Read the response body
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}

// 	// Parse the JSON response
// 	var meetLinkResponse MeetLinkResponse
// 	err = json.Unmarshal(body, &meetLinkResponse)
// 	if err != nil {
// 		fmt.Println("Error parsing JSON response:", err)
// 		return
// 	}

// 	// Print the generated meet link and event ID
// 	fmt.Println("Generated meet link:", meetLinkResponse.MeetLink)
// 	fmt.Println("Event ID:", meetLinkResponse.EventID)
// }
