package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //necessary for gorm
	"github.com/joaoprodrigo/shlink-go/config"
)

// DB is the global Db for use with application
var DB *gorm.DB

// StartupDB runs routines necessary for the DB initialization
func StartupDB() {

	db, err := gorm.Open("sqlite3", config.DatabaseFile)

	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	runMigrations()
}

func runMigrations() {

	DB.AutoMigrate(&ShortURL{})
	DB.AutoMigrate(&Domain{})
	DB.AutoMigrate(&Tag{})
	DB.AutoMigrate(&APIKey{})

}
