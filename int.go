package seq

import (
	"github.com/akramarenkov/safe"
	"golang.org/x/exp/constraints"
)

// Generates a slice of a sequence of integers from begin to end with the specified step.
//
// If the step is not a multiple of the begin-end range, the end value will not be
// returned.
//
// If begin is greater than end, the returned sequence will be decreasing, otherwise it
// will be increasing.
//
// If a zero or negative step is specified, the function will panic.
func Int[Type constraints.Integer](begin Type, end Type, step Type) []Type {
	sequence := make([]Type, safe.IterStepSize(begin, end, step))

	for id, number := range safe.IterStep2(begin, end, step) {
		sequence[id] = number
	}

	return sequence
}
