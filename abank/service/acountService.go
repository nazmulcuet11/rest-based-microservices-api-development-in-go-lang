package service

import (
	"abank/domain"
	"abank/dto"
	"abank/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func NewDefaultAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	account := domain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		Type:        req.AccountType,
		Amount:      req.Ammount,
		Status:      "1",
	}

	newAccount, err := s.repo.Create(account)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToNewAccountResponseDTO()
	return &response, nil
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	if req.IsWithdrawl() {
		acc, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}

		if !acc.CanWithDrawAmount(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}

	t := domain.Transaction{
		AccountId: req.AccountId,
		Amount:    req.Amount,
		Type:      domain.TransactionType(req.TransactionType),
		Date:      time.Now().Format("2006-01-02 15:04:05"),
	}

	transaction, appErr := s.repo.SaveTransaction(t)
	if appErr != nil {
		return nil, appErr
	}

	response := transaction.ToTransactionResponseDTO()
	return &response, nil
}
