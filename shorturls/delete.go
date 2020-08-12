package shorturls

import "github.com/joaoprodrigo/shlink-go/core/models"

// DeleteShortURL removes the ShortURL Model from the db even if it isn't valid
func (s *ShortURLService) DeleteShortURL(domain string, shortcode string) error {

	err := s.DB.Debug().
		Joins("JOIN domains ON domains.id = short_urls.domain_id").
		Where("domains.authority = ? AND short_code = ?", domain, shortcode).
		// Where("short_code = ?", shortcode).
		Delete(models.ShortURL{}).
		Error

	return err
}
