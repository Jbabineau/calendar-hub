package db

import (
	"os"

	"github.com/jbabineau/calendar-hub/internal/store"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func open(dbName string) (*gorm.DB, error) {
	err := os.MkdirAll("/tmp", 0755)
	if err != nil {
		return nil, err
	}

	return gorm.Open(sqlite.Open(dbName), &gorm.Config{})
}

func MustOpen(dbName string) *gorm.DB {
	if dbName == "" {
		dbName = "calendar.db"
	}

	db, err := open(dbName)
	if err != err {
		panic(err)
	}

	err = db.AutoMigrate(&store.User{}, &store.Session{})

	if err != nil {
		panic(err)
	}

	return db
}