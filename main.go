package main

import (
	"go_learning/database"
	_ "go_learning/docs"
	"go_learning/internal/routes"
	"log"
)

// @title           Golang Framework GIN Swagger
// @version         1.0
// @description     Golang Framwork GIN ,Database Postgress, Gorm, Swaggo

// @host      localhost:8080

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main()  {
 	err := database.Init();
	 if err != nil {
		log.Fatalf("Could not initialize the database: %v", err)
		return
	}
	router := routes.LoadRouter();
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not start the server: %v", err)
	}

}