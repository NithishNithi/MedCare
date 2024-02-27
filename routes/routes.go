package routes

import (
	"medcare/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Static("/signin", "/home/nithish/Desktop/medcare/frontend/signin")
	r.Static("/home", "/home/nithish/Desktop/medcare/frontend/home")
	r.Static("/signup", "/home/nithish/Desktop/medcare/frontend/signup")
	r.POST("/registercustomer", handlers.CustomerSignIn)
}
