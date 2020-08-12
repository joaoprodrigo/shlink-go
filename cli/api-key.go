package cli

import (
	"fmt"
	"time"
)

func (p BasicCliInterface) apiKeyGenerate(expirationTime string) {

	//TODO update to use utils.ParseDateString

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

	key, _ := p.auth.CreateAPIKey(expTime)
	fmt.Printf("Generated key %v\n", key)

	return
}

func (p BasicCliInterface) apiKeyDisable(apiKey string) {
	if err := p.auth.DisableAPIKey(apiKey); err != nil {
		fmt.Printf("An error has occured: %s\n", err)
		return
	}

	fmt.Printf("Disabled API Key %s\n", apiKey)
}

func (p BasicCliInterface) apiKeyList() {
	activeKeys := p.auth.ListAPIKeys()

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
