package handlers

import (
	"database/sql"
	"net/http"

	"github.com/duofinder/auth-microservice/repositories"
	"github.com/duofinder/auth-microservice/services"
	"github.com/duofinder/auth-microservice/types"
	"github.com/gin-gonic/gin"
)

func SignupRoute(r *gin.Engine, db *sql.DB) {
	r.POST("/signup", func(ctx *gin.Context) {
		var login types.SignupInput

		if err := ctx.ShouldBindJSON(&login); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp := services.SignupService(&login, db, &repositories.Repo{})

		ctx.JSON(resp.StatusCode, resp.Body)
	})
}
