package main

import (
	"github.com/joaoprodrigo/shlink-go/cli"
	"github.com/joaoprodrigo/shlink-go/config"
	"github.com/joaoprodrigo/shlink-go/core/database"
	"github.com/joaoprodrigo/shlink-go/core/security"
	"github.com/joaoprodrigo/shlink-go/core/shorturls"
)

func main() {

	configRepo := config.NewConfigurationRepo()
	db := database.NewGormDB(configRepo, true)
	auth := security.NewAuthService(configRepo, db)
	shortURLService := shorturls.NewService(configRepo, db)

	cliParser := cli.NewCliInterface(shortURLService, auth, configRepo)
	cliParser.ParseArguments()
}
