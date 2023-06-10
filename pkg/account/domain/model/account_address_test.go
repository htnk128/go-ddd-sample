package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewAccountAddress(t *testing.T) {
	assert := require.New(t)

	id, _ := NewAccountAddressID("test_account_address_id")

	aa := NewAccountAddress(*id, nil)
	assert.Equal(*id, aa.ID)
	assert.True(aa.DeletedAt == nil)
	assert.True(aa.IsAvailable())

	deletedAt := time.Now()
	aa2 := NewAccountAddress(*id, &deletedAt)
	assert.Equal(&deletedAt, aa2.DeletedAt)
	assert.False(aa2.IsAvailable())
}

func TestSameValueAs(t *testing.T) {
	assert := require.New(t)

	id, _ := NewAccountAddressID("test_account_address_id")
	id2, _ := NewAccountAddressID("test_account_address_id2")
	deletedAt := time.Now()
	deletedAt2 := deletedAt.Add(1 * time.Minute)

	aa := NewAccountAddress(*id, &deletedAt)
	aa2 := NewAccountAddress(*id, &deletedAt)
	assert.True(aa.SameValueAs(aa2))

	aa3 := NewAccountAddress(*id, nil)
	aa4 := NewAccountAddress(*id, nil)
	assert.True(aa3.SameValueAs(aa4))

	aa5 := NewAccountAddress(*id, &deletedAt)
	aa6 := NewAccountAddress(*id2, &deletedAt)
	assert.False(aa5.SameValueAs(aa6))

	aa7 := NewAccountAddress(*id, &deletedAt)
	aa8 := NewAccountAddress(*id, &deletedAt2)
	assert.False(aa7.SameValueAs(aa8))
}
