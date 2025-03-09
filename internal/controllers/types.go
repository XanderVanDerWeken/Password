package controllers

import "github.com/gin-gonic/gin"

type NewPassword struct {
	Length      int  `json:"length"`
	UseNumbers  bool `json:"useNumbers"`
	UseSpecials bool `json:"useSpecials"`
}

type PasswordDto struct {
	Password string `json:"password"`
}

type ErrorDto struct {
	Message string `json:"message"`
}

type PasswordController interface {
	CreatePassword(context *gin.Context)
}
