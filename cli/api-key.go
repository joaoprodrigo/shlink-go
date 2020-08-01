package cli

import (
	"fmt"
	"time"

	"github.com/joaoprodrigo/shlink-go/core/security"
)

func apiKeyGenerate(expirationTime string) {

	var expTime *time.Time

	if len(expirationTime) > 0 {
		const timeFormat = "2006-01-02"

		parsedTime, err := time.Parse(timeFormat, expirationTime)

		if err != nil {
			fmt.Printf("Error parsing given date")
			return
		}

		expTime = &parsedTime
	}

	key := security.CreateAPIKey(expTime)
	fmt.Printf("Generated key %v\n", key)

	return
}

func apiKeyDisable(apiKey string) {
	if err := security.DisableAPIKey(apiKey); err != nil {
		fmt.Printf("An error has occured: %s\n", err)
		return
	}

	fmt.Printf("Disabled API Key %s\n", apiKey)
}

func apiKeyList() {
	activeKeys := security.ListAPIKeys()

	if len(activeKeys) == 0 {
		fmt.Println("No active keys present")
		return
	}

	fmt.Println("Active Keys:")
	for _, v := range activeKeys {
		fmt.Printf("    %s\n", v)
	}
	fmt.Print("\n")
}
