package outputport

import (
	"github.com/htnk128/go-ddd-sample/pkg/account/domain/model"
	"github.com/htnk128/go-ddd-sample/pkg/account/usecase/outputport/dto"
	sharedDTO "github.com/htnk128/go-ddd-sample/pkg/shared/usecase/outputport/dto"
)

type AccountUseCase interface {
	DTO(account *model.Account) *dto.AccountDTO
	PaginationDTO(accounts []*model.Account, count, limit, offset int) *sharedDTO.PaginationDTO[dto.AccountDTO]
}
