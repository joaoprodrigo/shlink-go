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

// SetDomain sets DomainId in ShortURL, creating that ID if necessary
func (s *ShortURL) SetDomain(domainAuthority string) error {
	var domain Domain

	DB.Where(Domain{Authority: domainAuthority}).FirstOrCreate(&domain)

	s.DomainID = domain.ID

	return nil
}

// AssignTags Creates the Tags that don't exist and assigns them to the ShortURL
func (s *ShortURL) AssignTags(tagNames []string) error {
	var tags []Tag

	// Load Tags that already exist
	DB.Where("Name in (?)", tagNames).Find(&tags)

	// Create the ones that don't

newTagLoop:
	// Loop the Found Tags for each new tag looking for it
	for _, newTag := range tagNames {
		for _, oldTag := range tags {

			// If we find it, continue the outer loop so we dont append it
			if oldTag.Name == newTag {
				continue newTagLoop
			}
		}
		tags = append(tags, Tag{Name: newTag})
	}

	s.Tags = tags

	return nil
}

// ShortURLMeta represents metadata passed from REST or CLI to generate a short url
type ShortURLMeta struct {
	LongURL         string
	Tags            []string
	ValidSince      string
	ValidUntil      string
	CustomSlug      string
	MaxVisits       uint
	FindIfExists    bool
	Domain          string
	ShortCodeLength uint
}
