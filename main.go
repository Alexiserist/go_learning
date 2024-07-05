package main

import (
	"go_learning/config"
	"go_learning/database"
	"go_learning/internal/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main()  {
	config := config.LoadConfig();
	database.Init(config);

	router := routes.LoadRouter();
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler));
	router.Run(":8080");

}