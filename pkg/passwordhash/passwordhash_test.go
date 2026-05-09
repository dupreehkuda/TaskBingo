package passwordhash_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dupreehkuda/TaskBingo/pkg/passwordhash"
)

func TestHash_KnownVector(t *testing.T) {
	// Locked-in vector ensures byte-for-byte compat with the prior service.
	got := passwordhash.Hash("hunter2", "ABCDEFGHIJ")
	require.Equal(t, "823e6d16f3b75390321d695d9c59c028", got)
}

func TestSalt_Length(t *testing.T) {
	s, err := passwordhash.Salt(10)
	require.NoError(t, err)
	require.Len(t, s, 10)
}

func TestSalt_AlphabetOnly(t *testing.T) {
	const allowed = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	s, err := passwordhash.Salt(64)
	require.NoError(t, err)
	for _, c := range s {
		require.Contains(t, allowed, string(c))
	}
}
