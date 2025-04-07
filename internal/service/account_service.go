package service

import (
	"github.com/joaodematejr/imersao22/go-gateway/internal/domain"
	"github.com/joaodematejr/imersao22/go-gateway/internal/dto"
)

type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) CreateAccount(input dto.CreateAccountInput) (*dto.AccountOutputDTO, error) {
	account := dto.ToAccount(input)

	existingAccount, err := s.repository.FindByAPIKey(account.APIKey)

	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}
	if existingAccount != nil {
		return nil, domain.ErrAccountDuplicateKey
	}

	err = s.repository.Save(account)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutputDTO, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, domain.ErrAccountNotFound
	}

	account.AddBalance(amount)
	err = s.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) FindByAPIKey(apiKey string) (*dto.AccountOutputDTO, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, domain.ErrAccountNotFound
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) FindByID(id string) (*dto.AccountOutputDTO, error) {
	account, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, domain.ErrAccountNotFound
	}

	output := dto.FromAccount(account)
	return &output, nil
}
