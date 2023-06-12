package model

import (
	"fmt"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type City struct {
	*domain.SomeValueObject[string]
}

func NewCity(value string) (*City, error) {
	const (
		minLength = 1
		maxLength = 100
	)

	if minLength <= len(value) && len(value) <= maxLength {
		return &City{SomeValueObject: domain.NewSomeValueObject(value)}, nil
	}

	return nil, errors.New(fmt.Sprintf("City must be %d characters or less.", maxLength))
}
