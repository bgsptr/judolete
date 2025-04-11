package main

import (
	"judolete/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	oauthGoogleConfig := config.NewOAuthGoogle("../credential.json")

	r := gin.Default()

	configBootstrap := &config.BootstrapConfig{
		OAuthGoogleConfig: oauthGoogleConfig,
	}

	config.Bootstrap(configBootstrap, r)

	r.Run(":8080")
}