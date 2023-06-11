package model

import (
	"fmt"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type Line1 struct {
	*domain.SomeValueObject[string]
}

func NewLine1(value string) (*Line1, error) {
	const (
		minLength = 1
		maxLength = 100
	)

	if minLength <= len(value) && len(value) <= maxLength {
		return &Line1{SomeValueObject: domain.NewSomeValueObject(value)}, nil
	}

	return nil, errors.New(fmt.Sprintf("Line1 must be %d characters or less.", maxLength))
}
