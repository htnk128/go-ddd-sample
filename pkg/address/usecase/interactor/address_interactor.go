package interactor

import (
	"context"
	"fmt"

	"github.com/htnk128/go-ddd-sample/pkg/address/domain/model"
	"github.com/htnk128/go-ddd-sample/pkg/address/domain/repository"
	"github.com/htnk128/go-ddd-sample/pkg/address/usecase/inputport/command"
	sharedUseCase "github.com/htnk128/go-ddd-sample/pkg/shared/usecase"
)

type AddressInteractor struct {
	addressRepository repository.AddressRepository
	ownerService      model.OwnerService
}

func NewAddressInteractor(addressRepository repository.AddressRepository, ownerService model.OwnerService) *AddressInteractor {
	return &AddressInteractor{
		addressRepository: addressRepository,
		ownerService:      ownerService,
	}
}

func (ai *AddressInteractor) Find(ctx context.Context, command command.FindAddressCommand) (*model.Address, error) {
	id, err := model.NewAddressID(command.AddressID)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}

	a, err := ai.addressRepository.Find(ctx, *id, false)
	if err != nil {
		return nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
	}
	if a == nil {
		return nil, sharedUseCase.NewNotFoundError(fmt.Sprintf("address not found. (addressId=%s)", id.String()))
	}

	return a, nil
}
func (ai *AddressInteractor) FindAll(ctx context.Context, command command.FindAllAddressCommand) ([]*model.Address, error) {
	id, err := model.NewOwnerID(command.OwnerID)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}

	as, err := ai.addressRepository.FindAll(ctx, *id)
	if err != nil {
		return nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
	}

	return as, nil
}

func (ai *AddressInteractor) Create(ctx context.Context, command command.CreateAddressCommand) (*model.Address, error) {
	id := ai.addressRepository.NextAddressID()
	ownerID, err := model.NewOwnerID(command.OwnerID)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}
	fullName, err := model.NewFullName(command.FullName)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}
	zipCode, err := model.NewZipCode(command.ZipCode)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}
	StateOrRegion, err := model.NewStateOrRegion(command.StateOrRegion)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}
	line1, err := model.NewLine1(command.Line1)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}
	var line2 *model.Line2
	if command.Line2 != nil {
		line2, err = model.NewLine2(*command.Line2)
		if err != nil {
			return nil, sharedUseCase.NewInvalidRequestError(err.Error())
		}
	}
	phoneNumber, err := model.NewPhoneNumber(command.PhoneNumber)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}

	// TODO トランザクション

	o, err := ai.ownerService.Find(*ownerID)
	if err != nil {
		return nil, sharedUseCase.NewNotFoundError(err.Error())
	}
	if o == nil || (o != nil && !o.IsAvailable()) {
		return nil, sharedUseCase.NewNotFoundError(fmt.Sprintf("owner not found. (ownerId=%s)", ownerID.String()))
	}

	a := model.NewAddress(*id, *ownerID, *fullName, *zipCode, *StateOrRegion, *line1, line2, *phoneNumber)
	err = ai.addressRepository.Add(ctx, a)
	if err != nil {
		return nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
	}

	// TODO ドメインイベントのパブリッシュを実装
	return a, nil
}

func (ai *AddressInteractor) Update(ctx context.Context, command command.UpdateAddressCommand) (*model.Address, error) {
	id, err := model.NewAddressID(command.AddressID)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}
	var fullName *model.FullName
	if command.FullName != nil {
		fullName, err = model.NewFullName(*command.FullName)
		if err != nil {
			return nil, sharedUseCase.NewInvalidRequestError(err.Error())
		}
	}
	var zipCode *model.ZipCode
	if command.ZipCode != nil {
		zipCode, err = model.NewZipCode(*command.ZipCode)
		if err != nil {
			return nil, sharedUseCase.NewInvalidRequestError(err.Error())
		}
	}
	var stateOrRegion *model.StateOrRegion
	if command.StateOrRegion != nil {
		stateOrRegion, err = model.NewStateOrRegion(*command.StateOrRegion)
		if err != nil {
			return nil, sharedUseCase.NewInvalidRequestError(err.Error())
		}
	}
	var line1 *model.Line1
	if command.Line1 != nil {
		line1, err = model.NewLine1(*command.Line1)
		if err != nil {
			return nil, sharedUseCase.NewInvalidRequestError(err.Error())
		}
	}
	var line2 *model.Line2
	if command.Line2 != nil {
		line2, err = model.NewLine2(*command.Line2)
		if err != nil {
			return nil, sharedUseCase.NewInvalidRequestError(err.Error())
		}
	}
	var phoneNumber *model.PhoneNumber
	if command.PhoneNumber != nil {
		phoneNumber, err = model.NewPhoneNumber(*command.PhoneNumber)
		if err != nil {
			return nil, sharedUseCase.NewInvalidRequestError(err.Error())
		}
	}

	// TODO トランザクション

	a, err := ai.addressRepository.Find(ctx, *id, true)
	if err != nil {
		return nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
	}
	if a == nil {
		return nil, sharedUseCase.NewNotFoundError(fmt.Sprintf("address not found. (addressId=%s)", id.String()))
	}
	err = a.Update(fullName, zipCode, stateOrRegion, line1, line2, phoneNumber)
	if err != nil {
		return nil, sharedUseCase.NewInvalidDataStateError("address has been deleted.")
	}
	cnt, err := ai.addressRepository.Set(ctx, a)
	if err != nil {
		return nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
	}
	if cnt > 0 {
		// TODO ドメインイベントのパブリッシュを実装
		return a, nil
	}

	return nil, sharedUseCase.NewUpdateFailureError(fmt.Sprintf("address update failure. (addressId=%s)", id.String()))
}

func (ai *AddressInteractor) Delete(ctx context.Context, command command.DeleteAddressCommand) (*model.Address, error) {
	id, err := model.NewAddressID(command.AddressID)
	if err != nil {
		return nil, sharedUseCase.NewInvalidRequestError(err.Error())
	}

	// TODO トランザクション

	a, err := ai.addressRepository.Find(ctx, *id, true)
	if err != nil {
		return nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
	}
	if a == nil {
		return nil, sharedUseCase.NewNotFoundError(fmt.Sprintf("address not found. (addressId=%s)", id.String()))
	}
	a.Delete()
	cnt, err := ai.addressRepository.Set(ctx, a)
	if err != nil {
		return nil, sharedUseCase.NewServerError("internal server error. " + err.Error())
	}
	if cnt > 0 {
		// TODO ドメインイベントのパブリッシュを実装
		return a, nil
	}

	return nil, sharedUseCase.NewUpdateFailureError(fmt.Sprintf("address update failure. (addressId=%s)", id.String()))
}
