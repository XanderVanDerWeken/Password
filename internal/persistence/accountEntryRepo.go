package persistence

import (
	"gorm.io/gorm"
)

type accountEntryRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountEntryRepository {
	return &accountEntryRepository{
		db: db,
	}
}

func (r *accountEntryRepository) GetAccountByDomainName(domainName string) (*AccountEntry, error) {
	var accountEntry AccountEntry

	r.db.Where("domainName = ?", domainName).First(&accountEntry)

	return &accountEntry, nil
}

func (r *accountEntryRepository) CreateAccount(domainName string, username string, password string) (*AccountEntry, error) {
	accountEntry := AccountEntry{
		DomainName: domainName,
		Username:   username,
		Password:   password,
	}

	createRes := r.db.Create(&accountEntry)
	if err := createRes.Error; err != nil {
		return nil, err
	}

	return &accountEntry, nil
}

func (r *accountEntryRepository) DeleteAccount(accountId uint) error {
	var accountEntry AccountEntry

	r.db.Find(&accountEntry, accountId)

	r.db.Delete(&accountEntry)

	return nil
}
