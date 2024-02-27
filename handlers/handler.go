package handlers

import (
	"fmt"
	"log"
	"medcare/models"
	"medcare/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomerSignIn(c *gin.Context) {
	var request *models.CustomerSignUp
	err := c.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println(err)
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
	}

	err = services.CustomerSignUp(request)
	if err != nil {
		fmt.Println(err)
		log.Println("Error creating customer:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": "Account Created"})
}
