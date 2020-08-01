package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// APIKey is used for signing in via REST interface
type APIKey struct {
	gorm.Model

	Key            string `gorm:"type:varchar(255)"`
	ExpirationDate *time.Time
	Enabled        bool
}
