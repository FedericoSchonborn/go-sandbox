package dyntuple_test

import (
	"fmt"

	"github.com/FedericoSchonborn/go-sandbox/tuple/dyntuple"
)

func Example() {
	dt := dyntuple.New("Hello", 42, true)
	fmt.Println(dyntuple.Get[string](dt, 0))
	fmt.Println(dyntuple.Get[int](dt, 1))
	fmt.Println(dyntuple.Get[float32](dt, 2))

	// Output:
	// Hello <nil>
	// 42 <nil>
	// 0 failed to convert value of type bool to type float32
}
