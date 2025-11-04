package projectservice

import (
	repository "gintugas/modules/components/Project/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectMemberService struct {
	MemberRepo *repository.ProjectMemberRepo
}

func NewProjectMemberService(memberRepo *repository.ProjectMemberRepo) *ProjectMemberService {
	return &ProjectMemberService{MemberRepo: memberRepo}
}

type AddMemberRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

func (s *ProjectMemberService) AddMember(ctx *gin.Context) {
	projectID := ctx.Param("project_id")

	var req AddMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.MemberRepo.AddMember(projectID, req.UserID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Member added successfully"})
}

func (s *ProjectMemberService) RemoveMember(ctx *gin.Context) {
	projectID := ctx.Param("project_id")
	userID := ctx.Param("user_id")

	if err := s.MemberRepo.RemoveMember(projectID, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Member removed successfully"})
}

func (s *ProjectMemberService) GetProjectMembers(ctx *gin.Context) {
	projectID := ctx.Param("project_id")

	members, err := s.MemberRepo.GetProjectMembers(projectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"members": members})
}
