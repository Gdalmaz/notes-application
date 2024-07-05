package database

import (
	"fmt"
	"log"
	"notes-application/config"
	"notes-application/models"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

func Connect() {
	p := config.Config("POSTGRES_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Fatal("error parsing")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.Config("POSTGRES_HOST"), config.Config("POSTGRES_USER"),
		config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_DB"), port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	log.Println("Database connection established successfully")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")
	db.AutoMigrate(&models.User{}, &models.Session{}, &models.Notes{})
	DB = DBInstance{
		Db: db,
	}
}
