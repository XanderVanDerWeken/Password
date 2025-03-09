package controllers

import "github.com/gin-gonic/gin"

type PasswordDto struct {
	Password string `json:"password"`
}

type PasswordController interface {
	CreatePassword(context *gin.Context)
}
