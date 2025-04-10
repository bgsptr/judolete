package cmd

import (
	"github.com/bgsptr/judolete/internal/config"
)

func main() {
	oauthGoogleConfig := config.NewOAuthGoogle("credential.json")

	r := gin.Default()

	configBootstrap := &config.BootstrapConfig{
		OAuthGoogle: oauthGoogleConfig,
	}

	config.Bootstrap(configBootstrap, r)

	r.Run(":8080")
}