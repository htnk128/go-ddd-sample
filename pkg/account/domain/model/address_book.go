package model

import (
	"reflect"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type AddressBook struct {
	domain.ValueObject

	AllAccountAddresses []*AccountAddress
}

func (ab *AddressBook) AvailableAccountAddresses() []*AccountAddress {
	aas := make([]*AccountAddress, 0)
	for _, aa := range ab.AllAccountAddresses {
		if aa.IsAvailable() {
			aas = append(aas, aa)
		}
	}
	return aas
}

func (ab *AddressBook) Equals(other *AddressBook) bool {
	return ab.SameValueAs(other)
}

func (ab *AddressBook) SameValueAs(other *AddressBook) bool {
	return reflect.DeepEqual(ab.AllAccountAddresses, other.AllAccountAddresses)
}

func NewAddressBook(allAccountAddresses []*AccountAddress) *AddressBook {
	return &AddressBook{AllAccountAddresses: allAccountAddresses}
}
