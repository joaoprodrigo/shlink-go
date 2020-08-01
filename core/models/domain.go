package models

import (
	"github.com/jinzhu/gorm"
)

// Domain is something im not sure in shlink
type Domain struct {
	gorm.Model

	Authority string `gorm:"type:varchar(255)"`

	ShortURLs []ShortURL
}
