package model

import (
	"fmt"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type AccountName struct {
	*domain.SomeValueObject[string]
}

func (an *AccountName) Equals(other *AccountName) bool {
	return an.SameValueAs(other)
}

func (an *AccountName) SameValueAs(other *AccountName) bool {
	return an.Value() == other.Value()
}

func NewAccountName(value string) (*AccountName, error) {
	const (
		minLength = 1
		maxLength = 100
	)

	if minLength <= len(value) && len(value) <= maxLength {
		return &AccountName{SomeValueObject: domain.NewSomeValueObject(value)}, nil
	}

	return nil, errors.New(fmt.Sprintf("AccountName must be %d characters or less.", maxLength))
}
