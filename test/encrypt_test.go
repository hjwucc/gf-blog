package test

import (
	"fmt"
)

func ExampleSha1_test() {

}

func Example_test() {
	var x int
	inc := func() int {
		x++
		return x
	}
	fmt.Println(func() (a, b int) {
		return inc(), inc()
	}())

	// Output:
	// 1 2
}
