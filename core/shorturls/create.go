package shorturls

import (
	"fmt"
	"time"

	"github.com/joaoprodrigo/shlink-go/config"
	"github.com/joaoprodrigo/shlink-go/core/models"
	"github.com/joaoprodrigo/shlink-go/core/utils"
)

// CreateShortURL generates a new short url based on meta received
func CreateShortURL(meta models.ShortURLMeta) (*models.ShortURL, error) {

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
		fmt.Printf("time: %s\n", now)
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
		shortURL.SetDomain(config.ShortDomainHost)
	} else {
		shortURL.SetDomain(meta.Domain)
	}

	// check original url?
	shortURL.OriginalURL = meta.LongURL

	// generate short code with length or use slug
	if meta.CustomSlug != "" {
		shortURL.ShortCode = meta.CustomSlug
	} else {
		shortCode, err := makeSlug(int(meta.ShortCodeLength))
		if err != nil {
			return nil, err
		}

		shortURL.ShortCode = shortCode
	}

	// max visits
	shortURL.MaxVisits = meta.MaxVisits

	// tags
	err := shortURL.AssignTags(meta.Tags)
	if err != nil {
		return nil, err
	}

	models.DB.Debug().Create(&shortURL)
	return &shortURL, nil //edit
}
