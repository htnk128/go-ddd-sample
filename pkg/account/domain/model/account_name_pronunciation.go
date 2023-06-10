package model

import (
	"fmt"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type AccountNamePronunciation struct {
	*domain.SomeValueObject[string]
}

func NewAccountNamePronunciation(value string) (*AccountNamePronunciation, error) {
	const (
		minLength = 1
		maxLength = 100
	)

	if minLength <= len(value) && len(value) <= maxLength {
		return &AccountNamePronunciation{SomeValueObject: domain.NewSomeValueObject(value)}, nil
	}

	return nil, errors.New(fmt.Sprintf("AccountNamePronunciation must be %d characters or less.", maxLength))
}
