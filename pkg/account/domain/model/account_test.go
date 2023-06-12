package model

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewAccount(t *testing.T) {
	assert := require.New(t)

	now := time.Now()
	id, _ := NewAccountID("AC_test_account_id")
	name, _ := NewAccountName("あいうえお")
	namePronunciation, _ := NewAccountNamePronunciation("アイウエオ")
	email, _ := NewAccountEmail("example@example.com")
	password, _ := NewAccountPasswordWithHash(strings.Repeat("a", 100), *id)

	a := NewAccount(*id, *name, *namePronunciation, *email, *password)
	assert.Equal(*id, a.ID)
	assert.Equal(*name, a.Name)
	assert.Equal(*namePronunciation, a.NamePronunciation)
	assert.Equal(*email, a.Email)
	assert.Equal(*password, a.Password)
	assert.True(a.CreatedAt.After(now))
	assert.True(a.UpdatedAt.After(now))
	assert.Equal(1, len(a.OccurredEvents))
	assert.Equal(AccountCreated, a.OccurredEvents[0].Type)
	assert.Equal(*id, a.OccurredEvents[0].Account.ID)
	assert.Equal(*name, a.OccurredEvents[0].Account.Name)
	assert.Equal(*namePronunciation, a.OccurredEvents[0].Account.NamePronunciation)
	assert.Equal(*email, a.OccurredEvents[0].Account.Email)
	assert.Equal(*password, a.OccurredEvents[0].Account.Password)
	assert.True(a.OccurredEvents[0].Account.CreatedAt.After(now))
	assert.True(a.OccurredEvents[0].Account.UpdatedAt.After(now))
}

func TestUpdateAccount(t *testing.T) {
	assert := require.New(t)

	now := time.Now()
	id, _ := NewAccountID("AC_test_account_id")
	name, _ := NewAccountName("あいうえお")
	namePronunciation, _ := NewAccountNamePronunciation("アイウエオ")
	email, _ := NewAccountEmail("example@example.com")
	password, _ := NewAccountPasswordWithHash(strings.Repeat("a", 100), *id)

	a := NewAccount(*id, *name, *namePronunciation, *email, *password)
	name2, _ := NewAccountName("あいうえおかきくけこ")
	namePronunciation2, _ := NewAccountNamePronunciation("アイウエオカキクケコ")
	email2, _ := NewAccountEmail("example2@example.com")
	password2, _ := NewAccountPasswordWithHash(strings.Repeat("b", 100), *id)

	err := a.Update(name2, namePronunciation2, email2, password2)
	assert.NoError(err)
	assert.Equal(*name2, a.Name)
	assert.Equal(*namePronunciation2, a.NamePronunciation)
	assert.Equal(*email2, a.Email)
	assert.Equal(*password2, a.Password)
	assert.True(a.UpdatedAt.After(now))
	assert.Equal(2, len(a.OccurredEvents))
	assert.Equal(AccountUpdated, a.OccurredEvents[1].Type)
	assert.Equal(*name2, a.OccurredEvents[1].Account.Name)
	assert.Equal(*namePronunciation2, a.OccurredEvents[1].Account.NamePronunciation)
	assert.Equal(*email2, a.OccurredEvents[1].Account.Email)
	assert.Equal(*password2, a.OccurredEvents[1].Account.Password)
	assert.True(a.OccurredEvents[1].Account.UpdatedAt.After(now))

	err = a.Update(nil, nil, nil, nil)
	assert.NoError(err)
	assert.Equal(*name2, a.Name)
	assert.Equal(*namePronunciation2, a.NamePronunciation)
	assert.Equal(*email2, a.Email)
	assert.Equal(*password2, a.Password)
	assert.True(a.UpdatedAt.After(now))
	assert.Equal(3, len(a.OccurredEvents))
	assert.Equal(AccountUpdated, a.OccurredEvents[2].Type)
	assert.Equal(*name2, a.OccurredEvents[2].Account.Name)
	assert.Equal(*namePronunciation2, a.OccurredEvents[2].Account.NamePronunciation)
	assert.Equal(*email2, a.OccurredEvents[2].Account.Email)
	assert.Equal(*password2, a.OccurredEvents[2].Account.Password)
	assert.True(a.OccurredEvents[2].Account.UpdatedAt.After(now))

	a.Delete()
	err = a.Update(nil, nil, nil, nil)
	assert.Error(err)
}

func TestDeleteAccount(t *testing.T) {
	assert := require.New(t)

	now := time.Now()
	id, _ := NewAccountID("AC_test_account_id")
	name, _ := NewAccountName("あいうえお")
	namePronunciation, _ := NewAccountNamePronunciation("アイウエオ")
	email, _ := NewAccountEmail("example@example.com")
	password, _ := NewAccountPasswordWithHash(strings.Repeat("a", 100), *id)

	a := NewAccount(*id, *name, *namePronunciation, *email, *password)
	a.Delete()
	deletedAt := a.DeletedAt
	updatedAt := a.UpdatedAt
	assert.True(a.DeletedAt.After(now))
	assert.True(a.UpdatedAt.After(now))
	assert.Equal(2, len(a.OccurredEvents))
	assert.Equal(AccountDeleted, a.OccurredEvents[1].Type)
	assert.True(a.OccurredEvents[1].Account.DeletedAt.After(now))
	assert.True(a.OccurredEvents[1].Account.UpdatedAt.After(now))

	a.Delete()
	assert.Equal(deletedAt, a.DeletedAt)
	assert.Equal(updatedAt, a.UpdatedAt)
	assert.Equal(2, len(a.OccurredEvents))
}

func TestSameIdentityAs(t *testing.T) {
	assert := require.New(t)

	id, _ := NewAccountID("AC_test_account_id")
	name, _ := NewAccountName("あいうえお")
	name2, _ := NewAccountName("あいうえお1")
	namePronunciation, _ := NewAccountNamePronunciation("アイウエオ")
	namePronunciation2, _ := NewAccountNamePronunciation("アイウエオ2")
	email, _ := NewAccountEmail("example@example.com")
	email2, _ := NewAccountEmail("example@example.com3")
	password, _ := NewAccountPasswordWithHash(strings.Repeat("a", 100), *id)
	password2, _ := NewAccountPasswordWithHash(strings.Repeat("b", 100), *id)

	a1 := NewAccount(*id, *name, *namePronunciation, *email, *password)
	a2 := NewAccount(*id, *name2, *namePronunciation2, *email2, *password2)
	assert.True(a1.SameIdentityAs(a2))

	id2, _ := NewAccountID("AC_test_account_id2")
	a3 := NewAccount(*id2, *name, *namePronunciation, *email, *password)
	assert.False(a1.SameIdentityAs(a3))
}
