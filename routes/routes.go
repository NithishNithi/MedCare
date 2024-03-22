package routes

import (
	"medcare/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Static("/signin", "/home/nithish/Desktop/medcare/frontend/signin")
	r.Static("/home", "/home/nithish/Desktop/medcare/frontend/home")
	r.Static("/signup", "/home/nithish/Desktop/medcare/frontend/signup")
	r.Static("/doctorsignup", "/home/nithish/Desktop/medcare/frontend/doctorsignup")
	r.Static("/doctorsignin", "/home/nithish/Desktop/medcare/frontend/doctorsignin")
	r.Static("/doctorhome", "/home/nithish/Desktop/medcare/frontend/doctorhome")
	r.Static("/patienthome", "/home/nithish/Desktop/medcare/frontend/patienthome")
	r.Static("/patientcreateappointment", "/home/nithish/Desktop/medcare/frontend/patientcreateappointment")
	r.Static("/doctorenroll", "/home/nithish/Desktop/medcare/frontend/doctorenroll")
	r.Static("/listappointmentfordoctor", "/home/nithish/Desktop/medcare/frontend/doctorlistappointment")
	r.Static("/createprescriptionfordoctor", "/home/nithish/Desktop/medcare/frontend/doctorprescription")
	r.Static("/patientreport", "/home/nithish/Desktop/medcare/frontend/patientreport")
	r.Static("/doctorreport", "/home/nithish/Desktop/medcare/frontend/doctorreport")
	r.Static("/patientprescriptionlist", "/home/nithish/Desktop/medcare/frontend/patientprescription")

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
