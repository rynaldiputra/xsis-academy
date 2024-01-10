package config

import (
	"fmt"
	"log"
	"theater/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func Connect() (*gorm.DB, error) {
	dataConn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		AppConfig.Username,
		AppConfig.Password,
		AppConfig.Host,
		AppConfig.Port,
		AppConfig.DatabaseName,
	)

	db, err := gorm.Open(mysql.Open(dataConn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("Database Connected Successfully")

	Migrate(db)

	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entities.Movie{},
	)

	if err != nil {
		return err
	}

	log.Println("Database Migration Completed...")
	return nil
}
