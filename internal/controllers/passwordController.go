package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xandervanderweken/Password/internal/services"
)

type passwordController struct {
	generatorService services.GeneratorService
}

func NewPasswordController(
	generatorService services.GeneratorService) PasswordController {
	return &passwordController{
		generatorService: generatorService,
	}
}

func (c *passwordController) PostCreatePassword(context *gin.Context) {
	var newPasswordBody NewPassword
	context.BindJSON(&newPasswordBody)

	password, err := c.generatorService.CreateNewPassword(newPasswordBody.Length, newPasswordBody.UseNumbers, newPasswordBody.UseSpecials)

	if err != nil {
		context.JSON(http.StatusBadRequest, ErrorDto{
			Message: "Could not create Password",
		})
	} else {
		context.JSON(http.StatusCreated, PasswordDto{
			Password: password,
		})
	}
}

func (c *passwordController) GetCreatePassword(context *gin.Context) {
	passwordLengthStr := context.DefaultQuery("length", "12")
	useNumbersStr := context.DefaultQuery("useNumbers", "true")
	useSpecialsStr := context.DefaultQuery("useSpecials", "false")

	passwordLength, err := strconv.Atoi(passwordLengthStr)
	if err != nil {
		sendErrorResponse(context, "Invalid password length")
		return
	}

	useNumbers, err := strconv.ParseBool(useNumbersStr)
	if err != nil {
		sendErrorResponse(context, "Invalid useNumbers Value")
		return
	}

	useSpecials, err := strconv.ParseBool(useSpecialsStr)
	if err != nil {
		sendErrorResponse(context, "Invalid useSpecials Value")
		return
	}

	password, err := c.generatorService.CreateNewPassword(passwordLength, useNumbers, useSpecials)

	if err != nil {
		sendErrorResponse(context, "Cound not create Password")
		return
	}

	context.JSON(http.StatusCreated, PasswordDto{
		Password: password,
	})
}

func sendErrorResponse(context *gin.Context, message string) {
	context.JSON(http.StatusBadRequest, ErrorDto{Message: message})
}
