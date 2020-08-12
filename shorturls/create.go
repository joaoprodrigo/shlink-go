package shorturls

import (
	"time"

	"github.com/joaoprodrigo/shlink-go/core/models"
	"github.com/joaoprodrigo/shlink-go/core/utils"
)

// CreateShortURL generates a new short url based on meta received
func (s *ShortURLService) CreateShortURL(meta *ShortURLMeta) (*models.ShortURL, error) {

	// TODO Fix issue:
	//     The model creation is updating the Tag Model even for tags that exist... eventually should be fixed
	//     FindIfExists not implemented

	var shortURL models.ShortURL

	if meta.ValidSince != "" {
		validSince, err := utils.ParseDateString(meta.ValidSince)
		if err != nil {
			return nil, err
		}
		shortURL.ValidSince = validSince
	} else {
		now := time.Now()
		shortURL.ValidSince = &now
	}

	if meta.ValidUntil != "" {
		validUntil, err := utils.ParseDateString(meta.ValidUntil)
		if err != nil {
			return nil, err
		}
		shortURL.ValidUntil = validUntil
	}

	// if domain is not set, use default
	if meta.Domain == "" {
		meta.Domain = s.Config.ShortDomainHost
		s.SetDomain(&shortURL, s.Config.ShortDomainHost)
	} else {
		s.SetDomain(&shortURL, meta.Domain)
	}

	// check original url?
	shortURL.OriginalURL = meta.LongURL

	// generate short code with length or use slug
	if meta.CustomSlug != "" {
		shortURL.ShortCode = meta.CustomSlug
	} else {
		shortCode, err := utils.MakeSlug(int(meta.ShortCodeLength))
		if err != nil {
			return nil, err
		}

		shortURL.ShortCode = shortCode
	}

	// max visits
	shortURL.MaxVisits = meta.MaxVisits

	// tags
	err := s.AssignTags(&shortURL, meta.Tags)
	if err != nil {
		return nil, err
	}

	s.DB.Create(&shortURL)
	return &shortURL, nil //edit
}

// SetDomain sets DomainId in ShortURL, creating that ID if necessary
func (s *ShortURLService) SetDomain(shortURL *models.ShortURL, domainAuthority string) error {
	var domain models.Domain

	s.DB.Where(models.Domain{Authority: domainAuthority}).FirstOrCreate(&domain)

	shortURL.DomainID = domain.ID

	return nil
}

// AssignTags Creates the Tags that don't exist and assigns them to the ShortURL
func (s *ShortURLService) AssignTags(shortURL *models.ShortURL, tagNames []string) error {
	var tags []models.Tag

	// Load Tags that already exist
	s.DB.Where("Name in (?)", tagNames).Find(&tags)

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
		tags = append(tags, models.Tag{Name: newTag})
	}

	shortURL.Tags = tags

	return nil
}
