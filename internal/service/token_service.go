package service

import (
	"context"
	"log"

	"golang.org/x/oauth2"
)

type TokenService struct {
	OAuthGoogleConfig *oauth2.Config
}

func NewTokenService(oauth *oauth2.Config) *TokenService {
	return &TokenService {
		OAuthGoogleConfig: oauth,
	}
}

func (t *TokenService) GetAuthURL() string {
	authUrl := t.OAuthGoogleConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return authUrl
}

func (t *TokenService) FetchTokenFromCallback(authCode string) *oauth2.Token {
	ctx := context.Background()
	tok, err := t.OAuthGoogleConfig.Exchange(ctx, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}