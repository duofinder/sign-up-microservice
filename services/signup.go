package services

import (
	"net/http"

	"github.com/duofinder/auth-microservice/types"
	"github.com/gin-gonic/gin"
)

func SignupService(input *types.SignupInput) *types.Response {
	hash, err := input.EncryptPasswordFunc(input.Password)
	if err != nil {
		return &types.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       gin.H{"error": err},
		}
	}

	err = input.CreateAuthRepository(input.DB, input.Contact, hash)
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
