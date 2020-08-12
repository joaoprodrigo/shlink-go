package cli

import (
	"github.com/joaoprodrigo/shlink-go/config"
	"github.com/joaoprodrigo/shlink-go/core/security"
	"github.com/joaoprodrigo/shlink-go/core/shorturls"
)

// Parser is an interface for a ParseArguments function that triggers Service functions based on arguments from the command line
type Parser interface {
	ParseArguments()
}

// BasicCliInterface is the interface for interaction by command line arguments
type BasicCliInterface struct {
	shortURL *shorturls.ShortURLService
	auth     *security.AuthService
	config   *config.ConfigurationRepo
}

// NewCliInterface returns a new CLI Interface
func NewCliInterface(shortURL *shorturls.ShortURLService, auth *security.AuthService, config *config.ConfigurationRepo) Parser {
	return &BasicCliInterface{shortURL: shortURL, auth: auth, config: config}
}
