package main

import (
	"github.com/joaoprodrigo/shlink-go/cli"
	"github.com/joaoprodrigo/shlink-go/core/models"
)

func main() {

	models.StartupDB()

	cli.ParseArguments()
}
