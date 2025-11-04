package serviceroute

import (
	"database/sql"
	"fmt"
	projectrepo "gintugas/modules/components/Project/repository"
	projectservice "gintugas/modules/components/Project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProjectRouter(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			projekRepo = projectrepo.NewRepository(db)
			projekSrv  = projectservice.NewService(projekRepo)
		)

		Project, err := projekSrv.CreateProjekService(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": fmt.Sprintf("Data projek berhasil ditambahkan"),
			"Project": Project,
		})
	}
}

func GetAllProjektRouter(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			projekRepo = projectrepo.NewRepository(db)
			projekSrv  = projectservice.NewService(projekRepo)
		)

		Project, err := projekSrv.GetAllProjekService(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("successfully get all Project data"),
			"Project": Project,
		})
	}
}

func GetProjectRouter(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			projekRepo = projectrepo.NewRepository(db)
			projekSrv  = projectservice.NewService(projekRepo)
		)

		Project, err := projekSrv.GetProjekService(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("successfully get Project data"),
			"Project": Project,
		})
	}
}

func UpdateProjectRouter(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			projekRepo = projectrepo.NewRepository(db)
			projekSrv  = projectservice.NewService(projekRepo)
		)

		Project, err := projekSrv.UpdateProjekService(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Data Project Berhasil di Update"),
			"Project": Project,
		})
	}
}

func DeleteProjectRouter(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			projekRepo = projectrepo.NewRepository(db)
			projekSrv  = projectservice.NewService(projekRepo)
		)

		err := projekSrv.DeleteProjekService(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Project berhasil dihapus",
		})
	}
}
