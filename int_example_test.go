package seq_test

import (
	"fmt"

	"github.com/akramarenkov/seq"
)

func ExampleInt() {
	fmt.Println(seq.Int(1, 8, 3))
	// Output:
	// [1 4 7]
}
