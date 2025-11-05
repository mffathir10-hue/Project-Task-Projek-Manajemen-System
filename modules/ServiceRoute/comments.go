package serviceroute

import (
	. "gintugas/modules/components/command/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentsHandler struct {
	commentsService CommentsService
}

func NewCommentsHandler(commentsService CommentsService) *CommentsHandler {
	return &CommentsHandler{
		commentsService: commentsService,
	}
}

func (c *CommentsHandler) CreateComments(ctx *gin.Context) {
	comments, err := c.commentsService.CreateComments(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":  "Comments created successfully",
		"Comments": comments,
	})
}

func (c *CommentsHandler) GetTasksComments(ctx *gin.Context) {
	comments, err := c.commentsService.GetTasksComments(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Comments retrieved successfully",
		"comments": comments,
	})
}

func (c *CommentsHandler) GetCommentsByID(ctx *gin.Context) {
	comments, err := c.commentsService.GetCommentsByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Comments retrieved successfully",
		"comments": comments,
	})
}

func (c *CommentsHandler) UpdateComments(ctx *gin.Context) {
	comments, err := c.commentsService.UpdateComments(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Comments updated successfully",
		"comments": comments,
	})
}

func (c *CommentsHandler) DeleteComments(ctx *gin.Context) {
	if err := c.commentsService.DeleteComments(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Comments deleted successfully",
	})
}
