package models

// Domain is something im not sure in shlink
type Domain struct {
	ID uint `gorm:"primary_key"`

	Authority string `gorm:"type:varchar(255)"`

	ShortURLs []ShortURL
}
