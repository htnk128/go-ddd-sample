package model

import (
	"fmt"
	"time"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type Account struct {
	domain.Entity

	ID                AccountID
	Name              AccountName
	NamePronunciation AccountNamePronunciation
	Email             AccountEmail
	Password          AccountPassword
	CreatedAt         time.Time
	DeletedAt         *time.Time
	UpdatedAt         time.Time
	OccurredEvents    []*AccountEvent
}

func (a *Account) IsDeleted() bool {
	return a.DeletedAt != nil
}

func (a *Account) Update(
	name *AccountName,
	namePronunciation *AccountNamePronunciation,
	email *AccountEmail,
	password *AccountPassword,
) error {
	if a.IsDeleted() {
		return errors.New("Account has been deleted.")
	}

	if name != nil {
		a.Name = *name
	}
	if namePronunciation != nil {
		a.NamePronunciation = *namePronunciation
	}
	if email != nil {
		a.Email = *email
	}
	if password != nil {
		a.Password = *password
	}

	a.UpdatedAt = time.Now()
	a.OccurredEvents = append(a.OccurredEvents, NewAccountUpdated(a))

	return nil
}

func (a *Account) Delete() {
	if a.IsDeleted() {
		return
	}

	now := time.Now()
	a.DeletedAt = &now
	a.UpdatedAt = now
	a.OccurredEvents = append(a.OccurredEvents, NewAccountDeleted(a))
}

func (a *Account) Equals(other *Account) bool {
	return a.SameIdentityAs(other)
}

func (a *Account) SameIdentityAs(other *Account) bool {
	return a.ID == other.ID
}

func (a *Account) String() string {
	updatedAt := a.UpdatedAt.Format(time.RFC3339)
	createdAt := a.CreatedAt.Format(time.RFC3339)
	var deletedAt = "nil"
	if a.DeletedAt != nil {
		deletedAt = a.DeletedAt.Format(time.RFC3339)
	}

	return fmt.Sprintf("ID=%v, Name=%v, NamePronunciation=%v, Email=%v, Password=%v, CreatedAt=%v, DeletedAt=%v, UpdatedAt=%v",
		a.ID, a.Name, a.NamePronunciation, a.Email, a.Password, createdAt, deletedAt, updatedAt)
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
		ID:                id,
		Name:              name,
		NamePronunciation: namePronunciation,
		Email:             email,
		Password:          password,
		CreatedAt:         now,
		DeletedAt:         nil,
		UpdatedAt:         now,
	}
	a.OccurredEvents = []*AccountEvent{NewAccountCreated(a)}

	return a
}
