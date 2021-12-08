package spec_test

import (
	"fmt"

	"github.com/fdschonborn/go-sandbox/spec"
)

func Example() {
	fmt.Println(spec.Value(5).Should.Equal(5))
	fmt.Println(spec.Value("apple").Should.Not.Equal("orange"))
	fmt.Println(spec.Value(true).Should.Be.True())
	fmt.Println(spec.Value(false).Should.Not.Be.True())

	// Output:
	// true
	// true
	// true
	// true
}