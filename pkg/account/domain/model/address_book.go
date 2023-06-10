package model

import (
	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
	"reflect"
)

type AddressBook struct {
	domain.ValueObject

	AllAccountAddresses []*AccountAddress
}

func (ab *AddressBook) AvailableAccountAddresses() []*AccountAddress {
	aaa := make([]*AccountAddress, 0)
	for _, aa := range ab.AllAccountAddresses {
		if aa.IsAvailable() {
			aaa = append(aaa, aa)
		}
	}
	return aaa
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
