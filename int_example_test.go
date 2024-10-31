package seq_test

import (
	"fmt"

	"github.com/akramarenkov/seq"
)

func ExampleInt() {
	fmt.Println(seq.IntP(1, 8, 3))
	// Output:
	// [1 4 7 8]
}
