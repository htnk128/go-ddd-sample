package model

import (
	"time"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type AccountAddress struct {
	domain.ValueObject

	ID        AccountAddressID
	DeletedAt *time.Time
}

func (aa *AccountAddress) IsAvailable() bool {
	return aa.DeletedAt == nil
}

func (aa *AccountAddress) Equals(other *AccountAddress) bool {
	return aa.SameValueAs(other)
}

func (aa *AccountAddress) SameValueAs(other *AccountAddress) bool {
	return aa.ID.ID() == other.ID.ID() && aa.DeletedAt == other.DeletedAt
}

func NewAccountAddress(id AccountAddressID, deletedAt *time.Time) *AccountAddress {
	return &AccountAddress{ID: id, DeletedAt: deletedAt}
}
