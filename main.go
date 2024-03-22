package main

import (
	"log"
	"medcare/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	log.Fatal(r.Run("0.0.0.0:8080"))
}
