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

	r.POST("/signupcustomer", handlers.CustomerSignup)
	r.POST("/signincustomer",handlers.CustomerSignIn)
	r.POST("/signupdoctor", handlers.DoctorSignup)
	r.POST("/signindoctor",handlers.DoctorSignin)
}
