package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	dsn := fmt.Sprintf()
}