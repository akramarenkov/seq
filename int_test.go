package seq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {
	require.Equal(t,
		[]int8{
			-128, -112, -96, -80, -64, -48, -32, -16, 0, 16, 32, 48, 64, 80, 96, 112,
		},
		Int[int8](-128, 127, 16),
	)
	require.Equal(t,
		[]int8{
			-128, -111, -94, -77, -60, -43, -26, -9, 8, 25, 42, 59, 76, 93, 110, 127,
		},
		Int[int8](-128, 127, 17),
	)

	require.Equal(t,
		[]uint8{
			0, 16, 32, 48, 64, 80, 96, 112, 128, 144, 160, 176, 192, 208, 224, 240,
		},
		Int[uint8](0, 255, 16),
	)
	require.Equal(t,
		[]uint8{
			0, 17, 34, 51, 68, 85, 102, 119, 136, 153, 170, 187, 204, 221, 238, 255,
		},
		Int[uint8](0, 255, 17),
	)
}
