package modules

import (
	"database/sql"
	serviceroute "gintugas/modules/ServiceRoute"
	attachmentrepository "gintugas/modules/components/Attachments/repository"
	attachmentservice "gintugas/modules/components/Attachments/service"
	services "gintugas/modules/components/Mail/service"
	repositoryprojek "gintugas/modules/components/Project/repository"
	servissprj "gintugas/modules/components/Project/service"
	taskrepository "gintugas/modules/components/Tasks/repository"
	taskservice "gintugas/modules/components/Tasks/service"
	controllers "gintugas/modules/components/auth/controller"
	"gintugas/modules/components/auth/middleware"
	. "gintugas/modules/components/command/repository"
	. "gintugas/modules/components/command/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Initiator(router *gin.Engine, db *sql.DB, gormDB *gorm.DB) {
	authHandler := controllers.NewAuthHandler(gormDB)

	memberRepo := repositoryprojek.NewProjectMemberRepo(gormDB)
	memberService := servissprj.NewProjectMemberService(memberRepo)

	taskRepo := taskrepository.NewTaskRepository(gormDB)

	mailService := services.NewMailService()
	taskService := taskservice.NewTaskService(taskRepo, mailService)
	taskController := serviceroute.NewTaskController(taskService)

	commentsRepo := NewCommentsRepository(gormDB)
	commentsService := NewTaskService(commentsRepo)
	commentsHandler := serviceroute.NewCommentsHandler(commentsService)

	attachmentRepo := attachmentrepository.NewAttachmentRepository(gormDB)

	//path menyimpan file
	uploadPath := "./uploads/attachments"
	attachmentService := attachmentservice.NewAttachmentService(attachmentRepo, uploadPath)
	attachmentHandler := serviceroute.NewAttachmentHandler(attachmentService)

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		auth.Use(middleware.AuthMiddleware())
		{

			// User routes
			auth.GET("/users", serviceroute.GetAllUsersRouter(db))
			auth.GET("/users/:id", serviceroute.GetUsersRouter(db))
			auth.PUT("/users/:id", serviceroute.UpdateUsersRouter(db))
			auth.DELETE("/users/:id", serviceroute.DeleteUsersRouter(db))

			// Project routes
			auth.POST("/project", serviceroute.CreateProjectRouter(db))
			auth.GET("/project", serviceroute.GetAllProjektRouter(db))
			auth.GET("/project/:id", serviceroute.GetProjectRouter(db))
			auth.PUT("/project/:id", serviceroute.UpdateProjectRouter(db))
			auth.DELETE("/project/:id", serviceroute.DeleteProjectRouter(db))

			member := auth.Group("/projects/:project_id/members")
			{
				member.POST("", memberService.AddMember)
				member.GET("", memberService.GetProjectMembers)
				member.DELETE("/:user_id", memberService.RemoveMember)
			}

			// Tasks Route
			tasks := auth.Group("/projects/:project_id/tasks")
			{
				tasks.POST("", taskController.CreateTask)
				tasks.GET("", taskController.GetProjectTasks)
				tasks.GET("/:task_id", taskController.GetTaskByID)
				tasks.PUT("/:task_id", taskController.UpdateTask)
				tasks.DELETE("/:task_id", taskController.DeleteTask)
			}

			//Comments Route
			comments := auth.Group("/tasks/:task_id/comments")
			{
				comments.POST("", commentsHandler.CreateComments)
				comments.GET("", commentsHandler.GetTasksComments)
				comments.GET("/:comments_id", commentsHandler.GetCommentsByID)
				comments.PUT("/:comments_id", commentsHandler.UpdateComments)
				comments.DELETE("/:comments_id", commentsHandler.DeleteComments)
			}

			attachments := auth.Group("")
			{
				//upload ke task berdasarkan id
				attachments.POST("/tasks/:task_id/attachments", attachmentHandler.UploadAttachment)
				attachments.GET("/tasks/:task_id/attachments", attachmentHandler.GetTaskAttachments)
				attachments.DELETE("/attachments/:attachment_id", attachmentHandler.DeleteAttachment)
				attachments.GET("/attachments/:attachment_id/download", attachmentHandler.DownloadAttachment)
			}
		}

	}
}
