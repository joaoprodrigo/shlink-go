package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //necessary for gorm
)

// DB is the global Db for use with application
var DB *gorm.DB

// StartupDB runs routines necessary for the DB initialization
func StartupDB() {

	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&ShortURL{})
	db.AutoMigrate(&Domain{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&APIKey{})

	DB = db
}
