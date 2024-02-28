package handlers

import (
	"log"
	"medcare/models"
	"medcare/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoctorSignup(c *gin.Context) {
	var request *models.DoctorSignup
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal server error"})
		return
	}

	err = services.DoctorSignup(request)
	if err != nil {
		log.Println("Error creating customer:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": "Account Created"})
}

func DoctorSignin(c *gin.Context) {
	var request *models.DoctorSignin
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": "Internal Server error"})
		return
	}
	status, ans := services.IsValidUserDoctor(request)
	if status {
		token, err := services.CreateToken(request.EmailId, ans.DoctorID)
		if err != nil {
			log.Println("Error creating token:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		response, err := services.InsertTokenDoctor(ans.DoctorID, request.EmailId, token)
		if err != nil {
			log.Println("Error inserting token:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": response})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
	}
}
