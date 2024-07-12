package database

import (
	"fmt"
	"log"

	"go_learning/config"
	"go_learning/internal/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Init() error{
	config := config.LoadConfig();
	log.Print("Connecting To Database:", config.DatabaseHost, ":", config.DatabasePort);
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        config.DatabaseHost, config.DatabasePort, config.DatabaseUser, config.DatabasePassword, config.DatabaseName);

	db,err := gorm.Open("postgres",psqlInfo);
	if err != nil {
		log.Print("Failed to connect Database");
		return err
	}
	MigrateDB(db);
	DB = db;
	return nil
}	

func MigrateDB(db *gorm.DB){
	db.AutoMigrate(&models.User{});
}