package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPhoneNumber(t *testing.T) {
	assert := require.New(t)

	for _, testData := range []string{
		"1",
		strings.Repeat("1", 50),
		"01234567890123456789012345678901234567890123456789",
	} {
		ae, err := NewPhoneNumber(testData)
		assert.NoError(err)
		assert.Equal(testData, ae.Value())
	}

	for _, testData := range []string{
		"",
		strings.Repeat("1", 51),
		"あ",
	} {
		_, err := NewPhoneNumber(testData)
		assert.Error(err)
	}
}
