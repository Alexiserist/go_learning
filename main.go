package main

import (
	"go_learning/database"
	"go_learning/internal/routes"
	_ "go_learning/docs"
)


func main()  {
	database.Init();
	router := routes.LoadRouter();
	router.Run(":8080");

}