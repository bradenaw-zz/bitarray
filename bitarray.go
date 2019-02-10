// Package bitarray provides a structure for representing arrays of k-bit elements.
package bitarray

import (
	"fmt"
	"strconv"
	"strings"
)

type Array struct {
	inner []uint64
	n     int
	k     uint
}

// Returns a new bit array with n elements of k bits each. k <= 64.
func New(n int, k uint) Array {
	if k > 64 {
		panic("can't support elements over 64 bits")
	}
	return Array{
		inner: make([]uint64, (uint(n)*k+63)/64),
		n:     n,
		k:     k,
	}
}

// Returns the number of elements in the array.
func (a *Array) Len() int {
	return a.n
}

// Returns the length, in bits, of each element.
func (a *Array) K() uint {
	return a.k
}

func (a *Array) mask() uint64 {
	return (1 << a.k) - 1
}

// Get the item at i as the bottom k bits of the result. Panics if i is out of bounds.
func (a *Array) Get(i int) uint64 {
	if i < 0 {
		panic("out of bounds")
	}
	bitStart := uint(i) * a.k
	bitsInFirst := 64 - (bitStart % 64)
	if bitsInFirst >= a.k {
		return (a.inner[bitStart/64] >> (64 - (bitStart % 64) - a.k)) & a.mask()
	} else {
		bitsInSecond := a.k - bitsInFirst

		return (a.inner[bitStart/64]<<bitsInSecond | a.inner[bitStart/64+1]>>(64-bitsInSecond)) & a.mask()
	}
}

// Sets the item at i to be the bottom k bits of v. Panics if v is more than k bits, or if i is out
// of bounds.
func (a *Array) Set(i int, v uint64) {
	if v != v&a.mask() {
		panic("v has more than k bits")
	}
	if i < 0 {
		panic("out of bounds")
	}

	bitStart := uint(i) * a.k
	bitsInFirst := 64 - (bitStart % 64)
	if bitsInFirst >= a.k {
		shift := (64 - (bitStart % 64) - a.k)
		a.inner[bitStart/64] = a.inner[bitStart/64] & ^(a.mask()<<shift) | v<<shift
	} else {
		bitsInSecond := a.k - bitsInFirst

		firstMask := ((uint64(1) << bitsInFirst) - 1)
		a.inner[bitStart/64] = (a.inner[bitStart/64] & ^firstMask) | ((v >> bitsInSecond) & firstMask)
		secondMask := ^((uint64(1) << (64 - bitsInSecond)) - 1)
		a.inner[bitStart/64+1] = (a.inner[bitStart/64+1] & ^secondMask) | ((v << (64 - bitsInSecond)) & secondMask)
	}
}

func (a *Array) dump() {
	var b strings.Builder
	for _, item := range a.inner {
		fmt.Printf("%064b\n", item)
		_, _ = b.WriteString(fmt.Sprintf("%064b", item))
	}
	s := b.String()
	for i := 0; i < a.n; i++ {
		bitsS := s[i*int(a.k) : i*int(a.k)+int(a.k)]
		parsed, err := strconv.ParseUint(bitsS, 2, 64)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %016x\n", bitsS, parsed)
	}
}
