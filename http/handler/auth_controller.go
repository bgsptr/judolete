package handler

import (
	"net/http"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"judolete/internal/service"
)

type AuthController struct {
	TokenService *service.TokenService
}

func NewAuthController(tokenService *service.TokenService) *AuthController {
	return &AuthController{
		TokenService: tokenService,
	}
}

func (a *AuthController) RedirectToAuthURL(c *gin.Context) {
	url := a.TokenService.GetAuthURL()

	c.JSON(200, gin.H{"url": url})
}

func (a *AuthController) GetToken(c *gin.Context) {
	authCode, _ := c.GetQuery("code")
	// if err != nil {
	// 	log.Println("can't find code from auth url")
	// }

	token := a.TokenService.FetchTokenFromCallback(authCode);

	tokenJson, err := json.Marshal(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize token"})
		return
	}

	c.SetCookie(
		"token",
		string(tokenJson),
		3600 * 24,
		"/",
		"localhost",
		false, // secure
		true, // httponly cookie
	)

	c.JSON(200, gin.H{"message": "Cookie Set!"})
}