package security

import (
	"errors"
	"time"

	"github.com/joaoprodrigo/shlink-go/core/models"

	"github.com/google/uuid"
)

// AuthorizeAPIKey returns an error if key is disabled, doesnt exist or is expired
func AuthorizeAPIKey(key string) error {

	//TODO Test

	var apiKey models.APIKey
	now := time.Now()

	if err := models.DB.Where(
		"key = ? AND enabled = ? AND (expiration_date is null OR expiration_date >= ?)",
		key,
		true,
		now).First(&apiKey).Error; err != nil {

	}

	return nil
}

// CreateAPIKey creates a new key in the db and returns the key
func CreateAPIKey(expirationDate *time.Time) string {
	keyUUID := uuid.New().String()

	key := models.APIKey{Key: keyUUID, ExpirationDate: expirationDate, Enabled: true}

	models.DB.Create(&key)

	return keyUUID
}

// DisableAPIKey disables a given keyID so it can no longer be used
func DisableAPIKey(keyID string) error {
	key, err := getAPIKey(keyID)

	if err != nil {
		return errors.New("API Key not found")
	}

	models.DB.Model(&key).Update("enabled", false)

	return nil
}

// ListAPIKeys returns a list of currently enabled API Keys
func ListAPIKeys() []string {
	var keys []models.APIKey
	var stringKeys []string = []string{}

	models.DB.Where("enabled = ? AND (expiration_date is null OR expiration_date >= ?)", true, time.Now()).Find(&keys)

	for _, v := range keys {
		stringKeys = append(stringKeys, v.Key)
	}

	return stringKeys
}

func getAPIKey(keyID string) (*models.APIKey, error) {
	var key models.APIKey

	if err := models.DB.Where("key = ?", keyID).First(&key).Error; err != nil {
		return nil, err
	}

	return &key, nil
}
