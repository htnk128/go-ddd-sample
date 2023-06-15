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

func NewAddressCreated(address *Address) *AddressEvent {
	return newAddressEvent(AddressCreated, address)
}

func NewAddressUpdated(address *Address) *AddressEvent {
	return newAddressEvent(AddressUpdated, address)
}

func NewAddressDeleted(address *Address) *AddressEvent {
	return newAddressEvent(AddressDeleted, address)
}

func newAddressEvent(addressEventType AddressEventType, address *Address) *AddressEvent {
	return &AddressEvent{
		Type:       addressEventType,
		Address:    address,
		occurredOn: time.Now(),
	}
}
