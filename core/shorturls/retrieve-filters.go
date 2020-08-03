package shorturls

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/joaoprodrigo/shlink-go/config"
	"github.com/joaoprodrigo/shlink-go/core/utils"
)

// queryOrderBy orderBy string (query) The field from which you want to order the result. (Since v1.3.0) Available values : longUrl, shortCode, dateCreated, visits
func queryOrderBy(db *gorm.DB, orderBy string) (*gorm.DB, error) {
	if orderBy == "" {
		return db, nil
	}

	orderOptions := map[string]string{
		"longUrl":     "original_url",
		"shortCode":   "short_code",
		"dateCreated": "created_at",
		"visits":      "visits",
	}

	// TODO implement visit filtering (needs joins and also all visit functionality)
	if orderBy == "visits" {
		return db, errors.New("Visits are not implemented yet")
	}

	columnName, ok := orderOptions[orderBy]
	if !ok {
		return db, errors.New("Invalid order by " + orderBy)
	}

	return db.Order(columnName + " ASC"), nil
}

// queryFilterDate startDate, endDate The date (in ISO-8601 format) from which we want to get short URLs.
func queryFilterDate(db *gorm.DB, startDate string, endDate string) (*gorm.DB, error) {

	// startDate string (query) The date (in ISO-8601 format) from which we want to get short URLs.
	if startDate != "" {
		date, err := utils.ParseDateString(startDate)
		if err != nil {
			return db, err
		}
		db = db.Where("short_urls.created_at >= ?", date)
	}

	// endDate string (query)
	if endDate != "" {
		date, err := utils.ParseDateString(endDate)
		if err != nil {
			return db, err
		}
		db = db.Where("short_urls.created_at <= ?", date)
	}
	return db, nil
}

// queryByTerm searchTerm string (query) A query used to filter results by searching for it on the longUrl and shortCode fields
func queryByTerm(db *gorm.DB, searchTerm string) (*gorm.DB, error) {
	if searchTerm != "" {
		if !strings.Contains(searchTerm, "%") {
			searchTerm = "%" + searchTerm + "%"
		}
		db = db.Where("short_urls.original_url LIKE ? OR short_urls.short_code LIKE ?", searchTerm, searchTerm)
	}

	return db, nil
}

func queryPaginate(db *gorm.DB, pageNumber int) (*gorm.DB, error) {

	// page integer (query) The page to be displayed. Defaults to 1
	page := 1
	offset := 0
	if pageNumber > 0 {
		page = pageNumber
		offset = (page - 1) * config.DefaultItemsPerPage
		db = db.Offset(offset)
	}
	db = db.Limit(config.DefaultItemsPerPage)

	return db, nil
}

// queryByTags tags[] array[string] A list of tags used to filter the resultset. Only short URLs tagged with at least one of the provided tags will be returned
func queryByTags(db *gorm.DB, tags []string) (*gorm.DB, error) {
	if len(tags) == 0 {
		return db, nil
	}

	db = db.Joins("JOIN short_urls_in_tags ON short_urls_in_tags.short_url_id == short_urls.id").
		Joins("JOIN tags ON short_urls_in_tags.tag_id == tags.id").
		Where("tags.name IN (?)", tags)

	return db, nil
}
