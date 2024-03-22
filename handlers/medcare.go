package handlers

import (
	"medcare/models"
	"medcare/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListBlogs(c *gin.Context) {
	response, err := services.ListBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal Server Error"})
	}
	c.JSON(http.StatusAccepted, gin.H{"message": response})
}

func Count(c *gin.Context) {
	patientcount, err := services.CountPatient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	doctorcount, err := services.CountDoctors()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	response := models.Count{PatientCount: patientcount, DoctorCount: doctorcount}
	c.JSON(http.StatusAccepted, gin.H{"message": response})

}
