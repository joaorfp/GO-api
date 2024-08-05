package repository

import (
	"github.com/joaorfp/GO-api/internal/model"
)

type BankRepository interface {
	GetAll() []*model.BankAccount
	GetByUserId(userId int) (*model.BankAccount, bool)
	Create(account *model.BankAccount) *model.BankAccount
	UpdateByUserId(userId int, account *model.BankAccount) (*model.BankAccount, bool)
	DeleteByUserId(userId int) bool
}

type InMemoryBankRepository struct {
	bankAccounts []*model.BankAccount
}

func NewInMemoryBankRepository() *InMemoryBankRepository {
    return &InMemoryBankRepository{
		bankAccounts: []*model.BankAccount{},
    }
}

func (r *InMemoryBankRepository) GetAll() []*model.BankAccount {
    return r.bankAccounts
}


func (r *InMemoryBankRepository) GetByUserId(userId int) (*model.BankAccount, bool) {
    for _, account := range r.bankAccounts {
        if account.UserID == userId {
            return account, true
        }
    }
    return nil, false
}

func (r *InMemoryBankRepository) Create(account *model.BankAccount) *model.BankAccount {
    r.bankAccounts = append(r.bankAccounts, account)
    return account
}

func (r *InMemoryBankRepository) UpdateByUserId(userId int, account *model.BankAccount) (*model.BankAccount, bool) {
    for i, a := range r.bankAccounts {
        if a.UserID == userId {
            r.bankAccounts[i] = account
            return account, true
        }
    }
    return nil, false
}

func (r *InMemoryBankRepository) DeleteByUserId(userId int) bool {
    for i, account := range r.bankAccounts {
        if account.UserID == userId {
            r.bankAccounts = append(r.bankAccounts[:i], r.bankAccounts[i+1:]...)
            return true
        }
    }
    return false
}
