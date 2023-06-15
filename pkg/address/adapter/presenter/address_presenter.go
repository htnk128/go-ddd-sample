package presenter

import (
	"github.com/htnk128/go-ddd-sample/pkg/address/domain/model"
	"github.com/htnk128/go-ddd-sample/pkg/address/usecase/outputport"
	"github.com/htnk128/go-ddd-sample/pkg/address/usecase/outputport/dto"
)

type addressPresenter struct {
}

func NewAddressUseCase() outputport.AddressUseCase {
	return &addressPresenter{}
}

func (ap addressPresenter) DTO(address *model.Address) *dto.AddressDTO {
	var line2 *string
	if address.Line2 != nil {
		l2 := address.Line2.Value()
		line2 = &l2
	}
	var deletedAt *int64
	if address.DeletedAt != nil {
		d := address.DeletedAt.UnixMilli()
		deletedAt = &d
	}

	return &dto.AddressDTO{
		AddressID:     address.ID.ID(),
		OwnerID:       address.OwnerID.ID(),
		FullName:      address.FullName.Value(),
		ZipCode:       address.ZipCode.Value(),
		StateOrRegion: address.StateOrRegion.Value(),
		Line1:         address.Line1.Value(),
		Line2:         line2,
		PhoneNumber:   address.PhoneNumber.Value(),
		CreatedAt:     address.CreatedAt.UnixMilli(),
		DeletedAt:     deletedAt,
		UpdatedAt:     address.UpdatedAt.UnixMilli(),
	}
}
