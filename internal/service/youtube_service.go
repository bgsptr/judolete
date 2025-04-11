package service

import (
	"context"
	"log"
	
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YoutubeService struct {
	OAuthGoogleConfig *oauth2.Config
	YoutubeClient     *youtube.Service
}

func NewYoutubeService(oauth *oauth2.Config) *YoutubeService {
	return &YoutubeService{
		OAuthGoogleConfig: oauth,
	}
}

func (y *YoutubeService) InitYoutubeService(tok *oauth2.Token) *youtube.Service {
	ctx := context.Background()
	tokenSource := y.OAuthGoogleConfig.TokenSource(ctx, tok)

	youtubeSrv, err := youtube.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		// return err
		log.Println(err)
	}

	return youtubeSrv
}