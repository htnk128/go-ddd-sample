package repository

import (
	"github.com/htnk128/go-ddd-sample/pkg/account/domain/model"
)

type AccountRepository interface {
	Find(accountID model.AccountID, lock bool) (*model.Account, error)
	FindAll(limit, offset int) ([]*model.Account, error)
	Count() (int, error)
	Add(account *model.Account) error
	Set(account *model.Account) error
	Remove(account *model.Account) error
	NextAccountID() *model.AccountID
}
