package repository

import (
	"context"

	"github.com/htnk128/go-ddd-sample/pkg/address/domain/model"
)

type AddressRepository interface {
	Find(ctx context.Context, addressID model.AddressID, lock bool) (*model.Address, error)
	FindAll(ctx context.Context, ownerID model.OwnerID) ([]*model.Address, error)
	Count(ctx context.Context) (int64, error)
	Add(ctx context.Context, address *model.Address) error
	Set(ctx context.Context, address *model.Address) (int64, error)
	NextAddressID() *model.AddressID
}
