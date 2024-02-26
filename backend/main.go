package main

import (
	"backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)


func InitDatabase(){
	
}

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	log.Fatal(r.Run(":8080"))
}
