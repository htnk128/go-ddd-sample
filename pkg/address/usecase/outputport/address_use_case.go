package outputport

import (
	"github.com/htnk128/go-ddd-sample/pkg/address/domain/model"
	"github.com/htnk128/go-ddd-sample/pkg/address/usecase/outputport/dto"
)

type AddressUseCase interface {
	DTO(address *model.Address) *dto.AddressDTO
}
