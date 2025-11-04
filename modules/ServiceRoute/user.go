package serviceroute

import (
	"database/sql"
	"fmt"
	userrepo "gintugas/modules/components/auth/repository"
	userservice "gintugas/modules/components/auth/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsersRouter(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			usersrepo = userrepo.NewRepository(db)
			usersSrv  = userservice.NewService(usersrepo)
		)

		users, err := usersSrv.GetAllUsersService(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("successfully get all users data"),
			"Users":   users,
		})
	}
}

func GetUsersRouter(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			usersrepo = userrepo.NewRepository(db)
			usersSrv  = userservice.NewService(usersrepo)
		)

		users, err := usersSrv.GetUserService(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("successfully get users data"),
			"Users":   users,
		})
	}
}

func UpdateUsersRouter(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			usersrepo = userrepo.NewRepository(db)
			usersSrv  = userservice.NewService(usersrepo)
		)

		users, err := usersSrv.UpdateUserService(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Data users Berhasil di Update"),
			"Users":   users,
		})
	}
}

func DeleteUsersRouter(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			usersrepo = userrepo.NewRepository(db)
			usersSrv  = userservice.NewService(usersrepo)
		)

		err := usersSrv.DeleteUserService(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "users berhasil dihapus",
		})
	}
}
