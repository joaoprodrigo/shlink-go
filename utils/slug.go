package utils

import (
	gonanoid "github.com/matoous/go-nanoid"
)

// MakeSlug generates a slug acording to a length and the shlink specification for alphabet
func MakeSlug(length int) (string, error) {
	const defaultSlugLength = 5
	const minSlugLength = 4
	const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	slugLength := defaultSlugLength

	if length > 0 && length > minSlugLength {
		slugLength = length
	}

	slug, err := gonanoid.Generate(alphabet, slugLength)
	return slug, err
}
