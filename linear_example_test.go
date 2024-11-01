package seq_test

import (
	"fmt"

	"github.com/akramarenkov/seq"
)

func ExampleLinear() {
	sequence, err := seq.Linear(1, 8, 3)
	fmt.Println(err)
	fmt.Println(sequence)
	// Output:
	// <nil>
	// [1 4 7 8]
}

func ExampleLinearSure() {
	fmt.Println(seq.LinearSure(1, 8, 3))
	// Output:
	// [1 4 7 8]
}
