package models

// Tag is a GORM short url Tag
type Tag struct {
	ID        uint       `gorm:"primary_key"`
	Name      string     `gorm:"type:varchar(255)"`
	ShortURLs []ShortURL `gorm:"many2many:short_urls_in_tags"`
}
