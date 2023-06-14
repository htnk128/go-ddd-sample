package interactor

import (
	"context"
	"fmt"

	"github.com/htnk128/go-ddd-sample/pkg/account/domain/model"
	"github.com/htnk128/go-ddd-sample/pkg/account/domain/repository"
	"github.com/htnk128/go-ddd-sample/pkg/account/usecase/inputport/command"
	sharedUseCase "github.com/htnk128/go-ddd-sample/pkg/shared/usecase"
)

type AccountInteractor struct {
	accountRepository  repository.AccountRepository
	addressBookService model.AddressBookService
}

func NewAccountInteractor(accountRepository repository.AccountRepository, addressBookService model.AddressBookService) *AccountInteractor {
	return &AccountInteractor{
		accountRepository:  accountRepository,
		addressBookService: addressBookService,
	}
}

func (ai *AccountInteractor) Find(ctx context.Context, command command.FindAccountCommand) (*model.Account, error) {
	id, err := model.NewAccountID(command.AccountID)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}

	a, err := ai.accountRepository.Find(ctx, *id, false)
	if err != nil {
		return nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
	}
	if a == nil {
		return nil, sharedUseCase.NewNotFoundError(fmt.Sprintf("account not found. (accountId=%s)", id.String()))
	}

	return a, nil
}
func (ai *AccountInteractor) FindAll(ctx context.Context, command command.FindAllAccountCommand) (int64, []*model.Account, error) {
	cnt, err := ai.accountRepository.Count(ctx)
	if err != nil {
		return 0, nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
	}

	as, err := ai.accountRepository.FindAll(ctx, command.Limit, command.Offset)
	if err != nil {
		return 0, nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
	}

	return cnt, as, nil
}

func (ai *AccountInteractor) Create(ctx context.Context, command command.CreateAccountCommand) (*model.Account, error) {
	id := ai.accountRepository.NextAccountID()
	name, err := model.NewAccountName(command.Name)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}
	namePronunciation, err := model.NewAccountNamePronunciation(command.NamePronunciation)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}
	email, err := model.NewAccountEmail(command.Email)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}
	password, err := model.NewAccountPasswordWithHash(command.Password, *id)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}

	// TODO トランザクション

	a := model.NewAccount(*id, *name, *namePronunciation, *email, *password)
	err = ai.accountRepository.Add(ctx, a)
	if err != nil {
		return nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
	}

	// TODO ドメインイベントのパブリッシュを実装
	return a, nil
}
func (ai *AccountInteractor) Update(ctx context.Context, command command.UpdateAccountCommand) (*model.Account, error) {
	id, err := model.NewAccountID(command.AccountID)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}
	var name *model.AccountName
	if command.Name != nil {
		name, err = model.NewAccountName(*command.Name)
		if err != nil {
			return nil, sharedUseCase.NewInvalidRequestError(err.Error())
		}
	}
	var namePronunciation *model.AccountNamePronunciation
	if command.NamePronunciation != nil {
		namePronunciation, err = model.NewAccountNamePronunciation(*command.NamePronunciation)
		if err != nil {
			return nil, sharedUseCase.NewInvalidRequestError(err.Error())
		}
	}
	var email *model.AccountEmail
	if command.Email != nil {
		email, err = model.NewAccountEmail(*command.Email)
		if err != nil {
			return nil, sharedUseCase.NewInvalidRequestError(err.Error())
		}
	}
	var password *model.AccountPassword
	if command.Password != nil {
		password, err = model.NewAccountPasswordWithHash(*command.Password, *id)
		if err != nil {
			return nil, sharedUseCase.NewInvalidRequestError(err.Error())
		}
	}

	// TODO トランザクション

	a, err := ai.accountRepository.Find(ctx, *id, true)
	if err != nil {
		return nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
	}
	err = a.Update(name, namePronunciation, email, password)
	if err != nil {
		return nil, sharedUseCase.NewInvalidDataStateError("account has been deleted.")
	}
	cnt, err := ai.accountRepository.Set(ctx, a)
	if err != nil {
		return nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
	}
	if cnt > 0 {
		// TODO ドメインイベントのパブリッシュを実装
		return a, nil
	}

	return nil, sharedUseCase.NewUpdateFailureError(fmt.Sprintf("account update failure. (accountId=%s)", id.String()))
}

func (ai *AccountInteractor) Delete(ctx context.Context, command command.DeleteAccountCommand) (*model.Account, error) {
	id, err := model.NewAccountID(command.AccountID)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}

	// TODO トランザクション

	a, err := ai.accountRepository.Find(ctx, *id, true)
	if err != nil {
		return nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
	}
	a.Delete()
	cnt, err := ai.accountRepository.Set(ctx, a)
	if err != nil {
		return nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
	}
	if cnt > 0 {
		ab, err := ai.addressBookService.Find(*id)
		if err != nil {
			return nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
		}
		for _, aaa := range ab.AvailableAccountAddresses() {
			err = ai.addressBookService.Remove(aaa.ID)
			if err != nil {
				return nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
			}
		}

		// TODO ドメインイベントのパブリッシュを実装
		return a, nil
	}

	return nil, sharedUseCase.NewUpdateFailureError(fmt.Sprintf("account update failure. (accountId=%s)", id.String()))
}
