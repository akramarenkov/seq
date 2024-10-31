package seq

import "errors"

var (
	ErrStepNegative = errors.New("step is negative")
	ErrStepZero     = errors.New("step is zero")
)
