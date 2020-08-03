package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/joaoprodrigo/shlink-go/config"
	"github.com/joaoprodrigo/shlink-go/core/shorturls"

	"github.com/joaoprodrigo/shlink-go/core/models"
)

// shortURLGenerate is a CLI Command called to generate a short url based on some metadata
func shortURLGenerate(meta *models.ShortURLMeta) {
	shortURL, err := shorturls.CreateShortURL(meta)

	if err != nil {
		fmt.Printf("Error when generating Short URL: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf(
		"New ShortURL Created:\n"+
			"	Original URL: %s\n"+
			"	Short URL: %s://%s/%s\n",
		shortURL.OriginalURL,
		config.ShortDomainSchema, meta.Domain, shortURL.ShortCode,
	)
}

// parseShortURLMeta uses flags to determine all the passed options
// and generate a ShortURLMeta with them
func parseShortURLMeta(osArgs []string) *models.ShortURLMeta {

	metaFlags := flag.NewFlagSet("meta", flag.ExitOnError)

	longURL := metaFlags.String("long", "", "Long URL")
	tags := metaFlags.String("tags", "", "Comma separated list of tags")
	validSince := metaFlags.String("from", "", "Date when URL becomes valid (YYYY-MM-DD)")
	validUntil := metaFlags.String("until", "", "Date when URL expires (YYYY-MM-DD)")
	customSlug := metaFlags.String("slug", "", "Custom Slug to use instead of random string")
	maxVisits := metaFlags.Int("max-visits", 0, "Maximum number of visits until URL expires")
	findIfExists := metaFlags.Bool("find-exists", true, "Check if the long URL already exists")
	domain := metaFlags.String("domain", "", "Domain associated with the short url")
	length := metaFlags.Int("len", 0, "Short Code length, ignored if custom slug is provided")

	metaFlags.Parse(osArgs)

	meta := models.ShortURLMeta{
		LongURL:         *longURL,
		Tags:            tagsFromString(*tags),
		ValidSince:      *validSince,
		ValidUntil:      *validUntil,
		CustomSlug:      *customSlug,
		MaxVisits:       uint(*maxVisits),
		FindIfExists:    *findIfExists,
		Domain:          *domain,
		ShortCodeLength: uint(*length),
	}

	if meta.LongURL == "" {
		fmt.Println("Long URL (-long) is required")
		os.Exit(1)
	}

	return &meta
}

func tagsFromString(sTags string) []string {

	if len(strings.TrimSpace(sTags)) == 0 {
		return []string{}
	}

	rawTags := strings.Split(sTags, ",")

	for i, v := range rawTags {
		rawTags[i] = strings.TrimSpace(v)
	}

	return rawTags
}
