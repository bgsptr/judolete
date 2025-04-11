package handler

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"unicode"

	"github.com/gin-gonic/gin"
	"judolete/internal/service"
	"golang.org/x/oauth2"
)

type CommentController struct {
	CommentService *service.CommentService
	// YoutubeService *service.YoutubeService
}

type CommentRequest struct {
	VideoID     string `json:"videoId"`
}

type TextRequest struct {
	CommentId string
	TextDisplay string 
	TextOriginal string 
}

func NewCommentController(commentService *service.CommentService) *CommentController {
	return &CommentController{
		CommentService: commentService,
	}
}

var fancyUnicodeRanges = []*unicode.RangeTable{
	unicode.Cyrillic,
	unicode.Greek,
	&unicode.RangeTable{ // Mathematical Alphanumeric Symbols (U+1D400–U+1D7FF)
		R32: []unicode.Range32{
			{Lo: 0x1D400, Hi: 0x1D7FF, Stride: 1},
		},
	},
}

func containsFancyOnly(s string) bool {
	for _, r := range s {
		for _, table := range fancyUnicodeRanges {
			if unicode.Is(table, r) {
				return true
			}
		}
	}
	return false
}

func (cc *CommentController) FindAllCommentController(c *gin.Context) {
	var req CommentRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	decoded, errGetCookie := c.Cookie("token")
	if errGetCookie != nil {
		fmt.Errorf("access_token cookie not found: ", errGetCookie)
	}

	var token *oauth2.Token
    errUnmarshalToken := json.Unmarshal([]byte(decoded), &token)
    if errUnmarshalToken != nil {
        fmt.Errorf("failed to unmarshal oauth2.Token: ", errUnmarshalToken)
    }

	items, errArray := cc.CommentService.FindAllCommentInVideo(req.VideoID, token)
	if errArray != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errArray})
		return
	}

	c.JSON(200, gin.H{"message": items})
}

func (cc *CommentController) Delete(c *gin.Context) {
	// commentId := c.Param("id")
	videoId := c.Param("videoId")

	var token *oauth2.Token
	decoded, errGetCookie := c.Cookie("token")
	if errGetCookie != nil {
		fmt.Errorf("access_token cookie not found: ", errGetCookie)
	}

    errUnmarshalToken := json.Unmarshal([]byte(decoded), &token)
    if errUnmarshalToken != nil {
        fmt.Errorf("failed to unmarshal oauth2.Token: ", errUnmarshalToken)
    }
	
	items, errArray := cc.CommentService.FindAllCommentInVideo(videoId, token)
	if errArray != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errArray})
		return
	}


	var gamblerCommentList []TextRequest

	for _, comment := range items {
		id := comment.Id
		textDisplay := comment.Snippet.TopLevelComment.Snippet.TextDisplay
		textOriginal := comment.Snippet.TopLevelComment.Snippet.TextOriginal

		if containsFancyOnly(textDisplay) || containsFancyOnly(textOriginal) {
			gamblerCommentList = append(gamblerCommentList, TextRequest{
				CommentId: id,
				TextDisplay:  textDisplay,
				TextOriginal: textOriginal,
			})
			errDeleteComment := cc.CommentService.DeleteCommentById(id, token)
			if errDeleteComment != nil {
				log.Println(errDeleteComment)
			}
		}
	}

	c.JSON(200, gin.H{"message": "Comment deleted", "forbidden_comment": gamblerCommentList})
}