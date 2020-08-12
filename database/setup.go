package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //necessary for gorm
	"github.com/joaoprodrigo/shlink-go/config"
	"github.com/joaoprodrigo/shlink-go/core/models"
)

// NewGormDB creates a gorm DB access object with the configuration struct passed
func NewGormDB(config *config.ConfigurationRepo, autoMigrate bool) *gorm.DB {

	db, err := gorm.Open("sqlite3", config.DatabaseFile)

	if err != nil {
		panic("failed to connect database")
	}

	if autoMigrate {
		runMigrations(db)
	}

	return db
}

func runMigrations(db *gorm.DB) {

	db.AutoMigrate(&models.ShortURL{})
	db.AutoMigrate(&models.Domain{})
	db.AutoMigrate(&models.Tag{})
	db.AutoMigrate(&models.APIKey{})

}
