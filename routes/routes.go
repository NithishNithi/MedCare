package routes

import (
	"medcare/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Static("/signin", "./frontend/signin")
	r.Static("/home", "./frontend/home")
	r.Static("/signup", "./frontend/signup")
	r.Static("/doctorsignup", "./frontend/doctorsignup")
	r.Static("/doctorsignin", "./frontend/doctorsignin")
	r.Static("/doctorhome", "./frontend/doctorhome")
	r.Static("/patienthome", "./frontend/patienthome")
	r.Static("/patientcreateappointment", "./frontend/patientcreateappointment")
	r.Static("/doctorenroll", "./frontend/doctorenroll")
	r.Static("/listappointmentfordoctor", "./frontend/doctorlistappointment")
	r.Static("/createprescriptionfordoctor", "./frontend/doctorprescription")
	r.Static("/patientreport", "./frontend/patientreport")
	r.Static("/doctorreport", "./frontend/doctorreport")
	r.Static("/patientprescriptionlist", "./frontend/patientprescription")

	r.POST("/signupcustomer", handlers.CustomerSignup)
	r.POST("/signincustomer", handlers.CustomerSignIn)
	r.POST("/signupdoctor", handlers.DoctorSignup)
	r.POST("/signindoctor", handlers.DoctorSignin)
	r.POST("/bookappointment", handlers.BookAppointment)
	r.POST("/listappointments", handlers.ListAppointment)
	r.POST("/createprescription", handlers.CreatePrescription)
	r.POST("/listpatientreport", handlers.ListPatientReport)
	r.POST("/listpatientprescription", handlers.ListPrescription)
	r.POST("/listblogs", handlers.ListBlogs)
	r.POST("/listcounts", handlers.Count)
}
