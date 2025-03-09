package services

import (
	"log"

	"github.com/xandervanderweken/Password/internal/persistence"
)

type accountService struct {
	repo persistence.AccountEntryRepository
}

func NewAccountService(repo persistence.AccountEntryRepository) AccountService {
	return &accountService{
		repo: repo,
	}
}

func (s *accountService) GetAccountByDomainName(domainName string) *persistence.AccountEntry {
	acc, err := s.repo.GetAccountByDomainName(domainName)

	if err != nil {
		log.Fatalf("Could not get domain: %s", err)
		return nil
	}

	return acc
}

func (s *accountService) CreateAccount(domainName string, username string, password string) *persistence.AccountEntry {
	acc, err := s.repo.CreateAccount(domainName, username, password)

	if err != nil {
		log.Fatalf("Could not create account: %s", err)
		return nil
	}

	return acc
}

func (s *accountService) DeleteAccount(accountId uint) {
	err := s.repo.DeleteAccount(accountId)

	if err != nil {
		log.Fatalf("Could not delete account: %s", err)
	}
}
