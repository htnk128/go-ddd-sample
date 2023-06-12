package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAccountAddressID(t *testing.T) {
	assert := require.New(t)

	for _, testData := range []string{
		"a",
		strings.Repeat("a", 64),
		"a_b-c-d-e",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	} {
		ae, err := NewAccountAddressID(testData)
		assert.NoError(err)
		assert.Equal(testData, ae.ID())
	}

	for _, testData := range []string{
		"",
		strings.Repeat("a", 65),
		"あ",
	} {
		_, err := NewAccountAddressID(testData)
		assert.Error(err)
	}
}
