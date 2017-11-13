// golang 的位运算
package main

import (
	"fmt"
)

func main() {
	a := 1
	var p *int = &a
	fmt.Println(*p)

	if b := 1; b < 1 {
		fmt.Println("you find")
	} else {
		fmt.Println("you find out")
	}
}
