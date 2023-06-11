package model

import (
	"fmt"

	"github.com/friendsofgo/errors"
	"github.com/google/uuid"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type AddressID struct {
	*domain.SomeIdentity
}

func GenerateAddressID() *AddressID {
	return &AddressID{SomeIdentity: domain.NewSomeIdentity(fmt.Sprintf("ADDR_%s", uuid.New().String()))}
}

func NewAddressID(id string) (*AddressID, error) {
	if domain.SomeIdentityMinLength <= len(id) && len(id) <= domain.SomeIdentityMaxLength && domain.SomeIdentityRegexp.MatchString(id) {
		return &AddressID{SomeIdentity: domain.NewSomeIdentity(id)}, nil
	}

	return nil, errors.New(fmt.Sprintf("Address ID must be %d characters or less and alphanumeric, hyphen, underscore.", domain.SomeIdentityMaxLength))
}
