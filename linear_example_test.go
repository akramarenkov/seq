package seq_test

import (
	"fmt"

	"github.com/akramarenkov/seq"
)

func ExampleLinear() {
	fmt.Println(seq.Linear(1, 8, 3))
	// Output:
	// [1 4 7 8]
}
