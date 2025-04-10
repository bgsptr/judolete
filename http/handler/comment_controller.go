package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"judolete/internal/service"
)

type CommentController struct {
	CommentService *service.CommentService
	// YoutubeService *service.YoutubeService
}

func NewCommentController(commentService *service.CommentService) *CommentController {
	return &CommentController{
		CommentService: commentService,
	}
}

func (cc *CommentController) Delete(c *gin.Context) {
	commentId := c.Param("id")
	err := cc.CommentService.DeleteCommentById(commentId)
	if err != nil {
		log.Println(err)
		return err
	}

	c.JSON(200, gin.H{"message": "Comment deleted", "id": id})
}