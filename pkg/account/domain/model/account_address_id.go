package model

import (
	"fmt"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type AccountAddressID struct {
	*domain.SomeIdentity
}

func NewAccountAddressID(id string) (*AccountAddressID, error) {
	if domain.SomeIdentityMinLength <= len(id) && len(id) <= domain.SomeIdentityMaxLength && domain.SomeIdentityRegexp.MatchString(id) {
		return &AccountAddressID{SomeIdentity: domain.NewSomeIdentity(id)}, nil
	}

	return nil, errors.New(fmt.Sprintf("Account address ID must be %d characters or less and alphanumeric, hyphen, underscore.", domain.SomeIdentityMaxLength))
}
