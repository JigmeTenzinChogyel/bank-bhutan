package util

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	min := int64(0)
	max := int64(1000)
	results := make(map[int64]bool)
	for range 100 {
		r := RandomInt(min, max)
		require.GreaterOrEqual(t, r, min)
		require.LessOrEqual(t, r, max)
		results[r] = true
	}

	// Check that multiple different values were produced
	require.Greater(t, len(results), 1, "RandomInt seems to return same value repeatedly")
}

func TestRandomString(t *testing.T) {
	n := 6
	result := RandomString(n)

	require.Len(t, result, n)
	require.NotEmpty(t, result)

	for _, char := range result {
		require.True(t, unicode.IsLetter(rune(char)), "expected only letters, got: %q", char)
	}
}

func TestRandomOwner(t *testing.T) {
	result := RandomOwner()
	require.NotEmpty(t, result)
	for _, char := range result {
		require.True(t, unicode.IsLetter(rune(char)), "expected only letters, got: %q", char)
	}
}
