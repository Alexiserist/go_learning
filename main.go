package main

import (
	"go_learning/database"
	"go_learning/internal/routes"
	_ "go_learning/docs"
)


func main()  {
 	err := database.Init();if err != nil {
		return;
	}
	router := routes.LoadRouter();
	router.Run(":8080");

}