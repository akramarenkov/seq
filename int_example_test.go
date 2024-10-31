package seq_test

import (
	"fmt"

	"github.com/akramarenkov/seq"
)

func ExampleInt() {
	sequence, err := seq.Int(1, 8, 3)
	fmt.Println(err)
	fmt.Println(sequence)
	// Output:
	// <nil>
	// [1 4 7 8]
}

func ExampleIntSure() {
	fmt.Println(seq.IntSure(1, 8, 3))
	// Output:
	// [1 4 7 8]
}
