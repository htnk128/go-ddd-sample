package model

import (
	"fmt"

	"github.com/friendsofgo/errors"
	"github.com/google/uuid"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type AccountID struct {
	*domain.SomeIdentity
}

func GenerateAccountID() *AccountID {
	return &AccountID{SomeIdentity: domain.NewSomeIdentity(fmt.Sprintf("AC_%s", uuid.New().String()))}
}

func NewAccountID(id string) (*AccountID, error) {
	if domain.SomeIdentityMinLength <= len(id) && len(id) <= domain.SomeIdentityMaxLength && domain.SomeIdentityRegexp.MatchString(id) {
		return &AccountID{SomeIdentity: domain.NewSomeIdentity(id)}, nil
	}

	return nil, errors.New(fmt.Sprintf("Account ID must be %d characters or less and alphanumeric, hyphen, underscore.", domain.SomeIdentityMaxLength))
}
