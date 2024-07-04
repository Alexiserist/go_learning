package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

type Config struct {
    DatabaseUser     string
    DatabasePassword string
    DatabaseHost     string
    DatabasePort     string
    DatabaseName     string
}

func LoadConfig() Config {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return Config{
        DatabaseUser:     os.Getenv("DATABASE_USER"),
        DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
        DatabaseHost:     os.Getenv("DATABASE_HOST"),
        DatabasePort:     os.Getenv("DATABASE_PORT"),
        DatabaseName:     os.Getenv("DATABASE_NAME"),
    }
}