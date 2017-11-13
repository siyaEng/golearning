// golang 的位运算
package main

import (
	"fmt"
)

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
)

func main() {

	fmt.Println(MB)
}
