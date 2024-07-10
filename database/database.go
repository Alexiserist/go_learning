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

func Init(){
	config := config.LoadConfig();
	log.Print("Connected To Database:", config.DatabaseHost, ":", config.DatabasePort);
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        config.DatabaseHost, config.DatabasePort, config.DatabaseUser, config.DatabasePassword, config.DatabaseName);

	db,err := gorm.Open("postgres",psqlInfo);
	if err != nil {
		panic("failed to connect db")
	}
	MigrateDB(db);

	DB = db
}	

func MigrateDB(db *gorm.DB){
	db.AutoMigrate(&models.User{});
}