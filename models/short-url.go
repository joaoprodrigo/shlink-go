package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ShortURL is the struct that saves short-urls
type ShortURL struct {
	gorm.Model
	DomainID    uint
	OriginalURL string `gorm:"type:varchar(2048)"`
	ShortCode   string `gorm:"type:varchar(255)"`
	ValidSince  *time.Time
	ValidUntil  *time.Time
	MaxVisits   uint

	Tags []Tag `gorm:"many2many:short_urls_in_tags"`
}
