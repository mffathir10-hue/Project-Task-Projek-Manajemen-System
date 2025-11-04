package modules

import (
	"database/sql"
	serviceroute "gintugas/modules/ServiceRoute"
	services "gintugas/modules/components/Mail/service"
	repositoryprojek "gintugas/modules/components/Project/repository"
	servissprj "gintugas/modules/components/Project/service"
	taskrepository "gintugas/modules/components/Tasks/repository"
	taskservice "gintugas/modules/components/Tasks/service"
	controllers "gintugas/modules/components/auth/controller"
	"gintugas/modules/components/auth/middleware"

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
		}

	}
}
