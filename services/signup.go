package services

import (
	"database/sql"
	"net/http"

	"github.com/duofinder/auth-microservice/repositories"
	"github.com/duofinder/auth-microservice/types"
	"github.com/duofinder/auth-microservice/utils"
	"github.com/gin-gonic/gin"
)

func SignupService(login *types.SignupInput, db *sql.DB, repo *repositories.Repo) *types.Response {
	hash, err := utils.EncryptPassword(login.Password)
	if err != nil {
		return &types.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       gin.H{"error": err},
		}
	}

	err = repo.CreateAuthRepository(db, login.Contact, hash)
	if err != nil {
		return &types.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       gin.H{"error": err},
		}
	}

	return &types.Response{
		StatusCode: http.StatusCreated,
		Body:       gin.H{"message": "success"},
	}
}
