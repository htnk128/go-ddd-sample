package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAccountPasswordWithHash(t *testing.T) {
	assert := require.New(t)

	id, err := NewAccountID("AC_test_account_id")
	assert.NoError(err)
	for _, testData := range [][2]string{
		{strings.Repeat("a", 6), "2bfeb5406cda3f0d1b5bd7d349b1f4abdcce04c4b63c821eae3e166088109bb9"},
		{strings.Repeat("a", 100), "a0d50f9c776cdfda95109b6bed678fb4af071f2b241643441be76814299d67f0"},
	} {
		ae, err := NewAccountPasswordWithHash(testData[0], *id)
		assert.NoError(err)
		assert.Equal(testData[1], ae.Value())
	}

	for _, testData := range []string{
		"",
		strings.Repeat("a", 5),
		strings.Repeat("a", 101),
	} {
		_, err := NewAccountPasswordWithHash(testData, *id)
		assert.Error(err)
	}
}
