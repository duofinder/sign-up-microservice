package services

import (
	"net/http"

	"github.com/duofinder/auth-microservice/types"
	"github.com/gin-gonic/gin"
)

func SignupService(input *types.SignupInput) *types.Response {
	_, err := input.EncryptPasswordFunc(input.Password)
	if err != nil {
		return &types.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       gin.H{"error": err.Error()},
		}
	}

	// err = input.CreateAuthRepository(input.DB, input.Contact, hash)
	// if err != nil {
	// 	return &types.Response{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Body:       gin.H{"error": err.Error()},
	// 	}
	// }

	return &types.Response{
		StatusCode: http.StatusCreated,
		Body:       gin.H{"message": "success"},
	}
}
