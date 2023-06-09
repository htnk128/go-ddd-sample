package model

import (
	"fmt"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type AccountEmail struct {
	*domain.SomeValueObject[string]
}

func NewAccountEmail(value string) (*AccountEmail, error) {
	const (
		minLength = 1
		maxLength = 100
	)

	if minLength <= len(value) && len(value) <= maxLength {
		return &AccountEmail{SomeValueObject: &domain.SomeValueObject[string]{Value: value}}, nil
	}

	return nil, errors.New(fmt.Sprintf("AccountEmail must be %d characters or less.", maxLength))
}
