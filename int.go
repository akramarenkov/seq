package seq

import (
	"github.com/akramarenkov/safe"
	"golang.org/x/exp/constraints"
)

// Creates a slice of a sequence of integers from begin to end with the specified step.
//
// If begin is greater than end, the returned sequence will be decreasing, otherwise it
// will be increasing.
//
// If the step is not specified, a step of one will be used. If multiple steps are
// specified, the first one will be used.
//
// If the step is not a multiple of the begin-end range, the end value will not be
// returned.
//
// If a zero or negative step is specified, the function will panic.
func Int[Type constraints.Integer](begin Type, end Type, steps ...Type) []Type {
	step := Type(1)

	if len(steps) != 0 {
		step = steps[0]
	}

	sequence := make([]Type, safe.StepSize(begin, end, step))

	for id, number := range safe.Step(begin, end, step) {
		sequence[id] = number
	}

	return sequence
}
