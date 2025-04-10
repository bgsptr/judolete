package config

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	"github.com/bgsptr/judolete/internal/handler"
	"github.com/bgsptr/judolete/internal/router"
	"github.com/bgsptr/judolete/internal/service"
)

type BootstrapConfig struct {
	OAuthGoogle *oauth2.Config
}

func Bootstrap(config *BootstrapConfig, r *gin.Engine) {
	youtubeService := service.NewYoutubeService(config.OAuthGoogleConfig)
	tokenService := service.NewTokenService(config.OAuthGoogleConfig)
	commentService := service.NewCommentService(youtubeService.YoutubeClient)

	youtubeController := handler.NewYoutubeController(youtubeService)
	authController := handler.NewAuthController(tokenService)
	commentController := handler.NewCommentController(commentService)

	apiConfig := router.NewAPIConfig(r.Group("/api/v1"), authController, commentController, youtubeController)
	apiConfig.DefineAllRoutes()
}