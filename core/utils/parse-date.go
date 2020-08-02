package utils

import (
	"errors"
	"strings"
	"time"
)

// ParseDateString returns a *Time from a YYYY-MM-DD string
func ParseDateString(dateStr string) (*time.Time, error) {
	date := strings.TrimSpace(dateStr)
	if len(date) == 0 {
		return nil, errors.New("No date given")
	}

	const timeFormat = "2006-01-02"

	parsedTime, err := time.Parse(timeFormat, date)

	if err != nil {
		return nil, errors.New("Incorrect Date Format")
	}

	return &parsedTime, nil
}
