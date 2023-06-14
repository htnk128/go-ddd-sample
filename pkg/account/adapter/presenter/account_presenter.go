package presenter

import (
	"github.com/htnk128/go-ddd-sample/pkg/account/domain/model"
	"github.com/htnk128/go-ddd-sample/pkg/account/usecase/outputport"
	"github.com/htnk128/go-ddd-sample/pkg/account/usecase/outputport/dto"
	sharedDTO "github.com/htnk128/go-ddd-sample/pkg/shared/usecase/outputport/dto"
)

type accountPresenter struct {
}

func NewAccountUseCase() outputport.AccountUseCase {
	return &accountPresenter{}
}

func (ap accountPresenter) DTO(account *model.Account) *dto.AccountDTO {
	var deletedAt *int64
	if account.DeletedAt != nil {
		d := account.DeletedAt.UnixMilli()
		deletedAt = &d
	}

	return &dto.AccountDTO{
		AccountID:         account.ID.ID(),
		Name:              account.Name.Value(),
		NamePronunciation: account.NamePronunciation.Value(),
		Email:             account.Email.Value(),
		Password:          account.Password.Value(),
		CreatedAt:         account.CreatedAt.UnixMilli(),
		DeletedAt:         deletedAt,
		UpdatedAt:         account.UpdatedAt.UnixMilli(),
	}
}

func (ap accountPresenter) PaginationDTO(accounts []*model.Account, count, limit, offset int) *sharedDTO.PaginationDTO[dto.AccountDTO] {
	data := make([]*dto.AccountDTO, len(accounts))
	for i, d := range accounts {
		data[i] = ap.DTO(d)
	}
	return &sharedDTO.PaginationDTO[dto.AccountDTO]{
		Count:  count,
		Limit:  limit,
		Offset: offset,
		Data:   data,
	}
}
