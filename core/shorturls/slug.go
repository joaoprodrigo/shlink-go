package shorturls

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid"
)

// MakeSlug generates a slug acording to a length and the shlink specification for alphabet
func MakeSlug(length int) (string, error) {
	const defaultSlugLength = 5
	const minSlugLength = 4
	const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	slugLength := defaultSlugLength

	if length < minSlugLength {
		return "", errors.New("Length of slug is lower than the minimum required")
	}

	if length > 0 {
		slugLength = length
	}

	slug, err := gonanoid.Generate(alphabet, slugLength)
	return slug, err
}
