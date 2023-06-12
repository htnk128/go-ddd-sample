package model

import (
	"fmt"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type FullName struct {
	*domain.SomeValueObject[string]
}

func NewFullName(value string) (*FullName, error) {
	const (
		minLength = 1
		maxLength = 100
	)

	if minLength <= len(value) && len(value) <= maxLength {
		return &FullName{SomeValueObject: domain.NewSomeValueObject(value)}, nil
	}

	return nil, errors.New(fmt.Sprintf("Full name must be %d characters or less.", maxLength))
}
