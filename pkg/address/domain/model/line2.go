package model

import (
	"fmt"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type Line2 struct {
	*domain.SomeValueObject[string]
}

func NewLine2(value string) (*Line2, error) {
	const (
		minLength = 1
		maxLength = 100
	)

	if minLength <= len(value) && len(value) <= maxLength {
		return &Line2{SomeValueObject: domain.NewSomeValueObject(value)}, nil
	}

	return nil, errors.New(fmt.Sprintf("Line2 must be %d characters or less.", maxLength))
}
