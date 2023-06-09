package services

import (
	"user-jwt-auth/models"
	"user-jwt-auth/models/entities"

	"github.com/gin-gonic/gin"
)

type UserServiceInterface interface {
	SignUp(ctx *gin.Context, req models.SignUpReq) (entities.Users, error)
	Login(ctx *gin.Context, req models.LoginReq) (string, error)
	UserDetils(ctx *gin.Context, email string) (models.UserDetailsRes, error)
}
