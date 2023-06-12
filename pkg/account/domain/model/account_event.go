package model

import (
	"time"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type AccountEvent struct {
	domain.Event

	Type       AccountEventType
	Account    *Account
	occurredOn time.Time
}

func (ae *AccountEvent) OccurredOn() time.Time {
	return ae.occurredOn
}

func (ae *AccountEvent) SameEventAs(other *AccountEvent) bool {
	if ae.Type != other.Type {
		return false
	}
	if ae.occurredOn != other.occurredOn {
		return false
	}
	return true
}

type AccountEventType string

const (
	AccountCreated AccountEventType = "account.created"
	AccountUpdated AccountEventType = "account.updated"
	AccountDeleted AccountEventType = "account.deleted"
)

func (t *AccountEventType) SameValueAs(other *AccountEventType) bool {
	return t == other
}

func NewAccountCreated(account *Account) *AccountEvent {
	return newAccountEvent(AccountCreated, account)
}

func NewAccountUpdated(account *Account) *AccountEvent {
	return newAccountEvent(AccountUpdated, account)
}

func NewAccountDeleted(account *Account) *AccountEvent {
	return newAccountEvent(AccountDeleted, account)
}

func newAccountEvent(accountEventType AccountEventType, account *Account) *AccountEvent {
	return &AccountEvent{
		Type:       accountEventType,
		Account:    account,
		occurredOn: time.Now(),
	}
}
