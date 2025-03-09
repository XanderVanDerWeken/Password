package persistence

type AccountEntry struct {
	ID         uint   `gorm:"primaryKey"`
	DomainName string `gorm:"size:100; not null"`
	Username   string `gorm:"size:100;not null"`
	Password   string `gorm:"size:100;not null"`
}

type AccountEntryRepository interface {
	GetAccountByDomainName(domainName string) (*AccountEntry, error)
	CreateAccount(domainName string, username string, password string) (*AccountEntry, error)
	DeleteAccount(accountId uint) error
}
