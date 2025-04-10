package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/bgsptr/judolete/internal/service"
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

	c.JSON(200, h.Message{"url": url})
}

func (a *AuthController) GetToken(c *gin.Context) {
	authCode, err := c.GetQuery("code")
	if err != nil {
		log.Println("can't find code from auth url")
		return err
	}

	token := a.TokenService.FetchTokenFromCallback(authCode);

	c.SetCookie(
		"token",
		token,
		3600 * 24,
		"/",
		"localhost",
		false, // secure
		true, // httponly cookie
	)

	c.JSON(200, gin.H{"message": "Cookie Set!"})
}