package inputport

import (
	"context"

	"github.com/htnk128/go-ddd-sample/pkg/account/domain/model"
	"github.com/htnk128/go-ddd-sample/pkg/account/domain/repository"
	"github.com/htnk128/go-ddd-sample/pkg/account/usecase/inputport/command"
	"github.com/htnk128/go-ddd-sample/pkg/account/usecase/interactor"
)

type AccountUseCase interface {
	Find(ctx context.Context, command command.FindAccountCommand) (*model.Account, error)
	FindAll(ctx context.Context, command command.FindAllAccountCommand) (int64, []*model.Account, error)
	Create(ctx context.Context, command command.CreateAccountCommand) (*model.Account, error)
	Update(ctx context.Context, command command.UpdateAccountCommand) (*model.Account, error)
	Delete(ctx context.Context, command command.DeleteAccountCommand) (*model.Account, error)
}

func NewAccountUseCase(accountRepository repository.AccountRepository, addressBookService model.AddressBookService) AccountUseCase {
	return interactor.NewAccountInteractor(accountRepository, addressBookService)
}
