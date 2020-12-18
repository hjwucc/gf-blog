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
