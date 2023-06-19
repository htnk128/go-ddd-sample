package model

import (
	"fmt"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type StateOrRegion struct {
	*domain.SomeValueObject[string]
}

func (sor *StateOrRegion) Equals(other *StateOrRegion) bool {
	return sor.SameValueAs(other)
}

func (sor *StateOrRegion) SameValueAs(other *StateOrRegion) bool {
	return sor.Value() == other.Value()
}

func NewStateOrRegion(value string) (*StateOrRegion, error) {
	const (
		minLength = 1
		maxLength = 100
	)

	if minLength <= len(value) && len(value) <= maxLength {
		return &StateOrRegion{SomeValueObject: domain.NewSomeValueObject(value)}, nil
	}

	return nil, errors.New(fmt.Sprintf("State or region must be %d characters or less.", maxLength))
}
