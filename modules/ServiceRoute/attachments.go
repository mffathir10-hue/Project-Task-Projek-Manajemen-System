package serviceroute

import (
	attachmentservice "gintugas/modules/components/Attachments/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AttachmentHandler struct {
	attachmentService attachmentservice.AttachmentService
}

func NewAttachmentHandler(attachmentService attachmentservice.AttachmentService) *AttachmentHandler {
	return &AttachmentHandler{
		attachmentService: attachmentService,
	}
}

// POST /api/tasks/:task_id/attachments
func (h *AttachmentHandler) UploadAttachment(ctx *gin.Context) {
	response, err := h.attachmentService.UploadAttachment(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "File berhasil diupload",
		"data":    response,
	})
}

// GET /api/tasks/:task_id/attachments
func (h *AttachmentHandler) GetTaskAttachments(ctx *gin.Context) {
	attachments, err := h.attachmentService.GetTaskAttachments(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil attachments",
		"data":    attachments,
	})
}

// DELETE /api/attachments/:attachment_id
func (h *AttachmentHandler) DeleteAttachment(ctx *gin.Context) {
	if err := h.attachmentService.DeleteAttachment(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Attachment berhasil dihapus",
	})
}

// GET /api/attachments/:attachment_id/download
func (h *AttachmentHandler) DownloadAttachment(ctx *gin.Context) {
	filePath, err := h.attachmentService.DownloadAttachment(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.File(filePath)
}
