package repository

import (
	"context"

	"github.com/htnk128/go-ddd-sample/pkg/account/domain/model"
)

type AccountRepository interface {
	Find(ctx context.Context, accountID model.AccountID, lock bool) (*model.Account, error)
	FindAll(ctx context.Context, limit, offset int) ([]*model.Account, error)
	Count(ctx context.Context) (int64, error)
	Add(ctx context.Context, account *model.Account) error
	Set(ctx context.Context, account *model.Account) (int64, error)
	NextAccountID() *model.AccountID
}
