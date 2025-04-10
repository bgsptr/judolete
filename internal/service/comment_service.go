package service

import (
	"google.golang.org/api/youtube/v3"
)

type CommentService struct {
	YoutubeClient *youtube.Service
}

func NewCommentService(ytClient *youtube.Service) *CommentService {
	return &CommentService {
		YoutubeClient: ytClient,
	}
}

func (c *CommentService) DeleteCommentById(commentId string) error {
	deleteCommentCall := c.YoutubeClient.Comments.Delete(commentId)
	return deleteCommentCall.Do()
}