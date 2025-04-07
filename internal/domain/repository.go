package domain

type AccountRepository interface {
	Save(account *Account) error
	FindByAPIKey(APIKey string) (*Account, error)
	FindByID(id string) (*Account, error)
	UpdateBalance(account *Account) error
}
