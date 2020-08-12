package security

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joaoprodrigo/shlink-go/config"

	"github.com/joaoprodrigo/shlink-go/core/models"

	"github.com/google/uuid"
)

// AuthService is an Authorization Service for storing, retrieving and verifying authorization keys
type AuthService struct {
	Config *config.ConfigurationRepo
	DB     *gorm.DB
}

// NewAuthService returns a new initialized Authorization Service
func NewAuthService(config *config.ConfigurationRepo, db *gorm.DB) *AuthService {
	return &AuthService{Config: config, DB: db}
}

// AuthorizeAPIKey returns an error if key is disabled, doesnt exist or is expired
func (s *AuthService) AuthorizeAPIKey(key string) error {

	//TODO Test

	var apiKey models.APIKey
	now := time.Now()

	if err := s.DB.Where(
		"key = ? AND enabled = ? AND (expiration_date is null OR expiration_date >= ?)",
		key,
		true,
		now).First(&apiKey).Error; err != nil {

		return errors.New("Unauthorized key")
	}

	return nil
}

// CreateAPIKey creates a new key in the db and returns the key
func (s *AuthService) CreateAPIKey(expirationDate *time.Time) (string, error) {
	keyUUID := uuid.New().String()

	key := models.APIKey{Key: keyUUID, ExpirationDate: expirationDate, Enabled: true}

	s.DB.Create(&key)

	return keyUUID, nil
}

// DisableAPIKey disables a given keyID so it can no longer be used
func (s *AuthService) DisableAPIKey(keyID string) error {
	key, err := s.getAPIKey(keyID)

	if err != nil {
		return errors.New("API Key not found")
	}

	s.DB.Model(&key).Update("enabled", false)

	return nil
}

// ListAPIKeys returns a list of currently enabled API Keys
func (s *AuthService) ListAPIKeys() []string {
	var keys []models.APIKey
	var stringKeys []string = []string{}

	s.DB.Where("enabled = ? AND (expiration_date is null OR expiration_date >= ?)", true, time.Now()).Find(&keys)

	for _, v := range keys {
		stringKeys = append(stringKeys, v.Key)
	}

	return stringKeys
}

func (s *AuthService) getAPIKey(keyID string) (*models.APIKey, error) {
	var key models.APIKey

	if err := s.DB.Where("key = ?", keyID).First(&key).Error; err != nil {
		return nil, err
	}

	return &key, nil
}
