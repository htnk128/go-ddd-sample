package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewAddressBook(t *testing.T) {
	assert := require.New(t)

	id, _ := NewAccountAddressID("test_account_address_id")

	aa := NewAccountAddress(*id, nil)
	ab := NewAddressBook([]*AccountAddress{aa})
	assert.Equal(ab.AllAccountAddresses, ab.AllAccountAddresses)
	assert.Equal(ab.AllAccountAddresses, ab.AvailableAccountAddresses())

	deletedAt := time.Now()
	aa2 := NewAccountAddress(*id, &deletedAt)
	ab2 := NewAddressBook([]*AccountAddress{aa, aa2})
	assert.Equal(ab2.AllAccountAddresses, ab2.AllAccountAddresses)
	assert.Equal([]*AccountAddress{aa}, ab2.AvailableAccountAddresses())
}

func TestAddressBookSameValueAs(t *testing.T) {
	assert := require.New(t)

	id, _ := NewAccountAddressID("test_account_address_id")

	aa := NewAccountAddress(*id, nil)
	ab := NewAddressBook([]*AccountAddress{aa})

	deletedAt := time.Now()
	aa2 := NewAccountAddress(*id, &deletedAt)
	ab2 := NewAddressBook([]*AccountAddress{aa, aa2})
	assert.True(ab.SameValueAs(ab))
	assert.True(ab2.SameValueAs(ab2))
	assert.False(ab.SameValueAs(ab2))
}
