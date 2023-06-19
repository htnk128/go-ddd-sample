package model

import (
	"fmt"

	"github.com/friendsofgo/errors"
	"github.com/google/uuid"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type OwnerID struct {
	*domain.SomeIdentity
}

func (oi *OwnerID) Equals(other *OwnerID) bool {
	return oi.SameValueAs(other)
}

func (oi *OwnerID) SameValueAs(other *OwnerID) bool {
	return oi.ID() == other.ID()
}

func GenerateOwnerID() *OwnerID {
	return &OwnerID{SomeIdentity: domain.NewSomeIdentity(fmt.Sprintf("AC_%s", uuid.New().String()))}
}

func NewOwnerID(id string) (*OwnerID, error) {
	if domain.SomeIdentityMinLength <= len(id) && len(id) <= domain.SomeIdentityMaxLength && domain.SomeIdentityRegexp.MatchString(id) {
		return &OwnerID{SomeIdentity: domain.NewSomeIdentity(id)}, nil
	}

	return nil, errors.New(fmt.Sprintf("Owner ID must be %d characters or less and alphanumeric, hyphen, underscore.", domain.SomeIdentityMaxLength))
}
