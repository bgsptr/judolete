package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"judolete/internal/service"
)

type YoutubeController struct {
	YoutubeService *service.YoutubeService
	// YoutubeService *service.YoutubeService
}

func NewYoutubeController(youtubeService *service.YoutubeService) *YoutubeController {
	return &YoutubeController{
		YoutubeService: youtubeService,
	}
}

func (yc *YoutubeController) Init(c *gin.Context) {
	err := yc.YoutubeService.InitYoutubeService()
	if err != nil {
		log.Println(err)
		return err
	}

	c.JSON(200, gin.H{"message": "successfully init youtube service"})
}