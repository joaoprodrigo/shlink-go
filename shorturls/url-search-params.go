package shorturls

// URLSearchParams is a struct for passing params to search for a shortURL
type URLSearchParams struct {
	Page       int      // The page to be displayed. Defaults to 1
	SearchTerm string   // A query used to filter results by searching for it on the longUrl and shortCode fields.
	Tags       []string // A list of tags used to filter the resultset. Only short URLs tagged with at least one of the provided tags will be returned.
	OrderBy    string   // The field from which you want to order the result. (Since v1.3.0) Available values : longUrl, shortCode, dateCreated, visits
	StartDate  string   // The date (in ISO-8601 format) from which we want to get short URLs
	EndDate    string   // The date (in ISO-8601 format) until which we want to get short URLs
}
