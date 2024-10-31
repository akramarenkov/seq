package seq

import (
	"github.com/akramarenkov/safe"
	"github.com/akramarenkov/safe/intspec"
	"golang.org/x/exp/constraints"
)

// Creates a slice of a linear sequence of integers from begin to end inclusive with the
// specified step.
//
// If begin is greater than end, the returned sequence will be decreasing, otherwise it
// will be increasing.
//
// If the step is not specified, a step of one will be used. If multiple steps are
// specified, the first one will be used.
//
// If a zero or negative step is specified, an error is returned.
func Int[Type constraints.Integer](begin, end Type, steps ...Type) ([]Type, error) {
	step := Type(1)

	if len(steps) != 0 {
		step = steps[0]
	}

	if step < 0 {
		return nil, ErrStepNegative
	}

	if step == 0 {
		return nil, ErrStepZero
	}

	sequence := make([]Type, intSize(begin, end, step))

	for id, number := range safe.Step(begin, end, step) {
		sequence[id] = number
	}

	// If the begin-end range is a multiple of the step, then the last element
	// of the sequence alredy has the end value after the loop, otherwise the value
	// of the last element equal to end must be set separately
	sequence[len(sequence)-1] = end

	return sequence, nil
}

func intSize[Type constraints.Integer](begin, end, step Type) uint64 {
	distance := safe.Dist(begin, end)

	size := distance / uint64(step)
	remainder := distance % uint64(step)

	if size == intspec.MaxUint64 {
		return size
	}

	// +1 due to the constant presence of begin in the sequence
	size++

	// if the begin-end range is not a multiple of the step, then the size must be
	// increased by one for store the end value
	if remainder != 0 {
		size++
	}

	return size
}

// Creates a slice of a linear sequence of integers from begin to end inclusive with the
// specified step.
//
// If begin is greater than end, the returned sequence will be decreasing, otherwise it
// will be increasing.
//
// If the step is not specified, a step of one will be used. If multiple steps are
// specified, the first one will be used.
//
// If a zero or negative step is specified, the function will panic.
func IntP[Type constraints.Integer](begin, end Type, steps ...Type) []Type {
	sequence, err := Int(begin, end, steps...)
	if err != nil {
		panic(err)
	}

	return sequence
}
