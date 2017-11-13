// golang 的位运算
package main

import (
	"fmt"
)

func main() {
	a := 1
	switch {
	case a >= 0:
		fmt.Println("a=0")
		fallthrough //如果成立了继续执行代码 关键字
	case a >= 1:
		fmt.Println("a=1")
	default:
		fmt.Println("none")
	}
}
