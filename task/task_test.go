//go:build go1.18

package task_test

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/FedericoSchonborn/go-sandbox/task"
)

func Example() {
	waitOneSecondAndReturnFive := func() int {
		time.Sleep(1 * time.Second)
		return 5
	}

	fmt.Println(task.New(waitOneSecondAndReturnFive).Join())
	// Output:
	// 5
}

func ExampleJoin() {
	rand.Seed(time.Now().Unix())
	returnSix := func() int {
		return 6
	}

	handles := make([]*task.Handle[int], 10)
	for i := 0; i < 10; i++ {
		handles[i] = task.New(returnSix)
	}

	fmt.Println(task.Join(handles...))
	// Output:
	// [6 6 6 6 6 6 6 6 6 6]
}
