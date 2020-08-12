package shorturls

// ShortURLMeta represents metadata passed from REST or CLI to generate a short url
type ShortURLMeta struct {
	LongURL         string
	Tags            []string
	ValidSince      string
	ValidUntil      string
	CustomSlug      string
	MaxVisits       uint
	FindIfExists    bool
	Domain          string
	ShortCodeLength uint
}
