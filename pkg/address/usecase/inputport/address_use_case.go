package inputport

import (
	"context"

	"github.com/htnk128/go-ddd-sample/pkg/address/domain/model"
	"github.com/htnk128/go-ddd-sample/pkg/address/domain/repository"
	"github.com/htnk128/go-ddd-sample/pkg/address/usecase/inputport/command"
	"github.com/htnk128/go-ddd-sample/pkg/address/usecase/interactor"
)

type AddressUseCase interface {
	Find(ctx context.Context, command command.FindAddressCommand) (*model.Address, error)
	FindAll(ctx context.Context, command command.FindAllAddressCommand) ([]*model.Address, error)
	Create(ctx context.Context, command command.CreateAddressCommand) (*model.Address, error)
	Update(ctx context.Context, command command.UpdateAddressCommand) (*model.Address, error)
	Delete(ctx context.Context, command command.DeleteAddressCommand) (*model.Address, error)
}

func NewAddressUseCase(addressRepository repository.AddressRepository, ownerService model.OwnerService) AddressUseCase {
	return interactor.NewAddressInteractor(addressRepository, ownerService)
}
