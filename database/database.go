package database

import (
	"database/sql"
	"fmt"
	"log"

	"go_learning/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init(){
	config := config.LoadConfig();
	log.Print("Connected To Database:", config.DatabaseHost, ":", config.DatabasePort);
	var err error
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        config.DatabaseHost, config.DatabasePort, config.DatabaseUser, config.DatabasePassword, config.DatabaseName);

	DB,err = sql.Open("postgres",psqlInfo);
	if err != nil {
		log.Fatal("Failed to connect database",err);
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping database",err);
	}
}