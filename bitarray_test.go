package bitarray

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/bradenaw/trand"
)

func TestBitArray(t *testing.T) {
	a := New(3, 6)
	require.Equal(t, uint64(0), a.Get(0))
	require.Equal(t, uint64(0), a.Get(1))
	require.Equal(t, uint64(0), a.Get(2))

	a.Set(0, 0x2b)
	a.Set(1, 0x1c)
	a.Set(2, 0x3a)
	a.Set(1, 0x1e)
	require.Equal(t, uint64(0x2b), a.Get(0))
	require.Equal(t, uint64(0x1e), a.Get(1))
	require.Equal(t, uint64(0x3a), a.Get(2))
}

func TestBitArrayRandom(t *testing.T) {
	trand.RandomN(t, 500, func(t *testing.T, r *rand.Rand) {
		n := r.Int()%100 + 1
		k := uint(r.Int()%63 + 1)

		a := New(n, k)
		expected := make([]uint64, n)
		mask := uint64((1 << k) - 1)
		for i := 0; i < n*5; i++ {
			idx := r.Int() % n
			val := r.Uint64() & mask
			expected[idx] = val
			a.Set(idx, val)
		}

		for i := range expected {
			require.Equal(t, expected[i], a.Get(i))
		}
	})
}
