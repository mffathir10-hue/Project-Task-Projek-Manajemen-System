package modules

import (
	"database/sql"
	serviceroute "gintugas/modules/ServiceRoute"
	repositoryprojek "gintugas/modules/components/Project/repository"
	servissprj "gintugas/modules/components/Project/service"
	controllers "gintugas/modules/components/userauth/controller"
	"gintugas/modules/components/userauth/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Initiator(router *gin.Engine, db *sql.DB, gormDB *gorm.DB) {
	authHandler := controllers.NewAuthHandler(gormDB)

	memberRepo := repositoryprojek.NewProjectMemberRepo(gormDB)
	memberService := servissprj.NewProjectMemberService(memberRepo)

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		auth.Use(middleware.AuthMiddleware())
		{

			auth.POST("/projects/:project_id/members", memberService.AddMember)
			auth.GET("/projects/:project_id/members", memberService.GetProjectMembers)
			auth.DELETE("/projects/:project_id/members/:user_id", memberService.RemoveMember)
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
		}

	}
}
