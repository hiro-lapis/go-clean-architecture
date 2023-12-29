package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB create db instance
func NewDB() *gorm.DB {
	if os.Getenv(("GO_ENV")) == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTRES_USER"),
		os.Getenv("POSTRES_PW"),
		os.Getenv("POSTRES_HOST"),
		os.Getenv("POSTRES_PORT"),
		os.Getenv("POSTRES_DB"),
	)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	// retrive DB connection instance by calling gorm instance
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
