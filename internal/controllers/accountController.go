package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xandervanderweken/Password/internal/services"
)

type accountController struct {
	accountService services.AccountService
}

func NewAccountController(accountService services.AccountService) AccountController {
	return &accountController{
		accountService: accountService,
	}
}

func (c *accountController) PostCreateAccount(context *gin.Context) {

}

func (c *accountController) GetAccountByDomainName(context *gin.Context) {

}

func (c *accountController) DeleteAccoundById(context *gin.Context) {

}
