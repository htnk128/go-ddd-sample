package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewAddress(t *testing.T) {
	assert := require.New(t)

	now := time.Now()
	id, _ := NewAddressID("ADDR_test_address_id")
	ownerID, _ := NewOwnerID("AC_test_account_id")
	fullName, _ := NewFullName("あいうえお")
	zipCode, _ := NewZipCode("1234567")
	stateOrRegion, _ := NewStateOrRegion("かきくけこ")
	line1, _ := NewLine1("さしすせそ")
	line2, _ := NewLine2("たちつてと")
	phoneNumber, _ := NewPhoneNumber("11111111111")

	a := NewAddress(*id, *ownerID, *fullName, *zipCode, *stateOrRegion, *line1, line2, *phoneNumber)
	assert.Equal(*id, a.ID)
	assert.Equal(*ownerID, a.OwnerID)
	assert.Equal(*fullName, a.FullName)
	assert.Equal(*zipCode, a.ZipCode)
	assert.Equal(*stateOrRegion, a.StateOrRegion)
	assert.Equal(*line1, a.Line1)
	assert.Equal(line2, a.Line2)
	assert.True(a.CreatedAt.After(now))
	assert.True(a.UpdatedAt.After(now))
	assert.Equal(1, len(a.OccurredEvents))
	assert.Equal(AddressCreated, a.OccurredEvents[0].Type)
	assert.Equal(*id, a.OccurredEvents[0].Address.ID)
	assert.Equal(*ownerID, a.OccurredEvents[0].Address.OwnerID)
	assert.Equal(*fullName, a.OccurredEvents[0].Address.FullName)
	assert.Equal(*zipCode, a.OccurredEvents[0].Address.ZipCode)
	assert.Equal(*stateOrRegion, a.OccurredEvents[0].Address.StateOrRegion)
	assert.Equal(*line1, a.OccurredEvents[0].Address.Line1)
	assert.Equal(line2, a.OccurredEvents[0].Address.Line2)
	assert.True(a.OccurredEvents[0].Address.CreatedAt.After(now))
	assert.True(a.OccurredEvents[0].Address.UpdatedAt.After(now))
}

func TestUpdateAddress(t *testing.T) {
	assert := require.New(t)

	now := time.Now()
	id, _ := NewAddressID("ADDR_test_address_id")
	ownerID, _ := NewOwnerID("AC_test_account_id")
	fullName, _ := NewFullName("あいうえお")
	zipCode, _ := NewZipCode("1234567")
	stateOrRegion, _ := NewStateOrRegion("かきくけこ")
	line1, _ := NewLine1("さしすせそ")
	line2, _ := NewLine2("たちつてと")
	phoneNumber, _ := NewPhoneNumber("11111111111")

	a := NewAddress(*id, *ownerID, *fullName, *zipCode, *stateOrRegion, *line1, line2, *phoneNumber)
	fullName2, _ := NewFullName("あいうえおかきくけこ")
	zipCode2, _ := NewZipCode("1234567")
	stateOrRegion2, _ := NewStateOrRegion("かきくけこさしすせそ")
	line12, _ := NewLine1("さしすせそたちつてと")
	line22, _ := NewLine2("たちつてとなにぬねの")
	phoneNumber2, _ := NewPhoneNumber("22222222222")

	err := a.Update(fullName2, zipCode2, stateOrRegion2, line12, line22, phoneNumber2)
	assert.NoError(err)
	assert.Equal(*fullName2, a.FullName)
	assert.Equal(*zipCode2, a.ZipCode)
	assert.Equal(*stateOrRegion2, a.StateOrRegion)
	assert.Equal(*line12, a.Line1)
	assert.Equal(line22, a.Line2)
	assert.True(a.UpdatedAt.After(now))
	assert.Equal(2, len(a.OccurredEvents))
	assert.Equal(AddressUpdated, a.OccurredEvents[1].Type)
	assert.Equal(*fullName2, a.OccurredEvents[1].Address.FullName)
	assert.Equal(*zipCode2, a.OccurredEvents[1].Address.ZipCode)
	assert.Equal(*stateOrRegion2, a.OccurredEvents[1].Address.StateOrRegion)
	assert.Equal(*line12, a.OccurredEvents[1].Address.Line1)
	assert.Equal(line22, a.OccurredEvents[1].Address.Line2)
	assert.True(a.OccurredEvents[1].Address.UpdatedAt.After(now))

	err = a.Update(nil, nil, nil, nil, nil, nil)
	assert.NoError(err)
	assert.Equal(*fullName2, a.FullName)
	assert.Equal(*zipCode2, a.ZipCode)
	assert.Equal(*stateOrRegion2, a.StateOrRegion)
	assert.Equal(*line12, a.Line1)
	assert.Equal(line22, a.Line2)
	assert.True(a.UpdatedAt.After(now))
	assert.Equal(3, len(a.OccurredEvents))
	assert.Equal(AddressUpdated, a.OccurredEvents[2].Type)
	assert.Equal(*fullName2, a.OccurredEvents[2].Address.FullName)
	assert.Equal(*zipCode2, a.OccurredEvents[2].Address.ZipCode)
	assert.Equal(*stateOrRegion2, a.OccurredEvents[2].Address.StateOrRegion)
	assert.Equal(*line12, a.OccurredEvents[2].Address.Line1)
	assert.Equal(line22, a.OccurredEvents[2].Address.Line2)
	assert.True(a.OccurredEvents[2].Address.UpdatedAt.After(now))

	a.Delete()
	err = a.Update(nil, nil, nil, nil, nil, nil)
	assert.Error(err)
}

func TestDeleteAddress(t *testing.T) {
	assert := require.New(t)

	now := time.Now()
	id, _ := NewAddressID("ADDR_test_address_id")
	ownerID, _ := NewOwnerID("AC_test_account_id")
	fullName, _ := NewFullName("あいうえお")
	zipCode, _ := NewZipCode("1234567")
	stateOrRegion, _ := NewStateOrRegion("かきくけこ")
	line1, _ := NewLine1("さしすせそ")
	line2, _ := NewLine2("たちつてと")
	phoneNumber, _ := NewPhoneNumber("11111111111")

	a := NewAddress(*id, *ownerID, *fullName, *zipCode, *stateOrRegion, *line1, line2, *phoneNumber)
	a.Delete()
	deletedAt := a.DeletedAt
	updatedAt := a.UpdatedAt
	assert.True(a.DeletedAt.After(now))
	assert.True(a.UpdatedAt.After(now))
	assert.Equal(2, len(a.OccurredEvents))
	assert.Equal(AddressDeleted, a.OccurredEvents[1].Type)
	assert.True(a.OccurredEvents[1].Address.DeletedAt.After(now))
	assert.True(a.OccurredEvents[1].Address.UpdatedAt.After(now))

	a.Delete()
	assert.Equal(deletedAt, a.DeletedAt)
	assert.Equal(updatedAt, a.UpdatedAt)
	assert.Equal(2, len(a.OccurredEvents))
}

func TestSameIdentityAs(t *testing.T) {
	assert := require.New(t)

	id, _ := NewAddressID("ADDR_test_address_id")
	ownerID, _ := NewOwnerID("AC_test_account_id")
	ownerID2, _ := NewOwnerID("AC_test_account_id2")
	fullName, _ := NewFullName("あいうえお")
	fullName2, _ := NewFullName("あいうえお2")
	zipCode, _ := NewZipCode("1234567")
	zipCode2, _ := NewZipCode("12345672")
	stateOrRegion, _ := NewStateOrRegion("かきくけこ")
	stateOrRegion2, _ := NewStateOrRegion("かきくけこ2")
	line1, _ := NewLine1("さしすせそ")
	line12, _ := NewLine1("さしすせそ2")
	line2, _ := NewLine2("たちつてと")
	line22, _ := NewLine2("たちつてと2")
	phoneNumber, _ := NewPhoneNumber("11111111111")
	phoneNumber2, _ := NewPhoneNumber("111111111112")

	a1 := NewAddress(*id, *ownerID, *fullName, *zipCode, *stateOrRegion, *line1, line2, *phoneNumber)
	a2 := NewAddress(*id, *ownerID2, *fullName2, *zipCode2, *stateOrRegion2, *line12, line22, *phoneNumber2)
	assert.True(a1.SameIdentityAs(a2))

	id2, _ := NewAddressID("ADDR_test_address_id2")
	a3 := NewAddress(*id2, *ownerID, *fullName, *zipCode, *stateOrRegion, *line1, line2, *phoneNumber)
	assert.False(a1.SameIdentityAs(a3))
}
