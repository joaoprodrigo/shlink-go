package shorturls

import (
	"errors"
	"regexp"
	"time"

	"github.com/joaoprodrigo/shlink-go/core/models"
)

// GetShortURL retrieves the ShortURL Model from the db even if it isn't valid
func (s *ShortURLService) GetShortURL(domain string, shortcode string) (models.ShortURL, error) {
	var shortURL models.ShortURL

	err := s.DB.
		Joins("JOIN domains ON domains.id = short_urls.domain_id").
		Where("domains.authority = ? AND short_code = ?", domain, shortcode).
		// Where("short_code = ?", shortcode).
		First(&shortURL).
		Error

	return shortURL, err
}

// GetValidShortURL retrieves the ShortURL Model from the db even if it isn't valid
func (s *ShortURLService) GetValidShortURL(domain string, shortcode string) (models.ShortURL, error) {

	var shortURL models.ShortURL
	now := time.Now()

	err := s.DB.
		Joins("JOIN domains ON domains.id = short_urls.domain_id").
		Where("domains.authority = ? AND short_code = ?", domain, shortcode).
		Where("valid_since IS NULL OR valid_since <= ?", &now).
		Where("valid_until IS NULL OR valid_until > ?", &now).
		// Where("short_code = ?", shortcode).
		First(&shortURL).
		Error

	return shortURL, err
}

//GetShortURLs retrieves a list of shortURLs
func (s *ShortURLService) GetShortURLs(params URLSearchParams) (*[]models.ShortURL, error) {

	var shortURLs []models.ShortURL
	var err error
	queryBuilder := s.DB.Model(&models.ShortURL{})

	queryBuilder, err = queryFilterDate(queryBuilder, params.StartDate, params.EndDate)
	if err != nil {
		return nil, err
	}

	queryBuilder, err = queryByTerm(queryBuilder, params.SearchTerm)
	if err != nil {
		return nil, err
	}

	queryBuilder, err = queryPaginate(queryBuilder, params.Page, s.Config.DefaultItemsPerPage)
	if err != nil {
		return nil, err
	}

	queryBuilder, err = queryOrderBy(queryBuilder, params.OrderBy)
	if err != nil {
		return nil, err
	}

	queryBuilder, err = queryByTags(queryBuilder, params.Tags)
	if err != nil {
		return nil, err
	}

	// Using Rows instead of Find because I need "Distinct" to filter duplication with Tags
	rows, err := queryBuilder.Select("DISTINCT \"short_urls\".*").Rows()
	defer rows.Close()
	for rows.Next() {
		var shortURL models.ShortURL
		s.DB.ScanRows(rows, &shortURL)
		shortURLs = append(shortURLs, shortURL)
	}

	return &shortURLs, err
}

// ShortURLParams is a struct with Domain and ShortCode to be used with ParseShortURL
type ShortURLParams struct {
	Domain    string
	ShortCode string
}

// ParseShortURL returns the domain and shortcode of a given URL
func (s *ShortURLService) ParseShortURL(shortURL string) (ShortURLParams, error) {

	urlMatcher := `^(?:https?:\/\/)?([a-zA-Z0-9\.]*)\/?` + s.Config.BasePath + `\/([a-zA-Z0-9]*)$`

	re := regexp.MustCompile(urlMatcher)

	matchGroup := re.FindStringSubmatch(shortURL)

	if len(matchGroup) != 3 {
		return ShortURLParams{}, errors.New("URLParser: No match for given URL " + shortURL)
	}

	return ShortURLParams{Domain: matchGroup[1], ShortCode: matchGroup[2]}, nil
}
