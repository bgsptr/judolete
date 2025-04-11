package router

import (
	"github.com/gin-gonic/gin"
	"judolete/http/handler"
)

type APIConfig struct {
	Route *gin.RouterGroup
	AuthController *handler.AuthController
	CommentController *handler.CommentController
	YoutubeController *handler.YoutubeController
}

func NewAPIConfig(route *gin.RouterGroup, authController *handler.AuthController, commentController *handler.CommentController, youtubeController *handler.YoutubeController) *APIConfig {
	return &APIConfig{
		Route: route,
		AuthController: authController,
		CommentController: commentController,
		YoutubeController: youtubeController,
	}
}

func (a *APIConfig) DefineAllRoutes() {
	authRouter := a.Route.Group("/auth")
	{
		authRouter.POST("/token/callback", a.AuthController.GetToken)
		authRouter.GET("/", a.AuthController.RedirectToAuthURL)
	}

	// youtubeRouter := a.Route.Group("/youtube")
	// {
		// youtubeRouter.GET("/", a.YoutubeController.Init)
	// }

	commentRouter := a.Route.Group("/comment")
	{
		commentRouter.DELETE("/video/:videoId", a.CommentController.Delete)
		commentRouter.POST("/", a.CommentController.FindAllCommentController)
	}
}