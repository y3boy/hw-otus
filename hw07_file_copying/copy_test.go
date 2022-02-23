package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("Invalid limit", func(t *testing.T) {
		err := Copy("", "", 0, -10)
		require.Equal(t, err, ErrOffsetExceedsFileSize, "limit is invalid")
	})

	t.Run("Invalid offser", func(t *testing.T) {
		err := Copy("", "", -10, 0)
		require.Equal(t, err, ErrOffsetExceedsFileSize, "offser is invalid")
	})
}
