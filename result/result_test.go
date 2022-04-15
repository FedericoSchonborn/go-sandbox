//go:build go1.18

package result_test

import (
	"fmt"

	"github.com/FedericoSchonborn/go-sandbox/result"
	"github.com/FedericoSchonborn/go-sandbox/result/strconv"
)

func ExampleMap() {
	value := result.Map(strconv.Atoi("2"), func(i int) int {
		return i * 2
	})
	fmt.Println(value)

	// Output:
	// Ok(4)
}
