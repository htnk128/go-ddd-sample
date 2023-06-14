package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	gatewayModel "github.com/htnk128/go-ddd-sample/pkg/account/adapter/gateway/db/model"
	"github.com/htnk128/go-ddd-sample/pkg/account/domain/model"
	"github.com/htnk128/go-ddd-sample/pkg/account/domain/repository"
)

type accountSQLBoilerRepository struct {
	db *sql.DB
}

func newUserRepository(db *sql.DB) repository.AccountRepository {
	return &accountSQLBoilerRepository{db: db}
}

func (asr *accountSQLBoilerRepository) Find(ctx context.Context, accountID model.AccountID, lock bool) (*model.Account, error) {
	q := []qm.QueryMod{
		qm.Where("account_id = ?", accountID.ID()),
	}
	if lock {
		q = append(q, qm.For("UPDATE"))
	}
	ac, err := gatewayModel.Accounts(q...).One(ctx, asr.db)
	if err != nil {
		return nil, err
	}
	m, err := accountFromGatewayModel(ac)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (asr *accountSQLBoilerRepository) FindAll(ctx context.Context, limit, offset int) ([]*model.Account, error) {
	acs, err := gatewayModel.Accounts(
		qm.OrderBy(gatewayModel.AccountColumns.CreatedAt),
		qm.Limit(limit),
		qm.Offset(offset*limit),
	).All(ctx, asr.db)
	if err != nil {
		return nil, err
	}

	all := make([]*model.Account, len(acs))
	for i, ac := range acs {
		m, err := accountFromGatewayModel(ac)
		if err != nil {
			return nil, err
		}
		all[i] = m
	}

	return all, nil
}

func (asr *accountSQLBoilerRepository) Count(ctx context.Context) (int64, error) {
	cnt, err := gatewayModel.Accounts().Count(ctx, asr.db)
	if err != nil {
		return 0, err
	}

	return cnt, nil
}

func (asr *accountSQLBoilerRepository) Add(ctx context.Context, account *model.Account) error {
	return accountFromDomainModel(account).Insert(ctx, asr.db, boil.Infer())
}

func (asr *accountSQLBoilerRepository) Set(ctx context.Context, account *model.Account) (int64, error) {
	return accountFromDomainModel(account).Update(ctx, asr.db, boil.Infer())
}

func (asr *accountSQLBoilerRepository) NextAccountID() *model.AccountID {
	return model.GenerateAccountID()
}

func accountFromGatewayModel(account *gatewayModel.Account) (*model.Account, error) {
	id, err := model.NewAccountID(account.AccountID)
	if err != nil {
		return nil, err
	}
	name, err := model.NewAccountName(account.Name)
	if err != nil {
		return nil, err
	}
	namePronunciation, err := model.NewAccountNamePronunciation(account.NamePronunciation)
	if err != nil {
		return nil, err
	}
	email, err := model.NewAccountEmail(account.Email)
	if err != nil {
		return nil, err
	}
	password := model.NewAccountPassword(account.Password)
	var deletedAt *time.Time
	if account.DeletedAt.Valid {
		deletedAt = &account.DeletedAt.Time
	}

	return &model.Account{
		ID:                *id,
		Name:              *name,
		NamePronunciation: *namePronunciation,
		Email:             *email,
		Password:          *password,
		CreatedAt:         account.CreatedAt,
		DeletedAt:         deletedAt,
		UpdatedAt:         account.UpdatedAt,
	}, nil
}

func accountFromDomainModel(account *model.Account) *gatewayModel.Account {
	return &gatewayModel.Account{
		AccountID:         account.ID.ID(),
		Name:              account.Name.Value(),
		NamePronunciation: account.NamePronunciation.Value(),
		Email:             account.Email.Value(),
		Password:          account.Password.Value(),
		CreatedAt:         account.CreatedAt,
		DeletedAt:         null.TimeFromPtr(account.DeletedAt),
		UpdatedAt:         account.UpdatedAt,
	}
}
