package config

import (
	// "context"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewOAuthGoogle(credentialPath string) *oauth2.Config {
	// ctx := context.Background()

	b, err := os.ReadFile(credentialPath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/youtube.force-ssl")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	return config
}
