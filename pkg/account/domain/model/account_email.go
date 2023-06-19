package model

import (
	"fmt"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type AccountEmail struct {
	*domain.SomeValueObject[string]
}

func (ae *AccountEmail) Equals(other *AccountEmail) bool {
	return ae.SameValueAs(other)
}

func (ae *AccountEmail) SameValueAs(other *AccountEmail) bool {
	return ae.Value() == other.Value()
}

func NewAccountEmail(value string) (*AccountEmail, error) {
	const (
		minLength = 1
		maxLength = 100
	)

	if minLength <= len(value) && len(value) <= maxLength {
		return &AccountEmail{SomeValueObject: domain.NewSomeValueObject(value)}, nil
	}

	return nil, errors.New(fmt.Sprintf("AccountEmail must be %d characters or less.", maxLength))
}
