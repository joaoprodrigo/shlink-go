package shorturls

import (
	"errors"
	"regexp"
	"time"

	"github.com/joaoprodrigo/shlink-go/core/models"

	"github.com/joaoprodrigo/shlink-go/config"
)

// GetShortURL retrieves the ShortURL Model from the db even if it isn't valid
func GetShortURL(domain string, shortcode string) (models.ShortURL, error) {
	var shortURL models.ShortURL

	err := models.DB.
		Joins("JOIN domains ON domains.id = short_urls.domain_id").
		Where("domains.authority = ? AND short_code = ?", domain, shortcode).
		// Where("short_code = ?", shortcode).
		First(&shortURL).
		Error

	return shortURL, err
}

// GetValidShortURL retrieves the ShortURL Model from the db even if it isn't valid
func GetValidShortURL(domain string, shortcode string) (models.ShortURL, error) {

	var shortURL models.ShortURL
	now := time.Now()

	err := models.DB.
		Joins("JOIN domains ON domains.id = short_urls.domain_id").
		Where("domains.authority = ? AND short_code = ?", domain, shortcode).
		Where("valid_since IS NULL OR valid_since <= ?", &now).
		Where("valid_until IS NULL OR valid_until > ?", &now).
		// Where("short_code = ?", shortcode).
		First(&shortURL).
		Error

	return shortURL, err
}

// ShortURLParams is a struct with Domain and ShortCode to be used with ParseShortURL
type ShortURLParams struct {
	Domain    string
	ShortCode string
}

// ParseShortURL returns the domain and shortcode of a given URL
func ParseShortURL(shortURL string) (ShortURLParams, error) {

	urlMatcher := `^(?:https?:\/\/)?([a-zA-Z0-9\.]*)\/?` + config.BasePath + `\/([a-zA-Z0-9]*)$`

	re := regexp.MustCompile(urlMatcher)

	matchGroup := re.FindStringSubmatch(shortURL)

	if len(matchGroup) != 3 {
		return ShortURLParams{}, errors.New("URLParser: No match for given URL " + shortURL)
	}

	return ShortURLParams{Domain: matchGroup[1], ShortCode: matchGroup[2]}, nil
}
