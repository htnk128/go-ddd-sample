package model

import (
	"fmt"
	"regexp"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type PhoneNumber struct {
	*domain.SomeValueObject[string]
}

func (pn *PhoneNumber) Equals(other *PhoneNumber) bool {
	return pn.SameValueAs(other)
}

func (pn *PhoneNumber) SameValueAs(other *PhoneNumber) bool {
	return pn.Value() == other.Value()
}

const PhoneNumberPattern = "[0-9]+"

var PhoneNumberRegexp = regexp.MustCompile(PhoneNumberPattern)

func NewPhoneNumber(value string) (*PhoneNumber, error) {
	const (
		minLength = 1
		maxLength = 50
	)

	if minLength <= len(value) && len(value) <= maxLength && PhoneNumberRegexp.MatchString(value) {
		return &PhoneNumber{SomeValueObject: domain.NewSomeValueObject(value)}, nil
	}

	return nil, errors.New(fmt.Sprintf("Phone number must be %d characters or less and numeric.", maxLength))
}
