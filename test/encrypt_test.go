package test

import (
	"fmt"
	"github.com/gogf/gf/crypto/gsha1"
)

func ExampleSha1_test() {
	fmt.Println(gsha1.Encrypt(""))
	// Output:
	// true
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
