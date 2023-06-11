package model

import (
	"time"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type AddressEvent struct {
	domain.Event

	Type       AddressEventType
	Address    *Address
	occurredOn time.Time
}

func (ae *AddressEvent) OccurredOn() time.Time {
	return ae.occurredOn
}

func (ae *AddressEvent) SameEventAs(other *AddressEvent) bool {
	if ae.Type != other.Type {
		return false
	}
	if ae.occurredOn != other.occurredOn {
		return false
	}
	return true
}

type AddressEventType string

const (
	AddressCreated AddressEventType = "address.created"
	AddressUpdated AddressEventType = "address.updated"
	AddressDeleted AddressEventType = "address.deleted"
)

func (t *AddressEventType) SameValueAs(other *AddressEventType) bool {
	return t == other
}

func NewAddressCreated(account *Address) *AddressEvent {
	return newAddressEvent(AddressCreated, account)
}

func NewAddressUpdated(account *Address) *AddressEvent {
	return newAddressEvent(AddressUpdated, account)
}

func NewAddressDeleted(account *Address) *AddressEvent {
	return newAddressEvent(AddressDeleted, account)
}

func newAddressEvent(accountEventType AddressEventType, account *Address) *AddressEvent {
	return &AddressEvent{
		Type:       accountEventType,
		Address:    account,
		occurredOn: time.Now(),
	}
}
