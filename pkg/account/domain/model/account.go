package model

import (
	"fmt"
	"time"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type Account struct {
	domain.Entity

	id                AccountID
	name              AccountName
	namePronunciation AccountNamePronunciation
	email             AccountEmail
	password          AccountPassword
	createdAt         time.Time
	deletedAt         *time.Time
	updatedAt         time.Time
	occurredEvents    []*AccountEvent
}

func (a *Account) IsDeleted() bool {
	return a.deletedAt != nil
}

func (a *Account) Update(
	name *AccountName,
	namePronunciation *AccountNamePronunciation,
	email *AccountEmail,
	password *AccountPassword,
) (*Account, error) {
	if a.IsDeleted() {
		return nil, errors.New("Account has been deleted.")
	}

	if name != nil {
		a.name = *name
	}
	if namePronunciation != nil {
		a.namePronunciation = *namePronunciation
	}
	if email != nil {
		a.email = *email
	}
	if password != nil {
		a.password = *password
	}

	a.updatedAt = time.Now()
	a.occurredEvents = append(a.occurredEvents, NewAccountUpdated(a))

	return a, nil
}

func (a *Account) Delete() *Account {
	if a.IsDeleted() {
		return a
	}

	now := time.Now()
	a.deletedAt = &now
	a.updatedAt = now
	a.occurredEvents = append(a.occurredEvents, NewAccountDeleted(a))

	return a
}

func (a *Account) Equals(other *Account) bool {
	return a.SameIdentityAs(other)
}

func (a *Account) SameIdentityAs(other *Account) bool {
	return a.id == other.id
}

func (a *Account) String() string {
	updatedAt := a.updatedAt.Format(time.RFC3339)
	createdAt := a.createdAt.Format(time.RFC3339)
	var deletedAt = "nil"
	if a.deletedAt != nil {
		deletedAt = a.deletedAt.Format(time.RFC3339)
	}

	return fmt.Sprintf("id=%v, name=%v, namePronunciation=%v, email=%v, password=%v, createdAt=%v, deletedAt=%v, updatedAt=%v",
		a.id, a.name, a.namePronunciation, a.email, a.password, createdAt, deletedAt, updatedAt)
}

func NewAccount(
	id AccountID,
	name AccountName,
	namePronunciation AccountNamePronunciation,
	email AccountEmail,
	password AccountPassword,
) *Account {
	now := time.Now()
	a := &Account{
		id:                id,
		name:              name,
		namePronunciation: namePronunciation,
		email:             email,
		password:          password,
		createdAt:         now,
		deletedAt:         nil,
		updatedAt:         now,
	}
	a.occurredEvents = []*AccountEvent{NewAccountCreated(a)}

	return a
}
