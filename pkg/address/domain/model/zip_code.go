package model

import (
	"fmt"
	"regexp"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type ZipCode struct {
	*domain.SomeValueObject[string]
}

func (zc *ZipCode) Equals(other *ZipCode) bool {
	return zc.SameValueAs(other)
}

func (zc *ZipCode) SameValueAs(other *ZipCode) bool {
	return zc.Value() == other.Value()
}

const zipCodePattern = "[0-9A-Za-z]+"

var zipCodeRegexp = regexp.MustCompile(zipCodePattern)

func NewZipCode(value string) (*ZipCode, error) {
	const (
		minLength = 1
		maxLength = 50
	)

	if minLength <= len(value) && len(value) <= maxLength && zipCodeRegexp.MatchString(value) {
		return &ZipCode{SomeValueObject: domain.NewSomeValueObject(value)}, nil
	}

	return nil, errors.New(fmt.Sprintf("Zip code must be %d characters or less and alphanumeric.", maxLength))
}
