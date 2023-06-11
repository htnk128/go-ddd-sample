package model

import (
	"fmt"
	"time"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type Address struct {
	domain.Entity

	ID             AddressID
	OwnerID        OwnerID
	FullName       FullName
	ZipCode        ZipCode
	StateOrRegion  StateOrRegion
	Line1          Line1
	Line2          *Line2
	PhoneNumber    PhoneNumber
	CreatedAt      time.Time
	DeletedAt      *time.Time
	UpdatedAt      time.Time
	OccurredEvents []*AddressEvent
}

func (a *Address) IsDeleted() bool {
	return a.DeletedAt != nil
}

func (a *Address) Update(
	fullName *FullName,
	zipCode *ZipCode,
	stateOrRegion *StateOrRegion,
	line1 *Line1,
	line2 *Line2,
	phoneNumber *PhoneNumber,
) error {
	if a.IsDeleted() {
		return errors.New("Address has been deleted.")
	}

	if fullName != nil {
		a.FullName = *fullName
	}
	if zipCode != nil {
		a.ZipCode = *zipCode
	}
	if stateOrRegion != nil {
		a.StateOrRegion = *stateOrRegion
	}
	if line1 != nil {
		a.Line1 = *line1
	}
	if line2 != nil {
		a.Line2 = line2
	}
	if phoneNumber != nil {
		a.PhoneNumber = *phoneNumber
	}

	a.UpdatedAt = time.Now()
	a.OccurredEvents = append(a.OccurredEvents, NewAddressUpdated(a))

	return nil
}

func (a *Address) Delete() {
	if a.IsDeleted() {
		return
	}

	now := time.Now()
	a.DeletedAt = &now
	a.UpdatedAt = now
	a.OccurredEvents = append(a.OccurredEvents, NewAddressDeleted(a))
}

func (a *Address) Equals(other *Address) bool {
	return a.SameIdentityAs(other)
}

func (a *Address) SameIdentityAs(other *Address) bool {
	return a.ID == other.ID
}

func (a *Address) String() string {
	updatedAt := a.UpdatedAt.Format(time.RFC3339)
	createdAt := a.CreatedAt.Format(time.RFC3339)
	var deletedAt = "nil"
	if a.DeletedAt != nil {
		deletedAt = a.DeletedAt.Format(time.RFC3339)
	}

	return fmt.Sprintf("ID=%v, OwnerID=%v, FullName=%v, ZipCode=%v, StateOrRegion=%v, Line1=%v, Line2=%v, PhoneNumber=%v, CreatedAt=%v, DeletedAt=%v, UpdatedAt=%v",
		a.ID, a.OwnerID, a.FullName, a.ZipCode, a.StateOrRegion, a.Line1, a.Line2, a.PhoneNumber, createdAt, deletedAt, updatedAt)
}

func NewAddress(
	id AddressID,
	ownerID OwnerID,
	fullName FullName,
	zipCode ZipCode,
	stateOrRegion StateOrRegion,
	line1 Line1,
	line2 *Line2,
	phoneNumber PhoneNumber,
) *Address {
	now := time.Now()
	a := &Address{
		ID:            id,
		OwnerID:       ownerID,
		FullName:      fullName,
		ZipCode:       zipCode,
		StateOrRegion: stateOrRegion,
		Line1:         line1,
		Line2:         line2,
		PhoneNumber:   phoneNumber,
		CreatedAt:     now,
		DeletedAt:     nil,
		UpdatedAt:     now,
	}
	a.OccurredEvents = []*AddressEvent{NewAddressCreated(a)}

	return a
}
