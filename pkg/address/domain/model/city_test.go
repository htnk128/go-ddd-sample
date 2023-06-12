package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCity(t *testing.T) {
	assert := require.New(t)

	for _, testData := range []string{
		"a",
		strings.Repeat("a", 100),
	} {
		ae, err := NewCity(testData)
		assert.NoError(err)
		assert.Equal(testData, ae.Value())
	}

	for _, testData := range []string{
		"",
		strings.Repeat("a", 101),
	} {
		_, err := NewCity(testData)
		assert.Error(err)
	}
}
