package handler

import (
	// "log"
	// "net/http"
	// "encoding/json"

	// "github.com/gin-gonic/gin"
	"judolete/internal/service"
	// "golang.org/x/oauth2"
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

// func (yc *YoutubeController) Init(c *gin.Context) {
// 	cookieToken, err := c.Cookie("token")
// 	if err != nil {
// 		log.Println("Failed to retrieve token from cookie:", err)
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
// 		return
// 	}

// 	var token *oauth2.Token
// 	err = json.Unmarshal([]byte(cookieToken), &token)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	err = yc.YoutubeService.InitYoutubeService(token)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	c.JSON(200, gin.H{"message": "successfully init youtube service"})
// }