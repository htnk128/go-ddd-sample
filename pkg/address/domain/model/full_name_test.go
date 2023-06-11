package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewFullName(t *testing.T) {
	assert := require.New(t)

	for _, testData := range []string{
		"a",
		strings.Repeat("a", 100),
	} {
		ae, err := NewFullName(testData)
		assert.NoError(err)
		assert.Equal(testData, ae.Value())
	}

	for _, testData := range []string{
		"",
		strings.Repeat("a", 101),
	} {
		_, err := NewFullName(testData)
		assert.Error(err)
	}
}
