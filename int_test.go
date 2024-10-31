package seq

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {
	testIntSig(t)
	testIntUns(t)
}

func testIntSig(t *testing.T) {
	sequence, err := Int[int8](-2, 3)
	require.NoError(t, err)
	require.Equal(t, []int8{-2, -1, 0, 1, 2, 3}, sequence)

	sequence, err = Int[int8](-2, 3, 1)
	require.NoError(t, err)
	require.Equal(t, []int8{-2, -1, 0, 1, 2, 3}, sequence)

	sequence, err = Int[int8](-2, 3, 1, 2)
	require.NoError(t, err)
	require.Equal(t, []int8{-2, -1, 0, 1, 2, 3}, sequence)

	sequence, err = Int[int8](-2, 3, 2, 1)
	require.NoError(t, err)
	require.Equal(t, []int8{-2, 0, 2, 3}, sequence)

	sequence, err = Int[int8](-2, 2, 2)
	require.NoError(t, err)
	require.Equal(t, []int8{-2, 0, 2}, sequence)

	sequence, err = Int[int8](-128, 127, 16)
	require.NoError(t, err)
	require.Equal(
		t,
		[]int8{
			-128, -112, -96, -80, -64, -48, -32, -16, 0, 16, 32, 48, 64, 80, 96, 112, 127,
		},
		sequence,
	)

	sequence, err = Int[int8](-128, 127, 17)
	require.NoError(t, err)
	require.Equal(
		t,
		[]int8{
			-128, -111, -94, -77, -60, -43, -26, -9, 8, 25, 42, 59, 76, 93, 110, 127,
		},
		sequence,
	)
}

func testIntUns(t *testing.T) {
	sequence, err := Int[uint8](1, 4, 2)
	require.NoError(t, err)
	require.Equal(t, []uint8{1, 3, 4}, sequence)

	sequence, err = Int[uint8](1, 5, 2)
	require.NoError(t, err)
	require.Equal(t, []uint8{1, 3, 5}, sequence)

	sequence, err = Int[uint8](0, 255, 16)
	require.NoError(t, err)
	require.Equal(
		t,
		[]uint8{
			0, 16, 32, 48, 64, 80, 96, 112, 128, 144, 160, 176, 192, 208, 224, 240, 255,
		},
		sequence,
	)

	sequence, err = Int[uint8](0, 255, 17)
	require.NoError(t, err)
	require.Equal(
		t,
		[]uint8{
			0, 17, 34, 51, 68, 85, 102, 119, 136, 153, 170, 187, 204, 221, 238, 255,
		},
		sequence,
	)
}

func TestIntError(t *testing.T) {
	sequence, err := Int[int8](-2, 3, 0)
	require.Error(t, err)
	require.Equal(t, []int8(nil), sequence)

	sequence, err = Int[int8](-2, 3, -1)
	require.Error(t, err)
	require.Equal(t, []int8(nil), sequence)
}

func TestIntSize(t *testing.T) {
	require.Equal(t, uint64(6), intSize(-2, 3, 1))
	require.Equal(t, uint64(3), intSize(-2, 2, 2))
	require.Equal(t, uint64(4), intSize(-2, 3, 2))
	require.Equal(t, uint64(3), intSize(-2, 4, 3))
	require.Equal(t, uint64(4), intSize(-2, 5, 3))
	require.Equal(t, uint64(math.MaxUint64), intSize(math.MinInt64, math.MaxInt64, 1))
	require.Equal(t, uint64(math.MaxUint64), intSize(math.MinInt64+1, math.MaxInt64, 1))
	require.Equal(t, uint64(math.MaxUint64), intSize(math.MinInt64, math.MaxInt64-1, 1))
	require.Equal(t, uint64(math.MaxUint64-1), intSize(math.MinInt64+2, math.MaxInt64, 1))
	require.Equal(t, uint64(math.MaxUint64-1), intSize(math.MinInt64, math.MaxInt64-2, 1))
	require.Equal(t, uint64(math.MaxUint64/2+2), intSize(math.MinInt64, math.MaxInt64, 2))
	require.Equal(t, uint64(math.MaxUint64/2+1), intSize(math.MinInt64+1, math.MaxInt64, 2))
	require.Equal(t, uint64(math.MaxUint64/2+1), intSize(math.MinInt64, math.MaxInt64-1, 2))
}

func TestIntSure(t *testing.T) {
	require.Equal(t, []int8{-2, -1, 0, 1, 2, 3}, IntSure[int8](-2, 3))
}

func TestIntSurePanic(t *testing.T) {
	require.Panics(t, func() { IntSure[int8](-2, 3, 0) })
	require.Panics(t, func() { IntSure[int8](-2, 3, -1) })
}

func BenchmarkIntReference(b *testing.B) {
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

func BenchmarkInt(b *testing.B) {
	expected := []int{0, 1, 2, 3, 4, 5}

	var (
		sequence []int
		err      error
	)

	for range b.N {
		sequence, err = Int(0, len(expected)-1)
	}

	require.NoError(b, err)
	require.Equal(b, expected, sequence)
}

func BenchmarkIntStep(b *testing.B) {
	expected := []int{0, 1, 2, 3, 4, 5}

	var (
		sequence []int
		err      error
	)

	for range b.N {
		sequence, err = Int(0, len(expected)-1, 1)
	}

	require.NoError(b, err)
	require.Equal(b, expected, sequence)
}

func BenchmarkIntSure(b *testing.B) {
	expected := []int{0, 1, 2, 3, 4, 5}

	var sequence []int

	for range b.N {
		sequence = IntSure(0, len(expected)-1, 1)
	}

	require.Equal(b, expected, sequence)
}
