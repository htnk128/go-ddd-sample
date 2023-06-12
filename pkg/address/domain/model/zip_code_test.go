package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewZipCode(t *testing.T) {
	assert := require.New(t)

	for _, testData := range []string{
		"a",
		strings.Repeat("a", 50),
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMN0123456789",
	} {
		ae, err := NewZipCode(testData)
		assert.NoError(err)
		assert.Equal(testData, ae.Value())
	}

	for _, testData := range []string{
		"",
		strings.Repeat("a", 51),
		"あ",
	} {
		_, err := NewZipCode(testData)
		assert.Error(err)
	}
}
