package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
Options:
  -h, --help            Display this help message
  -q, --quiet           Do not output any message
  -V, --version         Display this application version
      --ansi            Force ANSI output
      --no-ansi         Disable ANSI output
  -n, --no-interaction  Do not ask any interactive question
  -v|vv|vvv, --verbose  Increase the verbosity of messages: 1 for normal output, 2 for more verbose output and 3 for debug

Available commands:
  help                Displays help for a command
  list                Lists commands
 api-key
  api-key:disable     Disables an API key.
  api-key:generate    Generates a new valid API key.
  api-key:list        Lists all the available API keys.
 db
  db:create           Creates the database needed for shlink to work. It will do nothing if the database already exists
  db:migrate          Runs database migrations, which will ensure the shlink database is up to date.
 short-url
  short-url:delete    Deletes a short URL
  short-url:generate  Generates a short URL for provided long URL and returns it
  short-url:list      List all short URLs
  short-url:parse     Returns the long URL behind a short code
  short-url:visits    Returns the detailed visits information for provided short code
 tag
  tag:create          Creates one or more tags.
  tag:delete          Deletes one or more tags.
  tag:list            Lists existing tags.
  tag:rename          Renames one existing tag.
 visit
  visit:locate        Resolves visits origin locations.
*/

// ParseArguments gets OS Args and executes the required function
func ParseArguments() {

	if len(os.Args) < 2 {
		// if no arguments, run the server
		fmt.Println("Expected command but got None")
		os.Exit(1)
	}

	apiGenerateCmd := flag.NewFlagSet("api-key:generate", flag.ExitOnError)
	apiGenerateExp := apiGenerateCmd.String("expires", "", "expiration date YYYY-MM-DD")

	command := strings.ToLower(os.Args[1])
	switch command {
	case "api-key:generate":
		apiGenerateCmd.Parse(os.Args[2:])
		apiKeyGenerate(*apiGenerateExp)

	case "api-key:disable":
		if len(os.Args) != 3 {
			fmt.Println("Disable requires API Key and no other arguments")
			os.Exit(1)
		}
		apiKeyDisable(os.Args[2])

	case "api-key:list":
		apiKeyList()

	case "short-url:list":
		fmt.Println("Not Implemented")

	case "short-url:generate":
		meta := parseShortURLMeta(os.Args[2:])
		shortURLGenerate(meta)

	default:
		fmt.Println("No command or unexpected command found, exiting")
		os.Exit(1)
	}
}
