package service

import (
	"google.golang.org/api/youtube/v3"
	"golang.org/x/oauth2"
)

type CommentService struct {
	YoutubeService *YoutubeService
}

func NewCommentService(ytClient *YoutubeService) *CommentService {
	return &CommentService {
		YoutubeService: ytClient,
	}
}

func (c *CommentService) FindAllCommentInVideo(videoId string, token *oauth2.Token) ([]*youtube.CommentThread, error) {

	// filter comment by word 'snippet'
	youtubeSrv := c.YoutubeService.InitYoutubeService(token)
	call := youtubeSrv.CommentThreads.List([]string{"snippet"}).
		VideoId(videoId).
		TextFormat("plainText").
		MaxResults(100) // adjust as needed, max 100

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}


func (c *CommentService) DeleteCommentById(commentId string, token *oauth2.Token) error {
	youtubeSrv := c.YoutubeService.InitYoutubeService(token)
	deleteCommentCall := youtubeSrv.Comments.Delete(commentId)
	return deleteCommentCall.Do()
}