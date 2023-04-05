package handlers

import (
	"database/sql"
	"net/http"

	"github.com/duofinder/sign-up-microservice/repositories"
	"github.com/duofinder/sign-up-microservice/services"
	"github.com/duofinder/sign-up-microservice/types"
	"github.com/duofinder/sign-up-microservice/utils"
	"github.com/gin-gonic/gin"
)

func SignupRoute(r *gin.Engine, db *sql.DB) {
	r.POST("/signup", func(ctx *gin.Context) {
		var login types.UserData

		if err := ctx.ShouldBindJSON(&login); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp := services.SignupService(&types.SignupInput{
			UserData:                 login,
			DB:                       db,
			EncryptPasswordFunc:      utils.EncryptPassword,
			GenerateRefreshTokenFunc: utils.GenerateRefreshToken,
			CreateAuthRepository:     repositories.CreateAuthRepository,
			GenerateAccessTokenFunc:  utils.GenerateAccessToken,
		})

		ctx.JSON(resp.StatusCode, resp.Body)
	})
}
