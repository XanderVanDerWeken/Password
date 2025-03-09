package persistence

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func LoadDatabase(connectionString string) *gorm.DB {
	newDb, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	err = newDb.AutoMigrate(&AccountEntry{})
	if err != nil {
		log.Fatal("failed to migrate database schema")
	}

	return newDb
}
