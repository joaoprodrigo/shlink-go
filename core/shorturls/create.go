package shorturls

import (
	"github.com/joaoprodrigo/shlink-go/core/models"
	"github.com/joaoprodrigo/shlink-go/core/utils"
)

// CreateShortURL generates a new short url based on meta received
func CreateShortURL(meta models.ShortURLMeta) (*models.ShortURL, error) {

	var shortURL models.ShortURL

	validSince, err := utils.ParseDateString(meta.ValidSince)
	if err != nil {
		return nil, err
	}

	validUntil, err := utils.ParseDateString(meta.ValidUntil)
	if err != nil {
		return nil, err
	}

	// domain?

	// check original url?

	// generate short code with length or use slug

	// max visits

	// tags

	shortURL.ValidSince = validSince
	shortURL.ValidUntil = validUntil

	return &shortURL, nil //edit
}
