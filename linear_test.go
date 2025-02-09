package seq

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinear(t *testing.T) {
	testLinearSig(t)
	testLinearUns(t)
}

func testLinearSig(t *testing.T) {
	require.Equal(t, []int8{-2, -1, 0, 1, 2, 3}, Linear[int8](-2, 3))
	require.Equal(t, []int8{-2, -1, 0, 1, 2, 3}, Linear[int8](-2, 3, 1))
	require.Equal(t, []int8{-2, -1, 0, 1, 2, 3}, Linear[int8](-2, 3, 1, 2))
	require.Equal(t, []int8{-2, 0, 2, 3}, Linear[int8](-2, 3, 2, 1))
	require.Equal(t, []int8{-2, 0, 2}, Linear[int8](-2, 2, 2))

	require.Equal(
		t,
		[]int8{
			-128, -112, -96, -80, -64, -48, -32, -16, 0, 16, 32, 48, 64, 80, 96, 112, 127,
		},
		Linear[int8](-128, 127, 16),
	)

	require.Equal(
		t,
		[]int8{
			-128, -111, -94, -77, -60, -43, -26, -9, 8, 25, 42, 59, 76, 93, 110, 127,
		},
		Linear[int8](-128, 127, 17),
	)
}

func testLinearUns(t *testing.T) {
	require.Equal(t, []uint8{1, 3, 4}, Linear[uint8](1, 4, 2))
	require.Equal(t, []uint8{1, 3, 5}, Linear[uint8](1, 5, 2))

	require.Equal(
		t,
		[]uint8{
			0, 16, 32, 48, 64, 80, 96, 112, 128, 144, 160, 176, 192, 208, 224, 240, 255,
		},
		Linear[uint8](0, 255, 16),
	)

	require.Equal(
		t,
		[]uint8{
			0, 17, 34, 51, 68, 85, 102, 119, 136, 153, 170, 187, 204, 221, 238, 255,
		},
		Linear[uint8](0, 255, 17),
	)
}

func TestLinearPanic(t *testing.T) {
	require.Panics(t, func() { Linear[int8](-2, 3, 0) })
	require.Panics(t, func() { Linear[int8](-2, 3, -1) })
}

func TestLinearSize(t *testing.T) {
	require.Equal(t, uint64(6), linearSize(-2, 3, 1))
	require.Equal(t, uint64(3), linearSize(-2, 2, 2))
	require.Equal(t, uint64(4), linearSize(-2, 3, 2))
	require.Equal(t, uint64(3), linearSize(-2, 4, 3))
	require.Equal(t, uint64(4), linearSize(-2, 5, 3))
	require.Equal(t, uint64(math.MaxUint64), linearSize(math.MinInt64, math.MaxInt64, 1))
	require.Equal(t, uint64(math.MaxUint64), linearSize(math.MinInt64+1, math.MaxInt64, 1))
	require.Equal(t, uint64(math.MaxUint64), linearSize(math.MinInt64, math.MaxInt64-1, 1))
	require.Equal(t, uint64(math.MaxUint64-1), linearSize(math.MinInt64+2, math.MaxInt64, 1))
	require.Equal(t, uint64(math.MaxUint64-1), linearSize(math.MinInt64, math.MaxInt64-2, 1))
	require.Equal(t, uint64(math.MaxUint64/2+2), linearSize(math.MinInt64, math.MaxInt64, 2))
	require.Equal(t, uint64(math.MaxUint64/2+1), linearSize(math.MinInt64+1, math.MaxInt64, 2))
	require.Equal(t, uint64(math.MaxUint64/2+1), linearSize(math.MinInt64, math.MaxInt64-1, 2))
}

func BenchmarkLinearReference(b *testing.B) {
	expected := []int{0, 1, 2, 3, 4, 5}

	var sequence []int

	for range b.N {
		sequence = make([]int, len(expected))

		for id := range expected {
			sequence[id] = id
		}
	}

	require.Equal(b, expected, sequence)
}

func BenchmarkLinear(b *testing.B) {
	expected := []int{0, 1, 2, 3, 4, 5}

	var sequence []int

	for range b.N {
		sequence = Linear(0, len(expected)-1, 1)
	}

	require.Equal(b, expected, sequence)
}
