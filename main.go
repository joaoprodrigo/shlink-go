package main

import (
	"github.com/joaoprodrigo/shlink-go/cli"
	"github.com/joaoprodrigo/shlink-go/config"
	"github.com/joaoprodrigo/shlink-go/core/models"
)

func main() {

	config.Setup()

	models.StartupDB()

	cli.ParseArguments()
}
