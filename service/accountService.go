package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"banking/helpers"
)

// AccountService ..
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	NewTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

// DefaultAccountService ..
type DefaultAccountService struct {
	repo domain.AccountRepository
}

// NewAccount ..
func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Account{
		CustomerID:  req.CustomerID,
		OpeningDate: helpers.CurrentDateTime,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	newAccount, err := s.repo.Save(a)

	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()
	return &response, nil
}

// NewTransaction ..
func (s DefaultAccountService) NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	acc, err := s.repo.FindByID(req.AccountID)
	if err != nil {
		return nil, err
	}
	if !acc.CanWithdraw(req.Amount) {
		return nil, errs.NewValidationError("Insufficent balance")
	}

	a := domain.Transaction{
		AccountID:       req.AccountID,
		TransactionDate: helpers.CurrentDateTime,
		TransactionType: req.TransactionType,
		Amount:          req.Amount,
	}

	t, err := s.repo.SaveTransaction(a)

	if err != nil {
		return nil, err
	}

	response := t.ToNewTransactionResponseDto()
	return &response, nil
}

// NewAccountService ..
func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repository}
}
