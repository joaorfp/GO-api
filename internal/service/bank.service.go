package service

import (
	"github.com/joaorfp/GO-api/internal/model"
	"github.com/joaorfp/GO-api/internal/repository"
)

type BankService interface {
    GetAllAccounts() []*model.BankAccount
    GetAccountByUserId(userId int) (*model.BankAccount, bool)
    CreateAccount(account *model.BankAccount) *model.BankAccount
    UpdateAccountByUserId(userId int, account *model.BankAccount) (*model.BankAccount, bool)
    DeleteAccountByUserId(userId int) bool
}

type bankService struct {
    repo repository.BankRepository
}

func NewBankService(repo repository.BankRepository) BankService {
    return &bankService{repo: repo}
}

func (s *bankService) GetAllAccounts() []*model.BankAccount {
    return s.repo.GetAll()
}

func (s *bankService) GetAccountByUserId(userId int) (*model.BankAccount, bool) {
    return s.repo.GetByUserId(userId)
}

func (s *bankService) CreateAccount(account *model.BankAccount) *model.BankAccount {
    return s.repo.Create(account)
}

func (s *bankService) UpdateAccountByUserId(userId int, account *model.BankAccount) (*model.BankAccount, bool) {
    return s.repo.UpdateByUserId(userId, account)
}

func (s *bankService) DeleteAccountByUserId(userId int) bool {
    return s.repo.DeleteByUserId(userId)
}
