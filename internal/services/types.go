package services

import "github.com/xandervanderweken/Password/internal/persistence"

type GeneratorService interface {
	CreateNewPassword(length int, useNumbers bool, useSpecial bool) (string, error)
}

type AccountService interface {
	GetAccountByDomainName(domainName string) *persistence.AccountEntry
	CreateAccount(domainName string, username string, password string) *persistence.AccountEntry
	DeleteAccount(accountId uint)
}
