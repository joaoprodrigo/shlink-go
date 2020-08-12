package shorturls

import (
	"github.com/jinzhu/gorm"
	"github.com/joaoprodrigo/shlink-go/config"
)

// ShortURLService is the Service Logic for creating, updating, retrieving and deleting shortURLs
type ShortURLService struct {
	Config *config.ConfigurationRepo
	DB     *gorm.DB
}

// NewService returns a new ShortURLs Service
func NewService(config *config.ConfigurationRepo, db *gorm.DB) *ShortURLService {

	return &ShortURLService{config, db}
}
