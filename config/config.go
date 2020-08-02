package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* Configurator for the Shlink-Go system

Sets up default variables, the database to use, etc.

Sets variables based on a priority list, from higher priority to lower:

- YAML file (to implement later)
- ENV variables
- Hardcoded defaults

If mandatory variables are not defined, prints an error and exits
*/

// ShortDomainHost is the custom short domain used for this shlink instance. For example doma.in. (Env SHORT_DOMAIN_HOST)
var ShortDomainHost string

// Setup loads the environment settings
func Setup() {

	setVariable(&ShortDomainHost, "SHORT_DOMAIN_HOST", nil)
}

func setVariable(varName interface{}, envName string, def interface{}) {

	// TODO check for valid

	if envValue, present := os.LookupEnv(envName); present {

		switch p := varName.(type) {

		case *string:
			*p = envValue

		case *bool:
			env, err := strconv.ParseBool(strings.ToLower(envValue))
			if err != nil {
				fmt.Printf("Env variable %s should be true or false, but is %s\n", envName, envValue)
				os.Exit(1)
			}
			*p = env

		case *int:
			env, err := strconv.Atoi(envValue)
			if err != nil {
				fmt.Printf("Env variable %s should be a number, but is %s\n", envName, envValue)
				os.Exit(1)
			}
			*p = env
		default:
			fmt.Printf("Unimplemented Variable Type for configuration: %T\n", envName)
			os.Exit(1)
		}
	} else {

		if def == nil {
			fmt.Printf("Undefined required variable: %s", envName)
			os.Exit(1)
		}

		switch p := varName.(type) {
		case *string:
			*p = *def.(*string)
		case *bool:
			*p = *def.(*bool)
		case *int:
			*p = *def.(*int)
		default:
			fmt.Printf("Unimplemented Variable Type for configuration (default): %T\n", envName)
			os.Exit(1)
		}
	}
}

/*
unimplemented:

SHORT_DOMAIN_SCHEMA: Either http or https.
DB_DRIVER: sqlite (which is the default value), mysql, maria, postgres or mssql.
DB_NAME: The database name to be used when using an external database driver. Defaults to shlink.
DB_USER: The username credential to be used when using an external database driver.
DB_PASSWORD: The password credential to be used when using an external database driver.
DB_HOST: The host name of the database server when using an external database driver.
DB_PORT: The port in which the database service is running when using an external database driver.
	Default value is based on the value provided for DB_DRIVER:
	mysql or maria -> 3306
	postgres -> 5432
	mssql -> 1433
DISABLE_TRACK_PARAM: The name of a query param that can be used to visit short URLs avoiding the visit to be tracked. This feature won't be available if not value is provided.
DELETE_SHORT_URL_THRESHOLD: The amount of visits on short URLs which will not allow them to be deleted. Defaults to 15.
VALIDATE_URLS: Boolean which tells if shlink should validate a status 20x is returned (after following redirects) when trying to shorten a URL. Defaults to false.
BASE_PATH (v1.19.0) : The base path from which you plan to serve shlink, in case you don't want to serve it from the root of the domain. Defaults to ''.
INVALID_SHORT_URL_REDIRECT_TO (v1.20.0) : If a URL is provided here, when a user tries to access an invalid short URL, he/she will be redirected to this value. If this env var is not provided, the user will see a generic 404 - not found page.
REGULAR_404_REDIRECT_TO (v1.20.0) : If a URL is provided here, when a user tries to access a URL not matching any one supported by the router, he/she will be redirected to this value. If this env var is not provided, the user will see a generic 404 - not found page.
BASE_URL_REDIRECT_TO (v1.20.0) : If a URL is provided here, when a user tries to access Shlink's base URL, he/she will be redirected to this value. If this env var is not provided, the user will see a generic 404 - not found page.
WEB_WORKER_NUM (v1.21.0) : The amount of concurrent http requests this shlink instance will be able to server. Defaults to 16.
TASK_WORKER_NUM (v1.21.0) : The amount of concurrent background tasks this shlink instance will be able to execute. Defaults to 16.
VISITS_WEBHOOKS (v1.21.0) : A comma-separated list of URLs that will receive a POST request when a short URL receives a visit.
DEFAULT_SHORT_CODES_LENGTH (v2.1.0) : The length you want generated short codes to have. It defaults to 5 and has to be at least 4, so any value smaller than that will fall back to 4.
GEOLITE_LICENSE_KEY (v2.1.4) : The license key used to download new GeoLite2 database files. This is not mandatory, as a default license key is provided, but it is strongly recommended that you provide your own. Go to GeoLite2 license key to know how to generate it.
REDIS_SERVERS: A comma-separated list of redis servers where Shlink locks are stored (locks are used to prevent some operations to be run more than once in parallel).
MERCURE_PUBLIC_HUB_URL (v2.2.0) : The public URL of a mercure hub server to which Shlink will sent updates. This URL will also be served to consumers that want to subscribe to those updates.
MERCURE_INTERNAL_HUB_URL (v2.2.0) : An internal URL for a mercure hub. Will be used only when publishing updates to mercure, and does not need to be public. If this is not provided but MERCURE_PUBLIC_HUB_URL was, the former one will be used to publish updates.
MERCURE_JWT_SECRET (v2.2.0) : The secret key that was provided to the mercure hub server, in order to be able to generate valid JWTs for publishing/subscribing to that server.
ANONYMIZE_REMOTE_ADDR (v2.2.0) : Tells if IP addresses from visitors should be obfuscated before storing them in the database. Default value is true.
Warning
Setting this to false will make your Shlink instance no longer be in compliance with the GDPR and other similar data protection regulations.
REDIRECT_STATUS_CODE (v2.3.0) : Either 301 or 302. Used to determine if redirects from short to long URLs should be done with a 301 or 302 status. Defaults to 302.
REDIRECT_CACHE_LIFETIME (v2.3.0) : Allows to set the amount of seconds that redirects should be cached when redirect status is 301. Default values is 30.
*/
