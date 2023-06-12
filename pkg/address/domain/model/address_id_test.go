package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateAddressID(t *testing.T) {
	assert := require.New(t)

	ai := GenerateAddressID()
	assert.True(strings.HasPrefix(ai.ID(), "ADDR_"))
}

func TestNewAddressID(t *testing.T) {
	assert := require.New(t)

	for _, testData := range []string{
		"a",
		strings.Repeat("a", 64),
		"a_b-c-d-e",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	} {
		ae, err := NewAddressID(testData)
		assert.NoError(err)
		assert.Equal(testData, ae.ID())
	}

	for _, testData := range []string{
		"",
		strings.Repeat("a", 65),
		"あ",
	} {
		_, err := NewAddressID(testData)
		assert.Error(err)
	}
}
