package models

import (
	"github.com/jinzhu/gorm"
)

// Tag is a GORM short url Tag
type Tag struct {
	gorm.Model

	Name string `gorm:"type:varchar(255)"`

	ShortURLs []ShortURL `gorm:"many2many:short_urls_in_tags"`
}
