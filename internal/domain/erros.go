package domain

import "errors"

var (
	ErrAccountNotFound               = errors.New("account not found")
	ErrAccountDuplicateKey           = errors.New("account duplicate key")
	ErrAccountAlreadyExists          = errors.New("account already exists")
	ErrInvalidAPIKey                 = errors.New("invalid API key")
	ErrInsufficientFunds             = errors.New("insufficient funds")
	ErrInvalidAmount                 = errors.New("invalid amount")
	ErrTransactionNotFound           = errors.New("transaction not found")
	ErrTransactionAlreadyExists      = errors.New("transaction already exists")
	ErrTransactionFailed             = errors.New("transaction failed")
	ErrTransactionLimitExceeded      = errors.New("transaction limit exceeded")
	ErrTransactionNotAllowed         = errors.New("transaction not allowed")
	ErrTransactionAlreadyProcessed   = errors.New("transaction already processed")
	ErrTransactionAlreadyCancelled   = errors.New("transaction already cancelled")
	ErrTransactionAlreadyRefunded    = errors.New("transaction already refunded")
	ErrTransactionAlreadyReversed    = errors.New("transaction already reversed")
	ErrTransactionAlreadyChargedBack = errors.New("transaction already charged back")
	ErrTransactionAlreadySettled     = errors.New("transaction already settled")
	ErrTransactionAlreadyDisputed    = errors.New("transaction already disputed")
)
