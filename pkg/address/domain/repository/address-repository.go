package repository

import (
	"github.com/htnk128/go-ddd-sample/pkg/address/domain/model"
)

type AddressRepository interface {
	Find(addressID model.AddressID, lock bool) (*model.Address, error)
	FindAll(ownerID model.OwnerID) ([]*model.Address, error)
	Count() (int, error)
	Add(address *model.Address) error
	Set(address *model.Address) error
	Remove(address *model.Address) error
	NextAddressID() *model.AddressID
}
